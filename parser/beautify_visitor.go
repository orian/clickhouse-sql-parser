package parser

import "strings"

// BeautifyVisitor emits indented, line-broken SQL — the "beautify" counterpart
// to PrintVisitor. Like every visitor on this branch, it embeds DefaultASTVisitor
// so unimplemented Visit methods fall back to default descent.
//
// Architectural note: formatting belongs to the visitor, not to the nodes.
// Upstream introduced a Formatter type with a FormatSQL(*Formatter) method on
// every AST node (commit 9275c63). That breaks our branch's invariant — see
// refactor-visitor.md — so we keep formatting logic inside the visitor and
// leave the nodes untouched.
//
// Coverage strategy: top-level statement types (SelectQuery, CreateTable,
// CreateView, CreateMaterializedView, AlterTable, InsertStmt) get hand-written
// VisitX methods that break across lines and indent. Everything else falls
// through to the node's String() method (compact). Add VisitX overrides
// incrementally as specific subtrees need beautified output.
// BeautifyOptions configures the per-clause inline-vs-break thresholds
// used by BeautifyVisitor. A given clause stays on a single line when its
// rendered form fits both the MaxLen budget and (when applicable) the
// MaxCount item limit; otherwise it falls back to the multi-line layout.
// A zero MaxLen or MaxCount means "always break".
type BeautifyOptions struct {
	// MaxWidth is the soft line-length budget used by the expression-
	// level wrap logic (e.g. splitting long function calls).
	MaxWidth int

	// SelectItemsMaxLen / SelectItemsMaxCount govern whether SELECT
	// items stay on the SELECT line:
	//
	//   SELECT a, b, c   FROM t
	//
	// instead of breaking onto separate indented lines.
	SelectItemsMaxLen   int
	SelectItemsMaxCount int

	// FromMaxLen governs whether `FROM <expr>` stays on one line.
	// FROM expressions that contain a parenthesised subquery always
	// break regardless of this value, so the subquery can beautify.
	FromMaxLen int

	// WhereMaxLen / WhereMaxConjuncts govern whether `WHERE <expr>`
	// stays on one line. A conjunct is one top-level AND/OR-separated
	// piece of the predicate; a default of 1 means any boolean
	// splitting forces the multi-line layout.
	WhereMaxLen       int
	WhereMaxConjuncts int

	// OrderByMaxLen / OrderByMaxCount govern whether ORDER BY items
	// stay on one line.
	OrderByMaxLen   int
	OrderByMaxCount int
}

// DefaultBeautifyOptions matches the conventions the visitor used before
// these knobs were exposed: 100-col wrap budget, with 80-char inline
// budgets per clause and reasonable item counts.
var DefaultBeautifyOptions = BeautifyOptions{
	MaxWidth:            100,
	SelectItemsMaxLen:   80,
	SelectItemsMaxCount: 7,
	FromMaxLen:          80,
	WhereMaxLen:         80,
	WhereMaxConjuncts:   1,
	OrderByMaxLen:       80,
	OrderByMaxCount:     5,
}

type BeautifyVisitor struct {
	DefaultASTVisitor
	builder  strings.Builder
	depth    int
	indent   string
	col      int // running column of the next byte to be written
	maxWidth int // soft wrap budget for compact renderings (0 = no wrap)
	opts     BeautifyOptions
}

func NewBeautifyVisitor() *BeautifyVisitor {
	return NewBeautifyVisitorWith(DefaultBeautifyOptions)
}

// NewBeautifyVisitorWith creates a BeautifyVisitor with the given option set.
// Useful when callers want different per-clause inline thresholds.
func NewBeautifyVisitorWith(opts BeautifyOptions) *BeautifyVisitor {
	v := &BeautifyVisitor{
		indent:   "  ",
		maxWidth: opts.MaxWidth,
		opts:     opts,
	}
	v.Self = v
	return v
}

func (b *BeautifyVisitor) String() string { return b.builder.String() }

// writeIndent emits depth-many copies of the indent string and resets the
// column to its trailing position.
func (b *BeautifyVisitor) writeIndent() {
	for i := 0; i < b.depth; i++ {
		b.builder.WriteString(b.indent)
	}
	b.col = b.depth * len(b.indent)
}

// newline emits a newline followed by the current indent on the next line.
func (b *BeautifyVisitor) newline() {
	b.builder.WriteByte('\n')
	b.writeIndent()
}

// indent / dedent manage the indentation depth.
func (b *BeautifyVisitor) indentIn()  { b.depth++ }
func (b *BeautifyVisitor) indentOut() { b.depth-- }

// writeSpace emits a single space.
func (b *BeautifyVisitor) writeSpace() {
	b.builder.WriteByte(' ')
	b.col++
}

