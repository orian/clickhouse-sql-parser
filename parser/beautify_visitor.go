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
type BeautifyVisitor struct {
	DefaultASTVisitor
	builder strings.Builder
	depth   int
	indent  string
}

func NewBeautifyVisitor() *BeautifyVisitor {
	v := &BeautifyVisitor{indent: "  "}
	v.Self = v
	return v
}

func (b *BeautifyVisitor) String() string { return b.builder.String() }

// writeIndent emits depth-many copies of the indent string.
func (b *BeautifyVisitor) writeIndent() {
	for i := 0; i < b.depth; i++ {
		b.builder.WriteString(b.indent)
	}
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
func (b *BeautifyVisitor) writeSpace() { b.builder.WriteByte(' ') }

func (b *BeautifyVisitor) writeString(s string) { b.builder.WriteString(s) }

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

	b.indentIn()
	for i, item := range s.SelectItems {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		b.writeString(item.String())
	}
	b.indentOut()

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
		b.writeString(s.Settings.String())
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
		b.writeString(strings.TrimPrefix(c.Engine.String(), " "))
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

// beautifyFrom emits FROM with the table expression on its own indented line.
//
//	FROM
//	  t
func (b *BeautifyVisitor) beautifyFrom(f *FromClause) {
	b.writeString("FROM")
	b.indentIn()
	b.newline()
	b.writeString(f.Expr.String())
	b.indentOut()
}

// beautifyWhere emits WHERE with each AND/OR conjunct on its own indented line.
//
//	WHERE
//	  a > 1
//	AND
//	  b < 10
func (b *BeautifyVisitor) beautifyWhere(w *WhereClause) {
	b.writeString("WHERE")
	conjuncts, ops := splitBoolean(w.Expr)
	b.indentIn()
	for i, c := range conjuncts {
		if i > 0 {
			b.indentOut()
			b.newline()
			b.writeString(ops[i-1])
			b.indentIn()
		}
		b.newline()
		b.writeString(c.String())
	}
	b.indentOut()
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

// beautifyOrderBy emits ORDER BY with each item on its own indented line.
func (b *BeautifyVisitor) beautifyOrderBy(o *OrderByClause) {
	b.writeString("ORDER BY")
	b.indentIn()
	for i, item := range o.Items {
		if i == 0 {
			b.newline()
		} else {
			b.writeString(",")
			b.newline()
		}
		b.writeString(item.String())
	}
	b.indentOut()
	if o.Interpolate != nil {
		b.newline()
		b.writeString(o.Interpolate.String())
	}
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
		b.writeString(col.String())
	}
	b.indentOut()
	b.newline()
	b.writeString(")")
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
		b.writeString(c.Settings.String())
	}
	if c.HasAppend {
		b.newline()
		b.writeString("APPEND")
	}
	if c.Engine != nil {
		b.newline()
		b.writeString(strings.TrimPrefix(c.Engine.String(), " "))
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