func (b *BeautifyVisitor) writeString(s string) {
	b.builder.WriteString(s)
	if i := strings.LastIndexByte(s, '\n'); i >= 0 {
		b.col = len(s) - i - 1
	} else {
		b.col += len(s)
	}
}

// selectItemsInlineMaxLen is the rendered-length budget for keeping SELECT
// items on the same line as the SELECT keyword (without the keyword/space
// prefix). Combined with selectItemsInlineMaxCount it controls when a
// short SELECT stays compact:
//
//	SELECT a, b, c FROM t
//
// instead of breaking each item onto its own indented line.
// selectItemsFitInline reports whether the SELECT items should stay on
// the SELECT line. They stay inline only when there are at most
// opts.SelectItemsMaxCount items and the joined `item1, item2, ...`
// rendering is shorter than opts.SelectItemsMaxLen.
func (b *BeautifyVisitor) selectItemsFitInline(items []*SelectItem) bool {
	if b.opts.SelectItemsMaxCount <= 0 || b.opts.SelectItemsMaxLen <= 0 {
		return false
	}
	if len(items) > b.opts.SelectItemsMaxCount {
		return false
	}
	total := 0
	for i, item := range items {
		if i > 0 {
			total += 2 // ", "
		}
		s := item.String()
		if strings.ContainsRune(s, '\n') {
			return false
		}
		total += len(s)
		if total >= b.opts.SelectItemsMaxLen {
			return false
		}
	}
	return true
}

// emitExpr writes an expression respecting the line-width budget. If its
// compact one-line rendering (via .String()) would still fit within
// maxWidth at the current column, it's emitted as-is; otherwise the
// expression is dispatched to a node-type-specific multi-line emitter,
// which recursively applies the same logic to children. The upshot is
// that splitting cascades top-down: a SELECT item that's a list of
// function calls splits each call onto its own line first; if any one
// call is still too long, that call's arguments split next, and so on.
func (b *BeautifyVisitor) emitExpr(expr Expr) {
	if expr == nil {
		return
	}
	if b.fitsCompact(expr) {
		b.writeString(expr.String())
		return
	}
	switch e := expr.(type) {
	case *SelectItem:
		b.emitSelectItemMultiLine(e)
	case *ColumnExpr:
		b.emitColumnExprMultiLine(e)
	case *FunctionExpr:
		b.emitFunctionMultiLine(e)
	case *AliasExpr:
		b.emitAliasExprMultiLine(e)
	case *BinaryOperation:
		b.emitBinaryOpMultiLine(e)
	case *OrderExpr:
		b.emitOrderExprMultiLine(e)
	case *TernaryOperation:
		b.emitTernaryMultiLine(e)
	case *ArrayParamList:
		b.emitArrayParamListMultiLine(e)
	default:
		// Unsplittable node type — fall back to compact even if it
		// blows past the budget. Add a dedicated emitX method here
		// to extend coverage.
		b.writeString(expr.String())
	}
}

// fitsCompact reports whether expr's compact rendering would stay within
// the wrap budget given the current column. A zero (or negative) maxWidth
// disables wrapping.
func (b *BeautifyVisitor) fitsCompact(expr Expr) bool {
	if b.maxWidth <= 0 {
		return true
	}
	s := expr.String()
	if strings.ContainsRune(s, '\n') {
		return false
	}
	return b.col+len(s) <= b.maxWidth
}

// emitSelectItemMultiLine prints a SELECT item with a multi-line expression.
// Modifiers and AS-alias stay on the closing line.
func (b *BeautifyVisitor) emitSelectItemMultiLine(s *SelectItem) {
	b.emitExpr(s.Expr)
	for _, mod := range s.Modifiers {
		b.writeSpace()
		b.writeString(mod.String())
	}
	if s.Alias != nil {
		b.writeString(" AS ")
		b.writeString(s.Alias.String())
	}
}

// emitColumnExprMultiLine prints a ColumnExpr (Expr [AS Alias]) with the
// inner expression possibly multi-line and the alias on the closing line.
func (b *BeautifyVisitor) emitColumnExprMultiLine(c *ColumnExpr) {
	b.emitExpr(c.Expr)
	if c.Alias != nil {
		b.writeString(" AS ")
		b.writeString(c.Alias.String())
	}
}

// emitAliasExprMultiLine handles `<expr> AS <alias>` when the wrapped expr
// is long enough to need splitting.
func (b *BeautifyVisitor) emitAliasExprMultiLine(a *AliasExpr) {
	b.emitExpr(a.Expr)
	if a.Alias != nil {
		b.writeString(" AS ")
		b.writeString(a.Alias.String())
	}
}

// emitTernaryMultiLine prints `cond ? then : else` with the `?` and `:`
// moved to the start of indented continuation lines:
//
//	condition
//	  ? then_expr
//	  : else_expr
//
// Each branch goes back through emitExpr so a long branch can recursively
// split further.
func (b *BeautifyVisitor) emitTernaryMultiLine(t *TernaryOperation) {
	b.emitExpr(t.Condition)
	b.indentIn()
	b.newline()
	b.writeString("? ")
	b.emitExpr(t.TrueExpr)
	b.newline()
	b.writeString(": ")
	b.emitExpr(t.FalseExpr)
	b.indentOut()
}

// emitArrayParamListMultiLine prints `[item1, item2, ...]` with each
// element on its own indented line when the inline form overshoots the
// budget. Each element is routed back through emitExpr for recursive
// splitting.
func (b *BeautifyVisitor) emitArrayParamListMultiLine(a *ArrayParamList) {
	if a.Items == nil || len(a.Items.Items) == 0 {
		b.writeString(a.String())
		return
	}
	b.writeString("[")
	b.indentIn()
	for i, item := range a.Items.Items {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		b.emitExpr(item)
	}
	b.indentOut()
	b.newline()
	b.writeString("]")
}

// emitOrderExprMultiLine prints an ORDER BY item where the underlying
// expression is too long to fit. The direction (ASC/DESC) and optional
// WITH FILL stay on the closing line, with the body recursively split:
//
//	if(
//	  has(...),
//	  2,
//	  if(has(...), 1, 0)
//	) ASC
func (b *BeautifyVisitor) emitOrderExprMultiLine(o *OrderExpr) {
	b.emitExpr(o.Expr)
	if o.Alias != nil {
		b.writeString(" AS ")
		b.writeString(o.Alias.String())
	}
	if o.Direction != OrderDirectionNone {
		b.writeSpace()
		b.writeString(string(o.Direction))
	}
	if o.Fill != nil {
		b.writeSpace()
		b.writeString(o.Fill.String())
	}
}

// emitBinaryOpMultiLine prints `left <op> right` with the operator at the
// start of a new line and each operand routed back through emitExpr so a
// long operand (typically a function call) further splits its arguments.
//
// Two operator categories are handled specially:
//
//   - Tight-binding operators (::-cast, `.` member access) are never
//     broken: their compact form is emitted even if it overshoots the
//     budget, because splitting them produces obviously-wrong-looking SQL.
//   - Lambda arrows (->) keep `<param> -> ` on one line and split only
//     the body. Putting `->` on a new line, separated from its parameter,
//     reads as broken SQL.
func (b *BeautifyVisitor) emitBinaryOpMultiLine(p *BinaryOperation) {
	if isTightBinaryOp(p.Operation) {
		b.writeString(p.String())
		return
	}
	if string(p.Operation) == "->" {
		b.writeString(p.LeftExpr.String())
		b.writeString(" -> ")
		b.emitExpr(p.RightExpr)
		return
	}
	b.emitExpr(p.LeftExpr)
	b.newline()
	if p.HasNot {
		b.writeString("NOT ")
	} else if p.HasGlobal {
		b.writeString("GLOBAL ")
	}
	b.writeString(string(p.Operation))
	b.writeSpace()
	b.emitExpr(p.RightExpr)
}

// isTightBinaryOp reports operators that should never be split across lines.
func isTightBinaryOp(op TokenKind) bool {
	switch string(op) {
	case "::", ".":
		return true
	}
	return false
}

// emitFunctionMultiLine prints a function call with each argument on its
// own indented line. Each argument is itself routed through emitExpr so
// nested function calls split recursively when they overflow.
//
//	some_func(
//	  arg1,
//	  another_func(
//	    inner_arg1,
//	    inner_arg2
//	  ),
//	  arg3
//	)
func (b *BeautifyVisitor) emitFunctionMultiLine(f *FunctionExpr) {
	b.writeString(f.Name.String())
	b.writeString("(")
	if f.Params == nil || f.Params.Items == nil || len(f.Params.Items.Items) == 0 {
		b.writeString(")")
		return
	}
	if f.Params.Items.HasDistinct {
		b.indentIn()
		b.newline()
		b.writeString("DISTINCT")
		// Continue with items on the same indent level as DISTINCT.
		for _, item := range f.Params.Items.Items {
			b.writeString(",")
			b.newline()
			b.emitExpr(item)
		}
		b.indentOut()
		b.newline()
	} else {
		b.indentIn()
		for i, item := range f.Params.Items.Items {
			if i == 0 {
				b.newline()
			} else {
				b.writeString(",")
				b.newline()
			}
			b.emitExpr(item)
		}
		b.indentOut()
		b.newline()
	}
	b.writeString(")")
	if f.Params.ColumnArgList != nil {
		b.writeString(f.Params.ColumnArgList.String())
	}
}

// VisitSelectQuery beautifies a SELECT statement:
//
//	SELECT
//	  col1,
//	  col2
//	FROM t
//	WHERE ...
func (b *BeautifyVisitor) VisitSelectQuery(s *SelectQuery) error {
	b.Enter(s)
	defer b.Leave(s)

	if s.With != nil {
		b.writeString("WITH")
		b.indentIn()
		for i, cte := range s.With.CTEs {
			if i == 0 {
				b.newline()
			} else {
				b.writeString(",")
				b.newline()
			}
			b.writeString(cte.String())
		}
		b.indentOut()
		b.newline()
	}

	b.writeString("SELECT")
	if s.HasDistinct {
		b.writeString(" DISTINCT")
		if s.DistinctOn != nil {
			b.writeSpace()
			b.writeString(s.DistinctOn.String())
		}
	}
	if s.Top != nil {
		b.writeSpace()
		b.writeString(s.Top.String())
	}

	if b.selectItemsFitInline(s.SelectItems) {
		b.writeSpace()
		for i, item := range s.SelectItems {
			if i > 0 {
				b.writeString(", ")
			}
			b.writeString(item.String())
		}
	} else {
		b.indentIn()
		for i, item := range s.SelectItems {
			if i == 0 {
				b.newline()
			} else {
				b.writeString(",")
				b.newline()
			}
			b.emitExpr(item)
		}
		b.indentOut()
	}

	if s.From != nil {
		b.newline()
		b.beautifyFrom(s.From)
	}
	if s.Window != nil {
		b.newline()
		b.writeString(s.Window.String())
	}
	if s.Prewhere != nil {
		b.newline()
		b.writeString("PREWHERE")
		b.indentIn()
		b.newline()
		b.writeString(s.Prewhere.Expr.String())
		b.indentOut()
	}
	if s.Where != nil {
		b.newline()
		b.beautifyWhere(s.Where)
	}
	if s.GroupBy != nil {
		b.newline()
		b.beautifyGroupBy(s.GroupBy)
	}
	if s.Having != nil {
		b.newline()
		b.writeString("HAVING")
		b.indentIn()
		b.newline()
		b.writeString(s.Having.Expr.String())
		b.indentOut()
	}
	if s.OrderBy != nil {
		b.newline()
		b.beautifyOrderBy(s.OrderBy)
	}
	if s.LimitBy != nil {
		b.newline()
		b.writeString(s.LimitBy.String())
	}
	if s.Limit != nil {
		b.newline()
		b.writeString(s.Limit.String())
	}
	if s.Settings != nil {
		b.newline()
		b.beautifySettings(s.Settings)
	}
	if s.Format != nil {
		b.newline()
		b.writeString(s.Format.String())
	}

	if s.UnionAll != nil {
		b.newline()
		b.writeString("UNION ALL")
		b.newline()
		if err := s.UnionAll.Accept(b.Self); err != nil {
			return err
		}
	} else if s.UnionDistinct != nil {
		b.newline()
		b.writeString("UNION DISTINCT")
		b.newline()
		if err := s.UnionDistinct.Accept(b.Self); err != nil {
			return err
		}
	} else if s.Except != nil {
		b.newline()
		b.writeString("EXCEPT")
		b.newline()
		if err := s.Except.Accept(b.Self); err != nil {
			return err
		}
	}
	return nil
}

// VisitCreateTable beautifies CREATE TABLE.
func (b *BeautifyVisitor) VisitCreateTable(c *CreateTable) error {
	b.Enter(c)
	defer b.Leave(c)

	b.writeString("CREATE")
	if c.OrReplace {
		b.writeString(" OR REPLACE")
	}
	if c.HasTemporary {
		b.writeString(" TEMPORARY")
	}
	b.writeString(" TABLE ")
	if c.IfNotExists {
		b.writeString("IF NOT EXISTS ")
	}
	b.writeString(c.Name.String())
	if c.UUID != nil {
		b.writeSpace()
		b.writeString(c.UUID.String())
	}
	if c.OnCluster != nil {
		b.newline()
		b.writeString(c.OnCluster.String())
	}
	if c.TableSchema != nil {
		// Column list — break per column.
		b.writeSpace()
		b.beautifyTableSchema(c.TableSchema)
	}
	if c.Engine != nil {
		b.newline()
		b.beautifyEngine(c.Engine)
	}
	if c.SubQuery != nil {
		b.newline()
		b.writeString("AS")
		b.newline()
		if err := c.SubQuery.Select.Accept(b.Self); err != nil {
			return err
		}
	}
	if c.TableFunction != nil {
		b.newline()
		b.writeString("AS ")
		b.writeString(c.TableFunction.String())
	}
	if c.Comment != nil {
		b.newline()
		b.writeString("COMMENT ")
		b.writeString(c.Comment.String())
	}
	return nil
}

// beautifyFrom emits FROM with the table expression. If the whole
// `FROM <expr>` rendering fits within the wrap budget at the current
// column, it stays on one line (e.g. `FROM t`); otherwise the source is
// moved to its own indented line and any subselect inside is recursively
// beautified:
//
//	FROM
//	  (
//	    SELECT ...
//	  )
func (b *BeautifyVisitor) beautifyFrom(f *FromClause) {
	if b.fromFitsInline(f) {
		b.writeString("FROM ")
		b.writeString(f.Expr.String())
		return
	}
	b.writeString("FROM")
	b.indentIn()
	b.newline()
	b.beautifyFromExpr(f.Expr)
	b.indentOut()
}

// fromFitsInline reports whether `FROM <expr>` should stay on a single
// line: the source must contain no subqueries (those always beautify
// into an indented block) and the rendered length must stay within
// opts.FromMaxLen.
func (b *BeautifyVisitor) fromFitsInline(f *FromClause) bool {
	if b.opts.FromMaxLen <= 0 {
		return false
	}
	if fromHasSubQuery(f.Expr) {
		return false
	}
	s := f.Expr.String()
	if strings.ContainsRune(s, '\n') {
		return false
	}
	return len(s) < b.opts.FromMaxLen
}

// fromHasSubQuery reports whether the FROM expression tree contains a
// parenthesised SELECT — i.e. anything that should beautify as an
// indented (\n  SELECT … \n) block rather than stay inline.
func fromHasSubQuery(expr Expr) bool {
	switch e := expr.(type) {
	case *SubQuery:
		return e.HasParen
	case *AliasExpr:
		return fromHasSubQuery(e.Expr)
	case *TableExpr:
		return fromHasSubQuery(e.Expr)
	case *JoinTableExpr:
		if e.Table != nil {
			return fromHasSubQuery(e.Table)
		}
	case *JoinExpr:
		return fromHasSubQuery(e.Left) || fromHasSubQuery(e.Right)
	}
	return false
}

// beautifyFromExpr renders whatever sits under FROM. For now only the
// JoinTableExpr → TableExpr → SubQuery shape is line-broken; everything else
// (joins, table functions, plain idents) falls back to compact .String().
func (b *BeautifyVisitor) beautifyFromExpr(expr Expr) {
	jt, ok := expr.(*JoinTableExpr)
	if !ok || jt.Table == nil {
		b.writeString(expr.String())
		return
	}
	b.beautifyTableExpr(jt.Table)
	if jt.SampleRatio != nil {
		b.writeSpace()
		b.writeString(jt.SampleRatio.String())
	}
	if jt.HasFinal {
		b.writeString(" FINAL")
	}
}

// beautifyTableExpr emits a TableExpr; if its inner expression is a SubQuery
// (optionally wrapped in an AliasExpr for `(SELECT …) AS x`), the subselect is
// rendered as a beautified, indented block.
func (b *BeautifyVisitor) beautifyTableExpr(t *TableExpr) {
	sub, suffix := unwrapSubQueryTable(t.Expr)
	if sub == nil || !sub.HasParen {
		b.writeString(t.String())
		return
	}
	b.writeString("(")
	b.indentIn()
	b.newline()
	// Recurse through the visitor so VisitSelectQuery formats the inner
	// SELECT (and any nested subqueries) the same way as the outer one.
	_ = sub.Select.Accept(b.Self)
	b.indentOut()
	b.newline()
	b.writeString(")")
	if suffix != "" {
		b.writeSpace()
		b.writeString(suffix)
	}
	if t.Alias != nil {
		b.writeSpace()
		b.writeString(t.Alias.String())
	}
	if t.HasFinal {
		b.writeString(" FINAL")
	}
}

// unwrapSubQueryTable returns the inner SubQuery if expr is one (possibly
// wrapped in an AliasExpr for `(SELECT …) AS alias`). The second return is
// the alias text to print after the closing paren, or "" if there isn't one.
func unwrapSubQueryTable(expr Expr) (*SubQuery, string) {
	switch e := expr.(type) {
	case *SubQuery:
		return e, ""
	case *AliasExpr:
		if inner, ok := e.Expr.(*SubQuery); ok {
			return inner, "AS " + e.Alias.String()
		}
	}
	return nil, ""
}

// beautifyWhere emits `WHERE <expr>`. Short single-conjunct predicates
// stay on the same line as WHERE (`WHERE x > 1`); anything longer or
// joined with AND/OR breaks each conjunct onto its own indented line:
//
//	WHERE
//	  a > 1
//	AND
//	  b < 10
func (b *BeautifyVisitor) beautifyWhere(w *WhereClause) {
	conjuncts, ops := splitBoolean(w.Expr)
	if b.whereFitsInline(conjuncts) {
		b.writeString("WHERE ")
		for i, c := range conjuncts {
			if i > 0 {
				b.writeSpace()
				b.writeString(ops[i-1])
				b.writeSpace()
			}
			b.writeString(c.String())
		}
		return
	}
	b.writeString("WHERE")
	b.indentIn()
	for i, c := range conjuncts {
		if i > 0 {
			b.indentOut()
			b.newline()
			b.writeString(ops[i-1])
			b.indentIn()
		}
		b.newline()
		b.emitExpr(c)
	}
	b.indentOut()
}

// whereFitsInline reports whether the WHERE conjuncts fit on a single
// line: at most opts.WhereMaxConjuncts of them (default 1), joined
// rendering shorter than opts.WhereMaxLen.
func (b *BeautifyVisitor) whereFitsInline(conjuncts []Expr) bool {
	if b.opts.WhereMaxConjuncts <= 0 || b.opts.WhereMaxLen <= 0 {
		return false
	}
	if len(conjuncts) > b.opts.WhereMaxConjuncts {
		return false
	}
	total := 0
	for i, c := range conjuncts {
		if i > 0 {
			total += 5 // " AND " / " OR " upper bound
		}
		s := c.String()
		if strings.ContainsRune(s, '\n') {
			return false
		}
		total += len(s)
		if total >= b.opts.WhereMaxLen {
			return false
		}
	}
	return true
}

// splitBoolean splits an expression on top-level AND/OR into a flat list.
// Returns the conjuncts and the operators joining them. For "a AND b OR c"
// returns ([a, b, c], ["AND", "OR"]).
func splitBoolean(expr Expr) ([]Expr, []string) {
	bop, ok := expr.(*BinaryOperation)
	if !ok {
		return []Expr{expr}, nil
	}
	opStr := strings.ToUpper(string(bop.Operation))
	if opStr != "AND" && opStr != "OR" {
		return []Expr{expr}, nil
	}
	left, lops := splitBoolean(bop.LeftExpr)
	right, rops := splitBoolean(bop.RightExpr)
	conjuncts := append(left, right...)
	ops := append(lops, opStr)
	ops = append(ops, rops...)
	return conjuncts, ops
}

// beautifyGroupBy emits GROUP BY with the expression(s) on their own indented line.
func (b *BeautifyVisitor) beautifyGroupBy(g *GroupByClause) {
	b.writeString("GROUP BY")
	if g.AggregateType != "" {
		b.writeSpace()
		b.writeString(g.AggregateType)
	}
	if g.Expr != nil {
		b.indentIn()
		b.newline()
		b.writeString(g.Expr.String())
		b.indentOut()
	}
	if g.WithCube {
		b.newline()
		b.writeString("WITH CUBE")
	}
	if g.WithRollup {
		b.newline()
		b.writeString("WITH ROLLUP")
	}
	if g.WithTotals {
		b.newline()
		b.writeString("WITH TOTALS")
	}
}

// beautifyEngine emits ENGINE = Name(args) and each engine clause
// (ORDER BY, PARTITION BY, PRIMARY KEY, SAMPLE BY, TTL, SETTINGS) on its own
// line at the same indent as ENGINE. SETTINGS recursively beautifies to one
// setting per line.
func (b *BeautifyVisitor) beautifyEngine(e *EngineExpr) {
	b.writeString("ENGINE = ")
	b.writeString(e.Name)
	if e.Params != nil {
		b.writeString(e.Params.String())
	}
	if e.OrderBy != nil {
		b.newline()
		b.writeString(e.OrderBy.String())
	}
	if e.PartitionBy != nil {
		b.newline()
		b.writeString(e.PartitionBy.String())
	}
	if e.PrimaryKey != nil {
		b.newline()
		b.writeString(e.PrimaryKey.String())
	}
	if e.SampleBy != nil {
		b.newline()
		b.writeString(e.SampleBy.String())
	}
	if e.TTL != nil {
		b.newline()
		b.writeString(e.TTL.String())
	}
	if e.Settings != nil {
		b.newline()
		b.beautifySettings(e.Settings)
	}
}

// beautifySettings emits SETTINGS with each item on its own indented line:
//
//	SETTINGS
//	  max_threads = 8,
//	  max_memory_usage = 1000000
func (b *BeautifyVisitor) beautifySettings(s *SettingsClause) {
	b.writeString("SETTINGS")
	b.indentIn()
	for i, item := range s.Items {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		// SettingExpr.String() emits `name=value` with no spaces; render
		// the beautified form with spaces around the `=`.
		b.writeString(item.Name.String())
		b.writeString(" = ")
		b.writeString(item.Expr.String())
	}
	b.indentOut()
}

// beautifyOrderBy emits ORDER BY. Short item lists stay on the same line
// (`ORDER BY a, b DESC`); longer lists break each item onto its own
// indented line.
func (b *BeautifyVisitor) beautifyOrderBy(o *OrderByClause) {
	if b.orderByFitsInline(o.Items) {
		b.writeString("ORDER BY ")
		for i, item := range o.Items {
			if i > 0 {
				b.writeString(", ")
			}
			b.writeString(item.String())
		}
		if o.Interpolate != nil {
			b.newline()
			b.writeString(o.Interpolate.String())
		}
		return
	}
	b.writeString("ORDER BY")
	b.indentIn()
	for i, item := range o.Items {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		b.emitExpr(item)
	}
	b.indentOut()
	if o.Interpolate != nil {
		b.newline()
		b.writeString(o.Interpolate.String())
	}
}

// orderByFitsInline reports whether the ORDER BY items fit on a single
// line: at most opts.OrderByMaxCount of them and joined rendering
// shorter than opts.OrderByMaxLen.
func (b *BeautifyVisitor) orderByFitsInline(items []Expr) bool {
	if b.opts.OrderByMaxCount <= 0 || b.opts.OrderByMaxLen <= 0 {
		return false
	}
	if len(items) > b.opts.OrderByMaxCount {
		return false
	}
	total := 0
	for i, item := range items {
		if i > 0 {
			total += 2 // ", "
		}
		s := item.String()
		if strings.ContainsRune(s, '\n') {
			return false
		}
		total += len(s)
		if total >= b.opts.OrderByMaxLen {
			return false
		}
	}
	return true
}

func (b *BeautifyVisitor) beautifyTableSchema(t *TableSchemaClause) {
	// Mirror TableSchemaClause.String but with each column on its own line.
	if len(t.Columns) == 0 {
		b.writeString(t.String())
		return
	}
	b.writeString("(")
	b.indentIn()
	for i, col := range t.Columns {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		if cd, ok := col.(*ColumnDef); ok {
			b.emitColumnDef(cd)
		} else {
			// INDEX / CONSTRAINT / PROJECTION entries: compact form.
			b.writeString(col.String())
		}
	}
	b.indentOut()
	b.newline()
	b.writeString(")")
}

// emitColumnDef prints a column definition, routing the DEFAULT /
// MATERIALIZED / ALIAS expression through emitExpr so a long function call
// (e.g. a MATERIALIZED replaceRegexpAll(...)) wraps across indented lines
// instead of running off as one giant line. The name+type prefix and the
// trailing modifiers (codec, TTL, COMMENT) stay on the same line as the
// keyword and the closing paren respectively:
//
//	`pmat_email` String MATERIALIZED replaceRegexpAll(
//	  JSONExtractRaw(properties, 'email'),
//	  concat(...),
//	  ''
//	) COMMENT 'column_materializer::email'
func (b *BeautifyVisitor) emitColumnDef(c *ColumnDef) {
	b.writeString(c.Name.String())
	if c.Type != nil {
		b.writeSpace()
		b.writeString(c.Type.String())
	}
	if c.NotNull != nil {
		b.writeString(" NOT NULL")
	} else if c.Nullable != nil {
		b.writeString(" NULL")
	}
	if c.DefaultExpr != nil {
		b.writeString(" DEFAULT ")
		b.emitExpr(c.DefaultExpr)
	}
	if c.MaterializedExpr != nil {
		b.writeString(" MATERIALIZED ")
		b.emitExpr(c.MaterializedExpr)
	}
	if c.IsEphemeral {
		b.writeString(" EPHEMERAL")
		if c.EphemeralExpr != nil {
			b.writeSpace()
			b.emitExpr(c.EphemeralExpr)
		}
	}
	if c.AliasExpr != nil {
		b.writeString(" ALIAS ")
		b.emitExpr(c.AliasExpr)
	}
	if c.Codec != nil {
		b.writeSpace()
		b.writeString(c.Codec.String())
	}
	if c.TTL != nil {
		b.writeSpace()
		b.writeString(c.TTL.String())
	}
	if c.Comment != nil {
		b.writeString(" COMMENT ")
		b.writeString(c.Comment.String())
	}
}

// VisitCreateView beautifies CREATE VIEW.
func (b *BeautifyVisitor) VisitCreateView(c *CreateView) error {
	b.Enter(c)
	defer b.Leave(c)

	b.writeString("CREATE")
	if c.OrReplace {
		b.writeString(" OR REPLACE")
	}
	b.writeString(" VIEW ")
	if c.IfNotExists {
		b.writeString("IF NOT EXISTS ")
	}
	b.writeString(c.Name.String())
	if c.UUID != nil {
		b.writeSpace()
		b.writeString(c.UUID.String())
	}
	if c.OnCluster != nil {
		b.newline()
		b.writeString(c.OnCluster.String())
	}
	if c.TableSchema != nil {
		b.writeSpace()
		b.beautifyTableSchema(c.TableSchema)
	}
	if c.Definer != nil {
		b.newline()
		b.writeString("DEFINER = ")
		b.writeString(c.Definer.String())
	}
	if c.SQLSecurity != "" {
		b.newline()
		b.writeString("SQL SECURITY ")
		b.writeString(c.SQLSecurity)
	}
	if c.Comment != nil {
		b.newline()
		b.writeString("COMMENT ")
		b.writeString(c.Comment.String())
	}
	if c.SubQuery != nil {
		b.newline()
		b.writeString("AS")
		b.newline()
		if err := c.SubQuery.Select.Accept(b.Self); err != nil {
			return err
		}
	}
	return nil
}

// VisitCreateMaterializedView beautifies CREATE MATERIALIZED VIEW.
func (b *BeautifyVisitor) VisitCreateMaterializedView(c *CreateMaterializedView) error {
	b.Enter(c)
	defer b.Leave(c)

	b.writeString("CREATE MATERIALIZED VIEW ")
	if c.IfNotExists {
		b.writeString("IF NOT EXISTS ")
	}
	b.writeString(c.Name.String())
	if c.OnCluster != nil {
		b.newline()
		b.writeString(c.OnCluster.String())
	}
	if c.Refresh != nil {
		b.newline()
		b.writeString(c.Refresh.String())
	}
	if c.RandomizeFor != nil {
		b.newline()
		b.writeString("RANDOMIZE FOR ")
		b.writeString(c.RandomizeFor.String())
	}
	if c.DependsOn != nil {
		b.newline()
		b.writeString("DEPENDS ON ")
		for i, dep := range c.DependsOn {
			if i > 0 {
				b.writeString(", ")
			}
			b.writeString(dep.String())
		}
	}
	if c.Settings != nil {
		b.newline()
		b.beautifySettings(c.Settings)
	}
	if c.HasAppend {
		b.newline()
		b.writeString("APPEND")
	}
	if c.Engine != nil {
		b.newline()
		b.beautifyEngine(c.Engine)
	}
	if c.Destination != nil {
		b.newline()
		b.writeString(c.Destination.String())
		if c.Destination.TableSchema != nil {
			b.writeSpace()
			b.beautifyTableSchema(c.Destination.TableSchema)
		}
	}
	if c.HasEmpty {
		b.newline()
		b.writeString("EMPTY")
	}
	if c.Definer != nil {
		b.newline()
		b.writeString("DEFINER = ")
		b.writeString(c.Definer.String())
	}
	if c.SQLSecurity != "" {
		b.newline()
		b.writeString("SQL SECURITY ")
		b.writeString(c.SQLSecurity)
	}
	if c.Populate {
		b.newline()
		b.writeString("POPULATE")
	}
	if c.SubQuery != nil {
		b.newline()
		b.writeString("AS")
		b.newline()
		if err := c.SubQuery.Select.Accept(b.Self); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		b.newline()
		b.writeString("COMMENT ")
		b.writeString(c.Comment.String())
	}
	return nil
}

// VisitAlterTable beautifies ALTER TABLE.
func (b *BeautifyVisitor) VisitAlterTable(a *AlterTable) error {
	b.Enter(a)
	defer b.Leave(a)

	b.writeString("ALTER TABLE ")
	b.writeString(a.TableIdentifier.String())
	if a.OnCluster != nil {
		b.newline()
		b.writeString(a.OnCluster.String())
	}
	b.indentIn()
	for i, expr := range a.AlterExprs {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		b.writeString(expr.String())
	}
	b.indentOut()
	return nil
}

// VisitInsertExpr beautifies INSERT INTO.
func (b *BeautifyVisitor) VisitInsertExpr(i *InsertStmt) error {
	b.Enter(i)
	defer b.Leave(i)

	b.writeString("INSERT INTO ")
	if i.HasTableKeyword {
		b.writeString("TABLE ")
	}
	b.writeString(i.Table.String())
	if i.ColumnNames != nil {
		b.indentIn()
		b.newline()
		b.writeString(i.ColumnNames.String())
		b.indentOut()
	}
	if i.Format != nil {
		b.newline()
		b.writeString(i.Format.String())
	}
	if i.SelectExpr != nil {
		b.newline()
		if err := i.SelectExpr.Accept(b.Self); err != nil {
			return err
		}
	} else if len(i.Values) > 0 {
		b.newline()
		b.writeString("VALUES")
		b.indentIn()
		for j, value := range i.Values {
			if j == 0 {
				b.newline()
			} else {
				b.writeString(",")
				b.newline()
			}
			b.writeString(value.String())
		}
		b.indentOut()
	}
	return nil
}

// VisitSubQueryExpr beautifies a subquery wrapped in parens.
func (b *BeautifyVisitor) VisitSubQueryExpr(s *SubQuery) error {
	b.Enter(s)
	defer b.Leave(s)

	if s.HasParen {
		b.writeString("(")
		b.indentIn()
		b.newline()
		if err := s.Select.Accept(b.Self); err != nil {
			return err
		}
		b.indentOut()
		b.newline()
		b.writeString(")")
	} else {
		return s.Select.Accept(b.Self)
	}
	return nil
}
