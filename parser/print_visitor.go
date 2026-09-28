package parser

import (
	"reflect"
	"strings"
)

type PrintVisitor struct {
	DefaultASTVisitor
	builder *strings.Builder
}

func NewPrintVisitor() *PrintVisitor {
	v := &PrintVisitor{
		builder: &strings.Builder{},
	}
	v.Self = v
	return v
}

func (v *PrintVisitor) String() string {
	return v.builder.String()
}

func (p *PrintVisitor) VisitAliasExpr(a *AliasExpr) error {
	if _, isSelect := a.Expr.(*SelectQuery); isSelect {
		p.builder.WriteByte('(')
		if err := a.Expr.Accept(p); err != nil {
			return err
		}
		p.builder.WriteByte(')')
	} else {
		if err := a.Expr.Accept(p); err != nil {
			return err
		}
	}
	p.builder.WriteString(" AS ")
	if err := a.Alias.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitAlterRole(a *AlterRole) error {
	builder := p.builder
	builder.WriteString("ALTER ROLE ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	for i, roleRenamePair := range a.RoleRenamePairs {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := roleRenamePair.Accept(p); err != nil {
			return err
		}
	}
	if len(a.Settings) > 0 {
		builder.WriteString(" SETTINGS ")
		for i, setting := range a.Settings {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := setting.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}

// printOutputClauses prints the trailing `FORMAT fmt` and `SETTINGS ...` of a
// statement embedding OutputClauses.
func (p *PrintVisitor) printOutputClauses(o *OutputClauses) error {
	if o.Format != nil {
		p.builder.WriteByte(' ')
		if err := o.Format.Accept(p); err != nil {
			return err
		}
	}
	if o.OutputSettings != nil {
		p.builder.WriteByte(' ')
		if err := o.OutputSettings.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTable(a *AlterTable) error {
	builder := p.builder
	builder.WriteString("ALTER TABLE ")
	if err := a.TableIdentifier.Accept(p); err != nil {
		return err
	}
	if a.OnCluster != nil {
		builder.WriteString(" ")
		if err := a.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	for i, expr := range a.AlterExprs {
		builder.WriteString(" ")
		if err := expr.Accept(p); err != nil {
			return err
		}
		if i != len(a.AlterExprs)-1 {
			builder.WriteString(",")
		}
	}
	if err := p.printOutputClauses(&a.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableAddColumn(a *AlterTableAddColumn) error {
	p.builder.WriteString("ADD COLUMN ")
	if a.IfNotExists {
		p.builder.WriteString("IF NOT EXISTS ")
	}
	if err := a.Column.Accept(p); err != nil {
		return err
	}
	if a.After != nil {
		p.builder.WriteString(" AFTER ")
		if err := a.After.Accept(p); err != nil {
			return err
		}
	}
	if a.Settings != nil {
		p.builder.WriteByte(' ')
		return a.Settings.Accept(p)
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableAddIndex(a *AlterTableAddIndex) error {
	builder := p.builder
	builder.WriteString("ADD INDEX ")
	// IF NOT EXISTS goes between INDEX and the index name.
	if a.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := p.printTableIndexBody(a.Index); err != nil {
		return err
	}
	if a.After != nil {
		builder.WriteString(" AFTER ")
		if err := a.After.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableAddProjection(a *AlterTableAddProjection) error {
	builder := p.builder
	builder.WriteString("ADD PROJECTION ")
	if a.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := a.TableProjection.Accept(p); err != nil {
		return err
	}
	if a.After != nil {
		builder.WriteString(" AFTER ")
		if err := a.After.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableAttachPartition(a *AlterTableAttachPartition) error {
	builder := p.builder
	builder.WriteString("ATTACH ")
	if err := a.Partition.Accept(p); err != nil {
		return err
	}
	if a.From != nil {
		builder.WriteString(" FROM ")
		if err := a.From.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableClearColumn(a *AlterTableClearColumn) error {
	builder := p.builder
	builder.WriteString("CLEAR COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.ColumnName.Accept(p); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		builder.WriteString(" IN ")
		if err := a.PartitionExpr.Accept(p); err != nil {
			return err
		}
	}

	return nil
}

func (p *PrintVisitor) VisitAlterTableClearIndex(a *AlterTableClearIndex) error {
	builder := p.builder
	builder.WriteString("CLEAR INDEX ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.IndexName.Accept(p); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		builder.WriteString(" IN ")
		if err := a.PartitionExpr.Accept(p); err != nil {
			return err
		}
	}

	return nil
}

func (p *PrintVisitor) VisitAlterTableClearProjection(a *AlterTableClearProjection) error {
	builder := p.builder
	builder.WriteString("CLEAR PROJECTION ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.ProjectionName.Accept(p); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		builder.WriteString(" IN ")
		if err := a.PartitionExpr.Accept(p); err != nil {
			return err
		}
	}

	return nil
}
func (p *PrintVisitor) VisitAlterTableDetachPartition(a *AlterTableDetachPartition) error {
	builder := p.builder
	builder.WriteString("DETACH ")
	if err := a.Partition.Accept(p); err != nil {
		return err
	}
	if a.Settings != nil {
		builder.WriteByte(' ')
		if err := a.Settings.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableDropColumn(a *AlterTableDropColumn) error {
	builder := p.builder
	builder.WriteString("DROP COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.ColumnName.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	builder := p.builder
	builder.WriteString("DROP INDEX ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.IndexName.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	builder := p.builder
	builder.WriteString("DROP ")
	if a.HasDetached {
		builder.WriteString("DETACHED ")
	}
	if err := a.Partition.Accept(p); err != nil {
		return err
	}
	if a.Settings != nil {
		builder.WriteByte(' ')
		if err := a.Settings.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	builder := p.builder
	builder.WriteString("DROP PROJECTION ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.ProjectionName.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	builder := p.builder
	builder.WriteString("FREEZE")
	if a.Partition != nil {
		builder.WriteByte(' ')
		if err := a.Partition.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableMaterializeIndex(a *AlterTableMaterializeIndex) error {
	builder := p.builder
	builder.WriteString("MATERIALIZE INDEX")

	if a.IfExists {
		builder.WriteString(" IF EXISTS")
	}
	builder.WriteString(" ")
	if err := a.IndexName.Accept(p); err != nil {
		return err
	}
	if a.Partition != nil {
		builder.WriteString(" IN ")
		if err := a.Partition.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableMaterializeProjection(a *AlterTableMaterializeProjection) error {
	builder := p.builder
	builder.WriteString("MATERIALIZE PROJECTION")

	if a.IfExists {
		builder.WriteString(" IF EXISTS")
	}
	builder.WriteString(" ")
	if err := a.ProjectionName.Accept(p); err != nil {
		return err
	}
	if a.Partition != nil {
		builder.WriteString(" IN ")
		if err := a.Partition.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableModifyColumn(a *AlterTableModifyColumn) error {
	builder := p.builder
	builder.WriteString("MODIFY COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.Column.Accept(p); err != nil {
		return err
	}
	if a.RemovePropertyType != nil {
		if err := a.RemovePropertyType.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableModifyQuery(a *AlterTableModifyQuery) error {
	builder := p.builder
	builder.WriteString("MODIFY QUERY ")
	if err := a.SelectExpr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableDelete(a *AlterTableDelete) error {
	builder := p.builder
	builder.WriteString("DELETE")
	if a.InPartition != nil {
		builder.WriteString(" IN ")
		if err := a.InPartition.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(" WHERE ")
	if err := a.WhereClause.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableUpdate(a *AlterTableUpdate) error {
	builder := p.builder
	builder.WriteString("UPDATE ")
	for i, assignment := range a.Assignments {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := assignment.Accept(p); err != nil {
			return err
		}
	}
	if a.InPartition != nil {
		builder.WriteString(" IN ")
		if err := a.InPartition.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(" WHERE ")
	if err := a.WhereClause.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitUpdateAssignment(u *UpdateAssignment) error {
	builder := p.builder
	if err := u.Column.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" = ")
	if err := u.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableModifySetting(a *AlterTableModifySetting) error {
	builder := p.builder
	builder.WriteString("MODIFY SETTING ")
	for i, setting := range a.Settings {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := setting.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableResetSetting(a *AlterTableResetSetting) error {
	builder := p.builder
	builder.WriteString("RESET SETTING ")
	for i, setting := range a.Settings {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := setting.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	// The TTL clause prints its own TTL keyword.
	p.builder.WriteString("MODIFY ")
	return a.TTL.Accept(p)
}
func (p *PrintVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	p.builder.WriteString("REMOVE TTL")
	return nil
}

func (p *PrintVisitor) VisitAlterTableRenameColumn(a *AlterTableRenameColumn) error {
	builder := p.builder
	builder.WriteString("RENAME COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := a.OldColumnName.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" TO ")
	if err := a.NewColumnName.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	builder := p.builder
	builder.WriteString("REPLACE ")
	if err := a.Partition.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" FROM ")
	if err := a.Table.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitArrayParamList(a *ArrayParamList) error {
	builder := p.builder
	builder.WriteString("[")
	for i, item := range a.Items.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := item.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString("]")
	return nil
}

func (p *PrintVisitor) VisitValuesExpr(v *AssignmentValues) error {
	builder := p.builder
	builder.WriteByte('(')
	for i, value := range v.Values {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := value.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitBetweenClause(f *BetweenClause) error {
	// Expr is nil in a window frame (`ROWS BETWEEN ... AND ...`).
	if f.Expr != nil {
		if err := f.Expr.Accept(p); err != nil {
			return err
		}
		p.builder.WriteByte(' ')
	}
	p.builder.WriteString("BETWEEN ")
	if err := f.Between.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(" AND ")
	return f.And.Accept(p)
}
func (pv *PrintVisitor) VisitBinaryExpr(p *BinaryOperation) error {
	builder := pv.builder
	if err := p.LeftExpr.Accept(pv); err != nil {
		return err
	}
	if p.Operation != TokenKindDash {
		builder.WriteByte(' ')
	}
	if p.HasNot {
		builder.WriteString("NOT ")
	} else if p.HasGlobal {
		builder.WriteString("GLOBAL ")
	}
	builder.WriteString(string(p.Operation))
	if p.Operation != TokenKindDash {
		builder.WriteByte(' ')
	}
	if err := p.RightExpr.Accept(pv); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitCTEExpr(c *CTEStmt) error {
	builder := p.builder
	if err := c.Expr.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" AS ")
	if _, isSelect := c.Alias.(*SelectQuery); isSelect {
		builder.WriteByte('(')
		if err := c.Alias.Accept(p); err != nil {
			return err
		}
		builder.WriteByte(')')
	} else {
		if err := c.Alias.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitCaseExpr(c *CaseExpr) error {
	builder := p.builder
	builder.WriteString("CASE ")
	if c.Expr != nil {
		if err := c.Expr.Accept(p); err != nil {
			return err
		}
		builder.WriteByte(' ')
	}
	for i, when := range c.Whens {
		if i > 0 {
			builder.WriteByte(' ')
		}
		if err := when.Accept(p); err != nil {
			return err
		}
	}
	if c.Else != nil {
		builder.WriteString(" ELSE ")
		if err := c.Else.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(" END")
	return nil
}

func (p *PrintVisitor) VisitCastExpr(c *CastExpr) error {
	builder := p.builder
	builder.WriteString("CAST(")
	if err := c.Expr.Accept(p); err != nil {
		return err
	}
	if c.Separator == "," {
		builder.WriteString(", ")
	} else {
		builder.WriteString(" AS ")
	}
	if err := c.AsType.Accept(p); err != nil {
		return err
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitCheckExpr(c *CheckStmt) error {
	builder := p.builder
	builder.WriteString("CHECK TABLE ")
	if err := c.Table.Accept(p); err != nil {
		return err
	}
	if c.Partition != nil {
		builder.WriteString(" ")
		if err := c.Partition.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&c.OutputClauses); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitOnClusterExpr(o *ClusterClause) error {
	builder := p.builder
	builder.WriteString("ON CLUSTER ")
	if err := o.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitColumnArgList(c *ColumnArgList) error {
	builder := p.builder
	builder.WriteByte('(')
	if c.Distinct {
		builder.WriteString("DISTINCT ")
	}
	for i, item := range c.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := item.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitColumnDef(c *ColumnDef) error {
	builder := p.builder
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.Type != nil {
		builder.WriteByte(' ')
		if err := c.Type.Accept(p); err != nil {
			return err
		}
	}
	if c.NotNull != nil {
		builder.WriteString(" NOT NULL")
	} else if c.Nullable != nil {
		builder.WriteString(" NULL")
	}
	if c.DefaultExpr != nil {
		builder.WriteString(" DEFAULT ")
		if err := c.DefaultExpr.Accept(p); err != nil {
			return err
		}
	}
	if c.MaterializedExpr != nil {
		builder.WriteString(" MATERIALIZED ")
		if err := c.MaterializedExpr.Accept(p); err != nil {
			return err
		}
	}
	if c.IsEphemeral {
		builder.WriteString(" EPHEMERAL")
		if c.EphemeralExpr != nil {
			builder.WriteByte(' ')
			if err := c.EphemeralExpr.Accept(p); err != nil {
				return err
			}
		}
	}
	if c.AliasExpr != nil {
		builder.WriteString(" ALIAS ")
		if err := c.AliasExpr.Accept(p); err != nil {
			return err
		}
	}
	if c.Codec != nil {
		builder.WriteByte(' ')
		if err := c.Codec.Accept(p); err != nil {
			return err
		}
	}
	if c.TTL != nil {
		builder.WriteByte(' ')
		if err := c.TTL.Accept(p); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		if err := c.Comment.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitColumnExpr(c *ColumnExpr) error {
	builder := p.builder
	if err := c.Expr.Accept(p); err != nil {
		return err
	}
	if c.Alias != nil {
		builder.WriteString(" AS ")
		if err := c.Alias.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitTypedPlaceholder(t *TypedPlaceholder) error {
	p.builder.WriteString("{")
	if err := t.Name.Accept(p); err != nil {
		return err
	}
	p.builder.WriteByte(':')
	if err := t.Type.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString("}")
	return nil
}

func (p *PrintVisitor) VisitColumnExprList(c *ColumnExprList) error {
	builder := p.builder
	if c.HasDistinct {
		builder.WriteString("DISTINCT ")
	}
	for i, item := range c.Items {
		if err := item.Accept(p); err != nil {
			return err
		}
		if i != len(c.Items)-1 {
			builder.WriteString(", ")
		}
	}
	if c.HasTrailingComma {
		builder.WriteString(",")
	}
	return nil
}
func (p *PrintVisitor) VisitColumnNamesExpr(c *ColumnNamesExpr) error {
	builder := p.builder
	builder.WriteByte('(')
	for i, column := range c.ColumnNames {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := column.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitComplexType(c *ComplexType) error {
	builder := p.builder
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteByte('(')
	for i, param := range c.Params {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := param.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitCompressionCodec(c *CompressionCodec) error {
	builder := p.builder
	builder.WriteString("CODEC(")
	if c.Type != nil {
		if err := c.Type.Accept(p); err != nil {
			return err
		}
		if c.TypeLevel != nil {
			builder.WriteByte('(')
			if err := c.TypeLevel.Accept(p); err != nil {
				return err
			}
			builder.WriteByte(')')
		}
		builder.WriteByte(',')
		builder.WriteByte(' ')
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.Level != nil {
		builder.WriteByte('(')
		if err := c.Level.Accept(p); err != nil {
			return err
		}
		builder.WriteByte(')')
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitConstraintExpr(c *ConstraintClause) error {
	builder := p.builder
	builder.WriteString("CONSTRAINT ")
	if err := c.Constraint.Accept(p); err != nil {
		return err
	}
	builder.WriteByte(' ')
	if c.Type != nil {
		if err := c.Type.Accept(p); err != nil {
			return err
		}
		builder.WriteByte(' ')
	}
	if err := c.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitCreateDatabase(c *CreateDatabase) error {
	builder := p.builder
	builder.WriteString(createVerb(c.IsAttach) + " DATABASE ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		// EngineExpr.String() already emits a leading " ENGINE = ...".
		if err := c.Engine.Accept(p); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		if err := c.Comment.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&c.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitCreateFunction(c *CreateFunction) error {
	builder := p.builder
	builder.WriteString("CREATE")
	if c.OrReplace {
		builder.WriteString(" OR REPLACE")
	}
	builder.WriteString(" FUNCTION ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.FunctionName.Accept(p); err != nil {
		return err
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(" AS ")
	if err := c.Params.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" -> ")
	if err := c.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitCreateLiveView(c *CreateLiveView) error {
	builder := p.builder
	builder.WriteString(createVerb(c.IsAttach) + " LIVE VIEW ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.UUID != nil {
		builder.WriteString(" ")
		if err := c.UUID.Accept(p); err != nil {
			return err
		}
	}

	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}

	if c.WithTimeout != nil {
		builder.WriteString(" ")
		if err := c.WithTimeout.Accept(p); err != nil {
			return err
		}
	}

	if c.Destination != nil {
		builder.WriteString(" ")
		if err := c.Destination.Accept(p); err != nil {
			return err
		}
	}

	if c.TableSchema != nil {
		builder.WriteString(" ")
		if err := c.TableSchema.Accept(p); err != nil {
			return err
		}
	}

	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		if err := c.SubQuery.Accept(p); err != nil {
			return err
		}
	}

	if err := p.printOutputClauses(&c.OutputClauses); err != nil {

		return err

	}
	return nil
}
func (p *PrintVisitor) VisitCreateMaterializedView(c *CreateMaterializedView) error {
	builder := p.builder
	builder.WriteString(createVerb(c.IsAttach) + " MATERIALIZED VIEW ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if c.Refresh != nil {
		builder.WriteString(" ")
		if err := c.Refresh.Accept(p); err != nil {
			return err
		}
	}
	if c.RandomizeFor != nil {
		builder.WriteString(" RANDOMIZE FOR ")
		if err := c.RandomizeFor.Accept(p); err != nil {
			return err
		}
	}
	if c.DependsOn != nil {
		builder.WriteString(" DEPENDS ON ")
		for i, dep := range c.DependsOn {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := dep.Accept(p); err != nil {
				return err
			}
		}
	}
	if c.Settings != nil {
		builder.WriteString(" ")
		if err := c.Settings.Accept(p); err != nil {
			return err
		}
	}
	if c.HasAppend {
		builder.WriteString(" APPEND")
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(p); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		builder.WriteString(" ")
		if err := c.Destination.Accept(p); err != nil {
			return err
		}
		if c.Destination.TableSchema != nil {
			builder.WriteString(" ")
			if err := c.Destination.TableSchema.Accept(p); err != nil {
				return err
			}
		}
	}
	if c.HasEmpty {
		builder.WriteString(" EMPTY")
	}
	if c.Definer != nil {
		builder.WriteString(" DEFINER = ")
		if err := c.Definer.Accept(p); err != nil {
			return err
		}
	}
	if c.SQLSecurity != "" {
		builder.WriteString(" SQL SECURITY ")
		builder.WriteString(c.SQLSecurity)
	}
	if c.Populate {
		builder.WriteString(" POPULATE")
	}
	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		if err := c.SubQuery.Accept(p); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		if err := c.Comment.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&c.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitRefreshExpr(r *RefreshExpr) error {
	builder := p.builder
	builder.WriteString("REFRESH ")
	builder.WriteString(r.Frequency)
	if r.Interval != nil {
		builder.WriteString(" ")
		if err := r.Interval.Accept(p); err != nil {
			return err
		}
	}
	if r.Offset != nil {
		builder.WriteString(" OFFSET ")
		if err := r.Offset.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitCreateDictionary(c *CreateDictionary) error {
	builder := p.builder
	builder.WriteString(createVerb(c.IsAttach))
	if c.OrReplace {
		builder.WriteString(" OR REPLACE")
	}
	builder.WriteString(" DICTIONARY ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.UUID != nil {
		builder.WriteString(" ")
		if err := c.UUID.Accept(p); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if c.Schema != nil {
		builder.WriteString(" ")
		if err := c.Schema.Accept(p); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		builder.WriteString(" ")
		if err := c.Engine.Accept(p); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		if err := c.Comment.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&c.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDictionarySchemaClause(d *DictionarySchemaClause) error {
	builder := p.builder
	builder.WriteString("(")
	for i, attr := range d.Attributes {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := attr.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(")")
	return nil
}

func (p *PrintVisitor) VisitDictionaryAttribute(d *DictionaryAttribute) error {
	builder := p.builder
	if err := d.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" ")
	if err := d.Type.Accept(p); err != nil {
		return err
	}
	if d.Default != nil {
		builder.WriteString(" DEFAULT ")
		if err := d.Default.Accept(p); err != nil {
			return err
		}
	}
	if d.Expression != nil {
		builder.WriteString(" EXPRESSION ")
		if err := d.Expression.Accept(p); err != nil {
			return err
		}
	}
	if d.Hierarchical {
		builder.WriteString(" HIERARCHICAL")
	}
	if d.Injective {
		builder.WriteString(" INJECTIVE")
	}
	if d.IsObjectId {
		builder.WriteString(" IS_OBJECT_ID")
	}
	return nil
}

func (p *PrintVisitor) VisitDictionaryEngineClause(d *DictionaryEngineClause) error {
	builder := p.builder
	start := builder.Len()
	if d.PrimaryKey != nil {
		if err := d.PrimaryKey.Accept(p); err != nil {
			return err
		}
	}
	if d.Source != nil {
		if builder.Len()-start > 0 {
			builder.WriteString(" ")
		}
		if err := d.Source.Accept(p); err != nil {
			return err
		}
	}
	if d.Lifetime != nil {
		if builder.Len()-start > 0 {
			builder.WriteString(" ")
		}
		if err := d.Lifetime.Accept(p); err != nil {
			return err
		}
	}
	if d.Layout != nil {
		if builder.Len()-start > 0 {
			builder.WriteString(" ")
		}
		if err := d.Layout.Accept(p); err != nil {
			return err
		}
	}
	if d.Range != nil {
		if builder.Len()-start > 0 {
			builder.WriteString(" ")
		}
		if err := d.Range.Accept(p); err != nil {
			return err
		}
	}
	if d.Settings != nil {
		if builder.Len()-start > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString("SETTINGS(")
		for i, item := range d.Settings.Items {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := item.Accept(p); err != nil {
				return err
			}
		}
		builder.WriteString(")")
	}
	return nil
}

func (p *PrintVisitor) VisitDictionaryPrimaryKeyClause(d *DictionaryPrimaryKeyClause) error {
	builder := p.builder
	builder.WriteString("PRIMARY KEY ")
	if err := d.Keys.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDictionarySourceClause(d *DictionarySourceClause) error {
	builder := p.builder
	builder.WriteString("SOURCE(")
	if err := d.Source.Accept(p); err != nil {
		return err
	}
	builder.WriteString("(")
	for i, arg := range d.Args {
		if i > 0 {
			builder.WriteString(" ")
		}
		if err := arg.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString("))")
	return nil
}

func (p *PrintVisitor) VisitDictionaryArgExpr(d *DictionaryArgExpr) error {
	builder := p.builder
	if err := d.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" ")
	if err := d.Value.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDictionaryLifetimeClause(d *DictionaryLifetimeClause) error {
	builder := p.builder
	builder.WriteString("LIFETIME(")
	if d.Value != nil {
		if err := d.Value.Accept(p); err != nil {
			return err
		}
	} else if d.Min != nil && d.Max != nil {
		builder.WriteString("MIN ")
		if err := d.Min.Accept(p); err != nil {
			return err
		}
		builder.WriteString(" MAX ")
		if err := d.Max.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(")")
	return nil
}

func (p *PrintVisitor) VisitDictionaryLayoutClause(d *DictionaryLayoutClause) error {
	builder := p.builder
	builder.WriteString("LAYOUT(")
	if err := d.Layout.Accept(p); err != nil {
		return err
	}
	builder.WriteString("(")
	for i, arg := range d.Args {
		if i > 0 {
			builder.WriteString(" ")
		}
		if err := arg.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString("))")
	return nil
}

func (p *PrintVisitor) VisitDictionaryRangeClause(d *DictionaryRangeClause) error {
	builder := p.builder
	builder.WriteString("RANGE(MIN ")
	if err := d.Min.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" MAX ")
	if err := d.Max.Accept(p); err != nil {
		return err
	}
	builder.WriteString(")")
	return nil
}

func (p *PrintVisitor) VisitShowExpr(s *ShowStmt) error {
	builder := p.builder
	builder.WriteString("SHOW ")
	builder.WriteString(s.ShowType)
	if s.Target != nil {
		builder.WriteByte(' ')
		if err := s.Target.Accept(p); err != nil {
			return err
		}
	}
	if s.LikeType != "" && s.LikePattern != nil {
		if s.NotLike {
			builder.WriteString(" NOT ")
		} else {
			builder.WriteByte(' ')
		}
		builder.WriteString(s.LikeType)
		builder.WriteByte(' ')
		if err := s.LikePattern.Accept(p); err != nil {
			return err
		}
	}
	if s.Limit != nil {
		builder.WriteString(" LIMIT ")
		if err := s.Limit.Accept(p); err != nil {
			return err
		}
	}
	if s.OutFile != nil {
		builder.WriteString(" INTO OUTFILE ")
		if err := s.OutFile.Accept(p); err != nil {
			return err
		}
	}
	return p.printOutputClauses(&s.OutputClauses)
}

func (p *PrintVisitor) VisitDescribeExpr(d *DescribeStmt) error {
	builder := p.builder
	builder.WriteString("DESCRIBE ")
	if d.DescribeType != "" {
		builder.WriteString(d.DescribeType)
		builder.WriteByte(' ')
	}
	if err := d.Target.Accept(p); err != nil {
		return err
	}
	return p.printOutputClauses(&d.OutputClauses)
}

func (p *PrintVisitor) VisitCreateNamedCollection(c *CreateNamedCollection) error {
	builder := p.builder
	builder.WriteString("CREATE NAMED COLLECTION ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(" AS ")
	for i, param := range c.Params {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := param.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitNamedCollectionParam(n *NamedCollectionParam) error {
	builder := p.builder
	if err := n.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" = ")
	if err := n.Value.Accept(p); err != nil {
		return err
	}
	if n.NotOverridable {
		builder.WriteString(" NOT OVERRIDABLE")
	} else if n.Overridable {
		builder.WriteString(" OVERRIDABLE")
	}
	return nil
}

func (p *PrintVisitor) VisitNamedParameterExpr(n *NamedParameterExpr) error {
	if err := n.Name.Accept(p); err != nil {
		return err
	}
	p.builder.WriteByte('=')
	return n.Value.Accept(p)
}

func (p *PrintVisitor) VisitBoolLiteral(b *BoolLiteral) error {
	p.builder.WriteString(b.Literal)
	return nil
}

func (p *PrintVisitor) VisitPath(path *Path) error {
	for i, ident := range path.Fields {
		if i > 0 {
			p.builder.WriteByte('.')
		}
		if err := ident.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitDistinctOn(s *DistinctOn) error {
	builder := p.builder
	builder.WriteString("ON (")
	for i, ident := range s.Idents {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := ident.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitTargetPairExpr(t *TargetPair) error {
	if err := t.Old.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(" TO ")
	return t.New.Accept(p)
}

func (p *PrintVisitor) VisitAuthenticationClause(a *AuthenticationClause) error {
	builder := p.builder
	if a.NotIdentified {
		builder.WriteString("NOT IDENTIFIED")
		return nil
	}
	builder.WriteString("IDENTIFIED")
	if a.AuthType != "" {
		builder.WriteString(" WITH ")
		builder.WriteString(a.AuthType)
	}
	if a.AuthValue != nil {
		builder.WriteString(" BY ")
		if err := a.AuthValue.Accept(p); err != nil {
			return err
		}
	}
	if a.LdapServer != nil {
		builder.WriteString(" WITH ldap SERVER ")
		if err := a.LdapServer.Accept(p); err != nil {
			return err
		}
	}
	if a.IsKerberos {
		builder.WriteString(" WITH kerberos")
		if a.KerberosRealm != nil && a.KerberosRealm.Literal != "" {
			builder.WriteString(" REALM ")
			if err := a.KerberosRealm.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PrintVisitor) VisitHostClause(h *HostClause) error {
	builder := p.builder
	builder.WriteString("HOST ")
	builder.WriteString(h.HostType)
	if h.HostValue != nil {
		builder.WriteString(" ")
		if err := h.HostValue.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitDefaultRoleClause(d *DefaultRoleClause) error {
	builder := p.builder
	builder.WriteString("DEFAULT ROLE ")
	if d.None {
		builder.WriteString("NONE")
	} else {
		for i, role := range d.Roles {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := role.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PrintVisitor) VisitGranteesClause(g *GranteesClause) error {
	builder := p.builder
	builder.WriteString("GRANTEES ")
	if g.Any {
		builder.WriteString("ANY")
	} else if g.None {
		builder.WriteString("NONE")
	} else {
		for i, grantee := range g.Grantees {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := grantee.Accept(p); err != nil {
				return err
			}
		}
	}
	if len(g.ExceptUsers) > 0 {
		builder.WriteString(" EXCEPT ")
		for i, except := range g.ExceptUsers {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := except.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PrintVisitor) VisitCreateUser(c *CreateUser) error {
	builder := p.builder
	builder.WriteString("CREATE USER ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if c.OrReplace {
		builder.WriteString("OR REPLACE ")
	}
	for i, userName := range c.UserNames {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := userName.Accept(p); err != nil {
			return err
		}
	}
	if c.Authentication != nil {
		builder.WriteString(" ")
		if err := c.Authentication.Accept(p); err != nil {
			return err
		}
	}
	if c.ValidUntil != nil {
		builder.WriteString(" VALID UNTIL ")
		if err := c.ValidUntil.Accept(p); err != nil {
			return err
		}
	}
	if len(c.Hosts) > 0 {
		builder.WriteString(" ")
		for i, host := range c.Hosts {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := host.Accept(p); err != nil {
				return err
			}
		}
	}
	if c.DefaultRole != nil {
		builder.WriteString(" ")
		if err := c.DefaultRole.Accept(p); err != nil {
			return err
		}
	}
	if c.DefaultDatabase != nil {
		builder.WriteString(" DEFAULT DATABASE ")
		if err := c.DefaultDatabase.Accept(p); err != nil {
			return err
		}
	} else if c.DefaultDbNone {
		builder.WriteString(" DEFAULT DATABASE NONE")
	}
	if c.Grantees != nil {
		builder.WriteString(" ")
		if err := c.Grantees.Accept(p); err != nil {
			return err
		}
	}
	if len(c.Settings) > 0 {
		builder.WriteString(" SETTINGS ")
		for i, setting := range c.Settings {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := setting.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PrintVisitor) VisitCreateRole(c *CreateRole) error {
	builder := p.builder
	builder.WriteString("CREATE ROLE ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if c.OrReplace {
		builder.WriteString("OR REPLACE ")
	}
	for i, roleName := range c.RoleNames {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := roleName.Accept(p); err != nil {
			return err
		}
	}
	if c.AccessStorageType != nil {
		builder.WriteString(" IN ")
		if err := c.AccessStorageType.Accept(p); err != nil {
			return err
		}
	}
	if len(c.Settings) > 0 {
		builder.WriteString(" SETTINGS ")
		for i, setting := range c.Settings {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := setting.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *PrintVisitor) VisitCreateTable(c *CreateTable) error {
	builder := p.builder
	builder.WriteString(createVerb(c.IsAttach))
	if c.OrReplace {
		builder.WriteString(" OR REPLACE")
	}
	if c.HasTemporary {
		builder.WriteString(" TEMPORARY")
	}
	builder.WriteString(" TABLE ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.UUID != nil {
		builder.WriteString(" ")
		if err := c.UUID.Accept(p); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		builder.WriteString(" ")
		if err := c.TableSchema.Accept(p); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(p); err != nil {
			return err
		}
	}
	for _, target := range c.TimeSeriesTargets {
		if err := target.Accept(p); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		if err := c.SubQuery.Accept(p); err != nil {
			return err
		}
	}
	if c.TableFunction != nil {
		builder.WriteString(" AS ")
		if err := c.TableFunction.Accept(p); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		if err := c.Comment.Accept(p); err != nil {
			return err
		}
	}
	if c.Settings != nil {
		builder.WriteString(" ")
		if err := c.Settings.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&c.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitCreateView(c *CreateView) error {
	builder := p.builder
	builder.WriteString(createVerb(c.IsAttach))
	if c.OrReplace {
		builder.WriteString(" OR REPLACE")
	}
	builder.WriteString(" VIEW ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if err := c.Name.Accept(p); err != nil {
		return err
	}
	if c.UUID != nil {
		builder.WriteString(" ")
		if err := c.UUID.Accept(p); err != nil {
			return err
		}
	}

	if c.OnCluster != nil {
		builder.WriteString(" ")
		if err := c.OnCluster.Accept(p); err != nil {
			return err
		}
	}

	if c.TableSchema != nil {
		builder.WriteString(" ")
		if err := c.TableSchema.Accept(p); err != nil {
			return err
		}
	}

	if c.Definer != nil {
		builder.WriteString(" DEFINER = ")
		if err := c.Definer.Accept(p); err != nil {
			return err
		}
	}
	if c.SQLSecurity != "" {
		builder.WriteString(" SQL SECURITY ")
		builder.WriteString(c.SQLSecurity)
	}

	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		if err := c.Comment.Accept(p); err != nil {
			return err
		}
	}

	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		if err := c.SubQuery.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&c.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDeduplicateExpr(d *DeduplicateClause) error {
	builder := p.builder
	builder.WriteString(" DEDUPLICATE")
	if d.By != nil {
		builder.WriteString(" BY ")
		if err := d.By.Accept(p); err != nil {
			return err
		}
	}
	if d.Except != nil {
		builder.WriteString(" EXCEPT ")
		if err := d.Except.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitDeleteFromExpr(d *DeleteClause) error {
	builder := p.builder
	builder.WriteString("DELETE FROM ")
	if err := d.Table.Accept(p); err != nil {
		return err
	}
	if d.OnCluster != nil {
		builder.WriteString(" ")
		if err := d.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if d.WhereExpr != nil {
		builder.WriteString(" WHERE ")
		if err := d.WhereExpr.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitDestinationExpr(d *DestinationClause) error {
	builder := p.builder
	builder.WriteString("TO ")
	if err := d.TableIdentifier.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDropDatabase(d *DropDatabase) error {
	builder := p.builder
	builder.WriteString(dropVerb(d.IsDetach) + " DATABASE ")
	if d.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := d.Name.Accept(p); err != nil {
		return err
	}
	if d.OnCluster != nil {
		builder.WriteString(" ")
		if err := d.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if d.Permanently {
		builder.WriteString(" PERMANENTLY")
	}
	if d.Modifier != "" {
		builder.WriteString(" " + d.Modifier)
	}
	if err := p.printOutputClauses(&d.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDropStmt(d *DropStmt) error {
	builder := p.builder
	builder.WriteString(dropVerb(d.IsDetach) + " ")
	if d.IsTemporary {
		builder.WriteString("TEMPORARY ")
	}
	builder.WriteString(d.DropTarget + " ")
	if d.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := d.Name.Accept(p); err != nil {
		return err
	}
	if d.OnCluster != nil {
		builder.WriteString(" ")
		if err := d.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if d.Permanently {
		builder.WriteString(" PERMANENTLY")
	}
	if len(d.Modifier) != 0 {
		builder.WriteString(" " + d.Modifier)
	}
	if err := p.printOutputClauses(&d.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitDropUserOrRole(d *DropUserOrRole) error {
	builder := p.builder
	builder.WriteString("DROP " + d.Target + " ")
	if d.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	for i, name := range d.Names {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := name.Accept(p); err != nil {
			return err
		}
	}
	if len(d.Modifier) != 0 {
		builder.WriteString(" " + d.Modifier)
	}
	if d.From != nil {
		builder.WriteString(" FROM ")
		if err := d.From.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitEngineExpr(e *EngineExpr) error {
	p.builder.WriteString(" ENGINE = ")
	p.builder.WriteString(e.Name)
	if e.Params != nil {
		if err := e.Params.Accept(p); err != nil {
			return err
		}
	}
	// Same clause order as EngineExpr.String().
	for _, clause := range []Expr{e.OrderBy, e.PartitionBy, e.PrimaryKey, e.SampleBy, e.TTL, e.Settings} {
		if isNilExpr(clause) {
			continue
		}
		p.builder.WriteByte(' ')
		if err := clause.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitEnumType(e *EnumType) error {
	builder := p.builder
	if err := e.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteByte('(')
	for i, enum := range e.Values {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := enum.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitEnumValue(e *EnumValue) error {
	builder := p.builder
	if err := e.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteByte('=')
	if err := e.Value.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitExplainExpr(e *ExplainStmt) error {
	builder := p.builder
	builder.WriteString("EXPLAIN")
	if e.Type != "" {
		builder.WriteByte(' ')
		builder.WriteString(e.Type)
	}
	for i, setting := range e.Settings {
		if i == 0 {
			builder.WriteByte(' ')
		} else {
			builder.WriteString(", ")
		}
		if err := setting.Accept(p); err != nil {
			return err
		}
	}
	if e.Statement != nil {
		builder.WriteByte(' ')
		if err := e.Statement.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&e.OutputClauses); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitExtractExpr(e *ExtractExpr) error {
	p.builder.WriteString("EXTRACT(")
	for i, param := range e.Parameters {
		if i > 0 {
			p.builder.WriteString(", ")
		}
		if err := param.Accept(p); err != nil {
			return err
		}
	}
	p.builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitIntervalFrom(i *IntervalFrom) error {
	if err := i.Interval.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(" FROM ")
	return i.FromExpr.Accept(p)
}

func (p *PrintVisitor) VisitFormatExpr(f *FormatClause) error {
	p.builder.WriteString("FORMAT ")
	if err := f.Format.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitFromExpr(f *FromClause) error {
	builder := p.builder
	builder.WriteString("FROM ")
	if err := f.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	if err := f.Name.Accept(p); err != nil {
		return err
	}
	if err := f.Params.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitGlobalInExpr(g *GlobalInOperation) error {
	p.builder.WriteString("GLOBAL ")
	return g.Expr.Accept(p)
}

func (p *PrintVisitor) VisitGrantPrivilegeExpr(g *GrantPrivilegeStmt) error {
	builder := p.builder
	builder.WriteString("GRANT ")
	if g.OnCluster != nil {
		builder.WriteString(" ")
		if err := g.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	for i, privilege := range g.Privileges {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := privilege.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString(" ON ")
	if err := g.On.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" TO ")
	for i, role := range g.To {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := role.Accept(p); err != nil {
			return err
		}
	}
	for _, option := range g.WithOptions {
		builder.WriteString(" WITH " + option + " OPTION")
	}

	return nil
}
func (p *PrintVisitor) VisitGroupByExpr(g *GroupByClause) error {
	p.builder.WriteString("GROUP BY ")
	p.builder.WriteString(g.AggregateType)
	// Expr is nil for `GROUP BY ALL`.
	if g.Expr != nil {
		if err := g.Expr.Accept(p); err != nil {
			return err
		}
	}
	if g.WithCube {
		p.builder.WriteString(" WITH CUBE")
	}
	if g.WithRollup {
		p.builder.WriteString(" WITH ROLLUP")
	}
	if g.WithTotals {
		p.builder.WriteString(" WITH TOTALS")
	}
	return nil
}

func (p *PrintVisitor) VisitHavingExpr(h *HavingClause) error {
	p.builder.WriteString("HAVING ")
	if err := h.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitIdent(i *Ident) error {
	if i.Param != nil {
		return i.Param.Accept(p)
	}
	switch i.QuoteType {
	case BackTicks:
		p.builder.WriteByte('`')
		p.builder.WriteString(i.Name)
		p.builder.WriteByte('`')
	case DoubleQuote:
		p.builder.WriteByte('"')
		p.builder.WriteString(i.Name)
		p.builder.WriteByte('"')
	case SingleQuote:
		p.builder.WriteByte('\'')
		p.builder.WriteString(i.Name)
		p.builder.WriteByte('\'')
	default:
		p.builder.WriteString(i.Name)
	}
	return nil
}
func (p *PrintVisitor) VisitIndexOperation(i *IndexOperation) error {
	builder := p.builder
	if err := i.Object.Accept(p); err != nil {
		return err
	}
	builder.WriteString(string(i.Operation))
	if err := i.Index.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitInsertExpr(i *InsertStmt) error {
	builder := p.builder
	builder.WriteString("INSERT INTO ")
	if i.HasTableKeyword {
		builder.WriteString("TABLE ")
	}
	if i.IsTableFunction() {
		builder.WriteString("FUNCTION ")
	}
	if err := i.Table.Accept(p); err != nil {
		return err
	}
	if i.ColumnNames != nil {
		builder.WriteString(" ")
		if err := i.ColumnNames.Accept(p); err != nil {
			return err
		}
	}
	if i.Format != nil {
		builder.WriteString(" ")
		if err := i.Format.Accept(p); err != nil {
			return err
		}
	}

	if i.SelectExpr != nil {
		builder.WriteString(" ")
		if err := i.SelectExpr.Accept(p); err != nil {
			return err
		}
	} else if len(i.Values) > 0 {
		builder.WriteString(" VALUES ")
		for j, value := range i.Values {
			if j > 0 {
				builder.WriteString(", ")
			}
			if err := value.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *PrintVisitor) VisitIntervalExpr(i *IntervalExpr) error {
	// The INTERVAL keyword is absent when the interval comes from a form such
	// as `toIntervalHour(1)` rewritten to `1 HOUR` (IntervalPos is zero).
	if i.IntervalPos != 0 {
		p.builder.WriteString("INTERVAL ")
	}
	if err := i.Expr.Accept(p); err != nil {
		return err
	}
	p.builder.WriteByte(' ')
	return i.Unit.Accept(p)
}

func (p *PrintVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	builder := p.builder
	if err := n.Expr.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" IS NOT NULL")
	return nil
}
func (p *PrintVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	builder := p.builder
	if err := n.Expr.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" IS NULL")
	return nil
}
func (p *PrintVisitor) VisitJSONOption(j *JSONOption) error {
	builder := p.builder
	if j.SkipPath != nil {
		builder.WriteString("SKIP ")
		if err := p.VisitJSONPath(j.SkipPath); err != nil {
			return err
		}
	}
	if j.SkipRegex != nil {
		builder.WriteString(" SKIP REGEXP ")
		if err := j.SkipRegex.Accept(p); err != nil {
			return err
		}
	}
	if j.MaxDynamicPaths != nil {
		builder.WriteString("max_dynamic_paths=")
		if err := j.MaxDynamicPaths.Accept(p); err != nil {
			return err
		}
	}
	if j.MaxDynamicTypes != nil {
		builder.WriteString("max_dynamic_types=")
		if err := j.MaxDynamicTypes.Accept(p); err != nil {
			return err
		}
	}
	// A type hint for a subcolumn path, e.g. `message String`.
	if j.Column != nil && j.Column.Path != nil && j.Column.Type != nil {
		if j.SkipPath != nil || j.SkipRegex != nil || j.MaxDynamicPaths != nil || j.MaxDynamicTypes != nil {
			builder.WriteByte(' ')
		}
		if err := p.VisitJSONPath(j.Column.Path); err != nil {
			return err
		}
		builder.WriteByte(' ')
		if err := j.Column.Type.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitJSONOptions(j *JSONOptions) error {
	builder := p.builder
	builder.WriteByte('(')
	// Same ordering as JSONOptions.String(): numeric options, then type-hint
	// items, then skip options, preserving relative order within each group.
	numericOptionItems := make([]*JSONOption, 0, len(j.Items))
	columnItems := make([]*JSONOption, 0, len(j.Items))
	skipOptionItems := make([]*JSONOption, 0, len(j.Items))
	for _, item := range j.Items {
		switch {
		case item.MaxDynamicPaths != nil || item.MaxDynamicTypes != nil:
			numericOptionItems = append(numericOptionItems, item)
		case item.Column != nil:
			columnItems = append(columnItems, item)
		case item.SkipPath != nil || item.SkipRegex != nil:
			skipOptionItems = append(skipOptionItems, item)
		default:
			numericOptionItems = append(numericOptionItems, item)
		}
	}
	first := true
	for _, items := range [][]*JSONOption{numericOptionItems, columnItems, skipOptionItems} {
		for _, item := range items {
			if !first {
				builder.WriteString(", ")
			}
			first = false
			if err := p.VisitJSONOption(item); err != nil {
				return err
			}
		}
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitJSONPath(j *JSONPath) error {
	builder := p.builder
	for i, ident := range j.Idents {
		if i > 0 {
			builder.WriteString(".")
		}
		if err := ident.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitJSONType(j *JSONType) error {
	if err := j.Name.Accept(p); err != nil {
		return err
	}
	if j.Options != nil {
		if err := p.VisitJSONOptions(j.Options); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitJoinConstraintExpr(j *JoinConstraintClause) error {
	builder := p.builder
	if j.On != nil {
		builder.WriteString("ON ")
		if err := j.On.Accept(p); err != nil {
			return err
		}
	} else {
		builder.WriteString("USING ")
		if err := j.Using.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitJoinExpr(j *JoinExpr) error {
	builder := p.builder
	if err := j.Left.Accept(p); err != nil {
		return err
	}
	if j.Right != nil {
		buildJoinString(builder, j.Right)
	}
	return nil
}
func (p *PrintVisitor) VisitJoinTableExpr(j *JoinTableExpr) error {
	builder := p.builder
	if err := j.Table.Accept(p); err != nil {
		return err
	}
	if j.SampleRatio != nil {
		builder.WriteByte(' ')
		if err := j.SampleRatio.Accept(p); err != nil {
			return err
		}
	}
	if j.HasFinal {
		builder.WriteString(" FINAL")
	}
	return nil
}

func (p *PrintVisitor) VisitLimitByExpr(l *LimitByClause) error {
	builder := p.builder
	if l.Limit != nil {
		if err := l.Limit.Accept(p); err != nil {
			return err
		}
	}
	if l.ByExpr != nil {
		builder.WriteString(" BY ")
		if err := l.ByExpr.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitLimitExpr(l *LimitClause) error {
	builder := p.builder
	if l.Limit != nil {
		builder.WriteString("LIMIT ")
		if err := l.Limit.Accept(p); err != nil {
			return err
		}
		if l.Offset != nil {
			builder.WriteString(" ")
		}
	}
	if l.Offset != nil {
		builder.WriteString("OFFSET ")
		if err := l.Offset.Accept(p); err != nil {
			return err
		}
	}
	if l.WithTies {
		builder.WriteString(" WITH TIES")
	}
	return nil
}
func (p *PrintVisitor) VisitMapLiteral(m *MapLiteral) error {
	builder := p.builder
	builder.WriteString("{")

	for i, value := range m.KeyValues {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := value.Key.Accept(p); err != nil {
			return err
		}
		builder.WriteString(": ")
		if err := value.Value.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteString("}")
	return nil
}

func (p *PrintVisitor) VisitNegateExpr(n *NegateExpr) error {
	p.builder.WriteString("-")
	return n.Expr.Accept(p)
}

func (p *PrintVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	if err := n.Ident.Accept(p); err != nil {
		return err
	}
	if n.DotIdent != nil {
		p.builder.WriteByte('.')
		if err := n.DotIdent.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitNestedType(n *NestedType) error {
	builder := p.builder

	if err := n.Name.Accept(p); err != nil {

		return err

	}
	builder.WriteByte('(')
	for i, column := range n.Columns {
		if err := column.Accept(p); err != nil {
			return err
		}
		if i != len(n.Columns)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitNotExpr(n *NotExpr) error {
	p.builder.WriteString("NOT ")
	return n.Expr.Accept(p)
}
func (p *PrintVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	p.builder.WriteString("NOT NULL")
	return nil
}
func (p *PrintVisitor) VisitNullLiteral(n *NullLiteral) error {
	p.builder.WriteString("NULL")
	return nil
}

func (p *PrintVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	p.builder.WriteString(n.Literal)
	return nil
}
func (p *PrintVisitor) VisitObjectParams(o *ObjectParams) error {
	if err := o.Object.Accept(p); err != nil {
		return err
	}
	return o.Params.Accept(p)
}

func (p *PrintVisitor) VisitOnExpr(o *OnClause) error {
	builder := p.builder
	builder.WriteString("ON ")
	if err := o.On.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitOperationExpr(o *OperationExpr) error {
	p.builder.WriteString(strings.ToUpper(string(o.Kind)))
	return nil
}

func (p *PrintVisitor) VisitOptimizeExpr(o *OptimizeStmt) error {
	builder := p.builder
	builder.WriteString("OPTIMIZE TABLE ")
	if err := o.Table.Accept(p); err != nil {
		return err
	}
	if o.OnCluster != nil {
		builder.WriteString(" ")
		if err := o.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if o.Partition != nil {
		builder.WriteString(" ")
		if err := o.Partition.Accept(p); err != nil {
			return err
		}
	}
	if o.HasFinal {
		builder.WriteString(" FINAL")
	}
	if o.Deduplicate != nil {
		if err := o.Deduplicate.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&o.OutputClauses); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitOrderByListExpr(o *OrderByClause) error {
	p.builder.WriteString("ORDER BY ")
	for i, item := range o.Items {
		if i > 0 {
			p.builder.WriteString(", ")
		}
		if err := item.Accept(p); err != nil {
			return err
		}
	}
	if o.Interpolate != nil {
		p.builder.WriteByte(' ')
		return o.Interpolate.Accept(p)
	}
	return nil
}
func (p *PrintVisitor) VisitOrderByExpr(o *OrderExpr) error {
	if err := o.Expr.Accept(p); err != nil {
		return err
	}
	if o.Alias != nil {
		p.builder.WriteString(" AS ")
		if err := o.Alias.Accept(p); err != nil {
			return err
		}
	}
	if o.Direction != OrderDirectionNone {
		p.builder.WriteByte(' ')
		p.builder.WriteString(string(o.Direction))
	}
	if o.Nulls != "" {
		p.builder.WriteString(" NULLS ")
		p.builder.WriteString(o.Nulls)
	}
	if o.Collate != nil {
		p.builder.WriteString(" COLLATE ")
		if err := o.Collate.Accept(p); err != nil {
			return err
		}
	}
	if o.Fill != nil {
		p.builder.WriteByte(' ')
		return o.Fill.Accept(p)
	}
	return nil
}

func (p *PrintVisitor) VisitParamExprList(f *ParamExprList) error {
	p.builder.WriteString("(")
	if err := f.Items.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(")")
	if f.ColumnArgList != nil {
		return f.ColumnArgList.Accept(p)
	}
	return nil
}
func (p *PrintVisitor) VisitPartitionByExpr(part *PartitionByClause) error {
	p.builder.WriteString("PARTITION BY ")
	return part.Expr.Accept(p)
}
func (p *PrintVisitor) VisitPartitionExpr(part *PartitionClause) error {
	p.builder.WriteString("PARTITION ")
	if part.ID != nil {
		p.builder.WriteString("ID ")
		part.ID.Accept(p)
	} else if part.All {
		p.builder.WriteString("ALL")
	} else {
		return part.Expr.Accept(p)
	}
	return nil
}

func (p *PrintVisitor) VisitPlaceHolderExpr(ph *PlaceHolder) error {
	p.builder.WriteString(ph.Type)
	return nil
}

func (p *PrintVisitor) VisitPrewhereExpr(w *PrewhereClause) error {
	p.builder.WriteString("PREWHERE ")
	return w.Expr.Accept(p)
}
func (p *PrintVisitor) VisitPrimaryKeyExpr(pkc *PrimaryKeyClause) error {
	p.builder.WriteString("PRIMARY KEY ")
	return pkc.Expr.Accept(p)
}

func (p *PrintVisitor) VisitPrivilegeExpr(pc *PrivilegeClause) error {
	builder := p.builder
	for i, keyword := range pc.Keywords {
		if i > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(keyword)
	}
	if pc.Params != nil {
		return pc.Params.Accept(p)
	}
	return nil
}
func (pv *PrintVisitor) VisitProjectionOrderBy(p *ProjectionOrderByClause) error {
	pv.builder.WriteString("ORDER BY ")
	return p.Columns.Accept(pv)
}

func (pv *PrintVisitor) VisitProjectionSelect(p *ProjectionSelectStmt) error {
	pv.builder.WriteString("(")
	if p.With != nil {
		if err := p.With.Accept(pv); err != nil {
			return err
		}
		pv.builder.WriteByte(' ')
	}
	pv.builder.WriteString("SELECT ")
	if err := p.SelectColumns.Accept(pv); err != nil {
		return err
	}
	if p.GroupBy != nil {
		pv.builder.WriteString(" ")
		if err := p.GroupBy.Accept(pv); err != nil {
			return err
		}
	}
	if p.OrderBy != nil {
		pv.builder.WriteString(" ")
		if err := p.OrderBy.Accept(pv); err != nil {
			return err
		}
	}
	pv.builder.WriteString(")")
	return nil
}

func (p *PrintVisitor) VisitPropertyType(c *PropertyType) error {
	return c.Name.Accept(p)
}

func (p *PrintVisitor) VisitQueryParam(q *QueryParam) error {
	p.builder.WriteString("{")
	if err := q.Name.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(": ")
	if err := q.Type.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString("}")
	return nil
}

func (p *PrintVisitor) VisitRatioExpr(r *RatioExpr) error {
	if err := r.Numerator.Accept(p); err != nil {
		return err
	}
	if r.Denominator != nil {
		p.builder.WriteString("/")
		if err := r.Denominator.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitRemovePropertyType(a *RemovePropertyType) error {
	p.builder.WriteString(" REMOVE ")

	return a.PropertyType.Accept(p)
}

func (p *PrintVisitor) VisitRenameStmt(r *RenameStmt) error {
	builder := p.builder
	builder.WriteString("RENAME " + r.RenameTarget + " ")
	for i, pair := range r.TargetPairList {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := pair.Old.Accept(p); err != nil {
			return err
		}
		builder.WriteString(" TO ")
		if err := pair.New.Accept(p); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		builder.WriteString(" ")
		if err := r.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&r.OutputClauses); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitRoleName(r *RoleName) error {
	builder := p.builder
	if err := r.Name.Accept(p); err != nil {
		return err
	}
	if r.Scope != nil {
		builder.WriteString("@")
		if err := r.Scope.Accept(p); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		builder.WriteByte(' ')
		if err := r.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitRoleRenamePair(r *RoleRenamePair) error {
	builder := p.builder
	if err := r.RoleName.Accept(p); err != nil {
		return err
	}
	if r.NewName != nil {
		builder.WriteString(" RENAME TO ")
		if err := r.NewName.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitRoleSetting(r *RoleSetting) error {
	builder := p.builder
	for i, settingPair := range r.SettingPairs {
		if i > 0 {
			builder.WriteString(" ")
		}
		if err := settingPair.Accept(p); err != nil {
			return err
		}
	}
	if r.Modifier != nil {
		if len(r.SettingPairs) > 0 {
			builder.WriteString(" ")
		}
		if err := r.Modifier.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitSampleByExpr(s *SampleByClause) error {
	builder := p.builder
	builder.WriteString("SAMPLE BY ")
	if err := s.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitSampleRatioExpr(s *SampleClause) error {
	builder := p.builder
	builder.WriteString("SAMPLE ")
	if err := s.Ratio.Accept(p); err != nil {
		return err
	}
	if s.Offset != nil {
		builder.WriteString(" OFFSET ")
		if err := s.Offset.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitScalarType(s *ScalarType) error {
	return s.Name.Accept(p)
}
func (p *PrintVisitor) VisitSelectItem(s *SelectItem) error {
	builder := p.builder
	if err := s.Expr.Accept(p); err != nil {
		return err
	}
	for _, modifier := range s.Modifiers {
		builder.WriteByte(' ')
		if err := modifier.Accept(p); err != nil {
			return err
		}
	}
	if s.Alias != nil {
		builder.WriteString(" AS ")
		if err := s.Alias.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitSelectQuery(s *SelectQuery) error {
	if s.Group != nil {
		p.builder.WriteString("(")
		if err := s.Group.Accept(p); err != nil {
			return err
		}
		p.builder.WriteString(")")
		return p.printSelectTail(s)
	}
	builder := p.builder
	if s.With != nil {
		builder.WriteString("WITH")
		if s.With.HasRecursive {
			builder.WriteString(" RECURSIVE")
		}
		for i, cte := range s.With.CTEs {
			builder.WriteString(" ")
			if err := cte.Accept(p); err != nil {
				return err
			}
			if i != len(s.With.CTEs)-1 {
				builder.WriteByte(',')
			}
		}
		builder.WriteString(" ")
	}
	builder.WriteString("SELECT ")
	if s.HasDistinct {
		builder.WriteString("DISTINCT ")
		if s.DistinctOn != nil {
			if err := s.DistinctOn.Accept(p); err != nil {
				return err
			}
			builder.WriteString(" ")
		}
	}
	if s.Top != nil {
		if err := s.Top.Accept(p); err != nil {
			return err
		}
		builder.WriteString(" ")
	}
	for i, selectItem := range s.SelectItems {
		if err := selectItem.Accept(p); err != nil {
			return err
		}
		if i != len(s.SelectItems)-1 {
			builder.WriteString(", ")
		}
	}
	if s.From != nil {
		builder.WriteString(" ")
		if err := s.From.Accept(p); err != nil {
			return err
		}
	}
	if s.Prewhere != nil {
		builder.WriteString(" ")
		if err := s.Prewhere.Accept(p); err != nil {
			return err
		}
	}
	if s.Where != nil {
		builder.WriteString(" ")
		if err := s.Where.Accept(p); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		builder.WriteString(" ")
		if err := s.GroupBy.Accept(p); err != nil {
			return err
		}
	}
	// WITH TOTALS without GROUP BY (a totals row over the whole query); with
	// GROUP BY it is part of the GROUP BY clause.
	if s.WithTotal {
		builder.WriteString(" WITH TOTALS")
	}
	if s.Having != nil {
		builder.WriteString(" ")
		if err := s.Having.Accept(p); err != nil {
			return err
		}
	}
	if s.Window != nil {
		builder.WriteString(" ")
		if err := s.Window.Accept(p); err != nil {
			return err
		}
	}
	if s.OrderBy != nil {
		builder.WriteString(" ")
		if err := s.OrderBy.Accept(p); err != nil {
			return err
		}
	}
	if s.LimitBy != nil {
		builder.WriteString(" ")
		if err := s.LimitBy.Accept(p); err != nil {
			return err
		}
	}
	if s.Limit != nil {
		builder.WriteString(" ")
		if err := s.Limit.Accept(p); err != nil {
			return err
		}
	}
	return p.printSelectTail(s)
}

// printSelectTail prints the clauses shared by a SELECT and a parenthesised
// group: SETTINGS, FORMAT, output SETTINGS and the set-operation
// continuation (UNION, EXCEPT, INTERSECT).
func (p *PrintVisitor) printSelectTail(s *SelectQuery) error {
	for _, clause := range []Expr{s.Settings, s.Format, s.OutputSettings} {
		if isNilExpr(clause) {
			continue
		}
		p.builder.WriteByte(' ')
		if err := clause.Accept(p); err != nil {
			return err
		}
	}
	if keyword, next := s.setOperation(); next != nil {
		p.builder.WriteByte(' ')
		p.builder.WriteString(keyword)
		p.builder.WriteByte(' ')
		return next.Accept(p)
	}
	return nil
}

func (p *PrintVisitor) VisitSetExpr(s *SetStmt) error {
	builder := p.builder
	builder.WriteString("SET ")
	for i, item := range s.Settings.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := item.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitSettingsExpr(s *SettingExpr) error {
	builder := p.builder
	if err := s.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteByte('=')
	if err := s.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitSettingPair(s *SettingPair) error {
	builder := p.builder
	if err := s.Name.Accept(p); err != nil {
		return err
	}
	if s.Value != nil {
		if s.Operation == TokenKindSingleEQ {
			builder.WriteString(string(s.Operation))
		} else {
			builder.WriteByte(' ')
		}
		if err := s.Value.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitSettingsExprList(s *SettingsClause) error {
	builder := p.builder
	builder.WriteString("SETTINGS ")
	for i, item := range s.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := item.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitStringLiteral(s *StringLiteral) error {
	p.builder.WriteString("'")
	p.builder.WriteString(s.Literal)
	p.builder.WriteString("'")
	return nil
}

func (p *PrintVisitor) VisitSubQueryExpr(s *SubQuery) error {
	if s.HasParen {
		p.builder.WriteString("(")
		if err := s.Select.Accept(p); err != nil {
			return err
		}
		p.builder.WriteString(")")
		return nil
	}
	return s.Select.Accept(p)
}

func (p *PrintVisitor) VisitSystemCtrlExpr(s *SystemCtrlExpr) error {
	builder := p.builder
	builder.WriteString(s.Command)
	builder.WriteByte(' ')
	builder.WriteString(s.Type)
	if s.OnCluster != nil {
		builder.WriteByte(' ')
		if err := s.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if s.Cluster != nil {
		builder.WriteByte(' ')
		if err := s.Cluster.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	p.builder.WriteString("DROP ")
	p.builder.WriteString(s.Type)
	return nil
}
func (p *PrintVisitor) VisitSystemFlushExpr(s *SystemFlushExpr) error {
	builder := p.builder
	builder.WriteString("FLUSH ")
	switch {
	case s.Logs:
		builder.WriteString("LOGS")
	case s.AsyncInsertQueue:
		builder.WriteString("ASYNC INSERT QUEUE")
	default:
		builder.WriteString("DISTRIBUTED")
	}
	if s.OnCluster != nil {
		builder.WriteByte(' ')
		if err := s.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	for i, table := range s.Tables {
		if i == 0 {
			builder.WriteByte(' ')
		} else {
			builder.WriteString(", ")
		}
		if err := table.Accept(p); err != nil {
			return err
		}
	}
	if s.Distributed != nil {
		builder.WriteByte(' ')
		if err := s.Distributed.Accept(p); err != nil {
			return err
		}
	}
	if s.Settings != nil {
		builder.WriteByte(' ')
		if err := s.Settings.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitSystemReloadExpr(s *SystemReloadExpr) error {
	builder := p.builder
	builder.WriteString("RELOAD ")
	builder.WriteString(s.Type)
	if s.Dictionary != nil {
		builder.WriteByte(' ')
		if err := s.Dictionary.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitSystemExpr(s *SystemStmt) error {
	p.builder.WriteString("SYSTEM ")
	return s.Expr.Accept(p)
}

func (p *PrintVisitor) VisitSystemSyncExpr(s *SystemSyncExpr) error {
	builder := p.builder
	builder.WriteString("SYNC ")
	builder.WriteString(s.Target)
	if s.CacheName != nil {
		builder.WriteByte(' ')
		if err := s.CacheName.Accept(p); err != nil {
			return err
		}
	}
	if s.OnCluster != nil {
		builder.WriteByte(' ')
		if err := s.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if s.Cluster != nil {
		builder.WriteByte(' ')
		if err := s.Cluster.Accept(p); err != nil {
			return err
		}
	}
	if s.Database != nil {
		builder.WriteByte(' ')
		if err := s.Database.Accept(p); err != nil {
			return err
		}
	}
	if s.IfExists {
		builder.WriteString(" IF EXISTS")
	}
	if s.Mode != "" {
		builder.WriteByte(' ')
		builder.WriteString(s.Mode)
	}
	for i, from := range s.From {
		if i == 0 {
			builder.WriteString(" FROM ")
		} else {
			builder.WriteString(", ")
		}
		if err := from.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitTTLExprList(t *TTLClause) error {
	p.builder.WriteString("TTL ")
	for i, item := range t.Items {
		if i > 0 {
			p.builder.WriteString(", ")
		}
		if err := item.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitTTLExpr(t *TTLExpr) error {
	builder := p.builder
	if err := t.Expr.Accept(p); err != nil {
		return err
	}
	if t.Policy != nil {
		builder.WriteString(" ")
		if err := t.Policy.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitTTLPolicy(t *TTLPolicy) error {
	builder := p.builder
	start := builder.Len()
	writeSep := func() {
		if builder.Len()-start > 0 {
			builder.WriteString(" ")
		}
	}
	if t.Item != nil {
		if err := t.Item.Accept(p); err != nil {
			return err
		}
	}
	if t.Where != nil {
		writeSep()
		if err := t.Where.Accept(p); err != nil {
			return err
		}
	}
	if t.GroupBy != nil {
		writeSep()
		if err := t.GroupBy.Accept(p); err != nil {
			return err
		}
	}
	if len(t.Assignments) > 0 {
		builder.WriteString(" SET ")
		for i, assignment := range t.Assignments {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := assignment.Accept(p); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PrintVisitor) VisitTTLPolicyRule(t *TTLPolicyRule) error {
	builder := p.builder
	if t.ToVolume != nil {
		builder.WriteString("TO VOLUME ")
		if err := t.ToVolume.Accept(p); err != nil {
			return err
		}
	} else if t.ToDisk != nil {
		builder.WriteString("TO DISK ")
		if err := t.ToDisk.Accept(p); err != nil {
			return err
		}
	} else if t.Action != nil {
		if err := t.Action.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitTTLPolicyItemAction(t *TTLPolicyRuleAction) error {
	builder := p.builder
	builder.WriteString(t.Action)
	if t.Codec != nil {
		builder.WriteString(" ")
		if err := t.Codec.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitTableArgListExpr(t *TableArgListExpr) error {
	builder := p.builder
	builder.WriteByte('(')
	for i, arg := range t.Args {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := arg.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitStreamClause(s *StreamClause) error {
	p.builder.WriteString("STREAM")
	for _, modifier := range s.Modifiers {
		p.builder.WriteByte(' ')
		p.builder.WriteString(modifier)
	}
	return nil
}

func (p *PrintVisitor) VisitTableExpr(t *TableExpr) error {
	builder := p.builder
	if err := t.Expr.Accept(p); err != nil {
		return err
	}
	if t.Alias != nil {
		builder.WriteByte(' ')
		if err := t.Alias.Accept(p); err != nil {
			return err
		}
	}
	if t.HasFinal {
		builder.WriteString(" FINAL")
	}
	if t.Stream != nil {
		builder.WriteByte(' ')
		if err := t.Stream.Accept(p); err != nil {
			return err
		}
	}
	return nil
}
func (p *PrintVisitor) VisitTableFunctionExpr(t *TableFunctionExpr) error {
	if err := t.Name.Accept(p); err != nil {
		return err
	}
	if err := t.Args.Accept(p); err != nil {
		return err
	}
	return nil
}
func (p *PrintVisitor) VisitTableIdentifier(t *TableIdentifier) error {
	if t.Database != nil {
		if err := t.Database.Accept(p); err != nil {
			return err
		}
		p.builder.WriteString(".")
	}
	return t.Table.Accept(p)
}

func (p *PrintVisitor) VisitTableIndex(a *TableIndex) error {
	p.builder.WriteString("INDEX ")
	return p.printTableIndexBody(a)
}

// printTableIndexBody prints an index definition after its INDEX keyword:
// `name expr TYPE type GRANULARITY n`.
func (p *PrintVisitor) printTableIndexBody(a *TableIndex) error {
	if err := a.Name.Accept(p); err != nil {
		return err
	}
	// As in String(): no space before a parenthesised column expression.
	if _, parenthesised := a.ColumnExpr.Expr.(*ParamExprList); !parenthesised {
		p.builder.WriteByte(' ')
	}
	if err := a.ColumnExpr.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(" TYPE ")
	if err := a.ColumnType.Accept(p); err != nil {
		return err
	}
	if a.Granularity == nil {
		return nil
	}
	p.builder.WriteString(" GRANULARITY ")
	return a.Granularity.Accept(p)
}
func (p *PrintVisitor) VisitTableProjection(t *TableProjection) error {
	builder := p.builder
	if t.IncludeProjectionKeyword {
		builder.WriteString("PROJECTION ")
	}
	if err := t.Identifier.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" ")
	if err := t.Select.Accept(p); err != nil {
		return err
	}
	if t.Settings != nil {
		builder.WriteString(" WITH SETTINGS (")
		for i, item := range t.Settings.Items {
			if i > 0 {
				builder.WriteString(", ")
			}
			if err := item.Accept(p); err != nil {
				return err
			}
		}
		builder.WriteString(")")
	}
	return nil
}

func (p *PrintVisitor) VisitTableSchemaExpr(t *TableSchemaClause) error {
	if len(t.Columns) > 0 {
		p.builder.WriteString("(")
		for i, column := range t.Columns {
			if i > 0 {
				p.builder.WriteString(", ")
			}
			if err := column.Accept(p); err != nil {
				return err
			}
		}
		p.builder.WriteByte(')')
	}
	// Callers write the separating space, so the AS forms start without one.
	if t.AliasTable != nil {
		p.builder.WriteString("AS ")
		if err := t.AliasTable.Accept(p); err != nil {
			return err
		}
	}
	if t.TableFunction != nil {
		p.builder.WriteString("AS ")
		return t.TableFunction.Accept(p)
	}
	return nil
}

func (p *PrintVisitor) VisitTimeSeriesTargetClause(t *TimeSeriesTargetClause) error {
	builder := p.builder
	// Every part repeats the target keyword, e.g.
	// ` SAMPLES INNER COLUMNS (...) SAMPLES INNER ENGINE = ...`.
	writeKeyword := func() {
		builder.WriteString(" ")
		builder.WriteString(t.Keyword)
	}
	if t.External != nil {
		writeKeyword()
		builder.WriteString(" ")
		if err := t.External.Accept(p); err != nil {
			return err
		}
	}
	if t.InnerUUID != nil {
		writeKeyword()
		builder.WriteString(" INNER ")
		if err := t.InnerUUID.Accept(p); err != nil {
			return err
		}
	}
	if t.InnerColumns != nil {
		writeKeyword()
		builder.WriteString(" INNER COLUMNS ")
		if err := t.InnerColumns.Accept(p); err != nil {
			return err
		}
	}
	if t.InnerEngine != nil {
		writeKeyword()
		if !t.EngineShorthand {
			builder.WriteString(" INNER")
		}
		// EngineExpr.String() already emits a leading " ENGINE = ...".
		if err := t.InnerEngine.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitTargetPair(t *TargetPair) error {
	if err := t.Old.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(" TO ")
	return t.New.Accept(p)
}
func (p *PrintVisitor) VisitTernaryExpr(t *TernaryOperation) error {
	builder := p.builder
	if err := t.Condition.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" ? ")
	if err := t.TrueExpr.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" : ")
	if err := t.FalseExpr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitTopExpr(t *TopClause) error {
	p.builder.WriteString("TOP ")
	if err := t.Number.Accept(p); err != nil {
		return err
	}
	if t.WithTies {
		p.builder.WriteString(" WITH TIES")
	}
	return nil
}

func (p *PrintVisitor) VisitTruncateTable(t *TruncateTable) error {
	builder := p.builder
	builder.WriteString("TRUNCATE ")
	if t.IsTemporary {
		builder.WriteString("TEMPORARY ")
	}
	builder.WriteString("TABLE ")
	if t.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	if err := t.Name.Accept(p); err != nil {
		return err
	}
	if t.OnCluster != nil {
		builder.WriteString(" ")
		if err := t.OnCluster.Accept(p); err != nil {
			return err
		}
	}
	if err := p.printOutputClauses(&t.OutputClauses); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitTypeWithParams(s *TypeWithParams) error {
	builder := p.builder
	if err := s.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteByte('(')
	for i, size := range s.Params {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := size.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitIndexTypeKwargs(s *IndexTypeKwargs) error {
	builder := p.builder
	if err := s.Name.Accept(p); err != nil {
		return err
	}
	builder.WriteByte('(')
	for i, kw := range s.Kwargs {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := kw.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitIndexTypeKwarg(k *IndexTypeKwarg) error {
	if err := k.Name.Accept(p); err != nil {
		return err
	}
	p.builder.WriteString(" = ")
	if err := k.Value.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitUUID(u *UUID) error {
	p.builder.WriteString("UUID ")
	return u.Value.Accept(p)
}

func (p *PrintVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	p.builder.WriteString(string(n.Kind))
	p.builder.WriteString(n.operandSeparator())
	return n.Expr.Accept(p)
}

func (p *PrintVisitor) VisitUseExpr(u *UseStmt) error {
	p.builder.WriteString("USE ")
	return u.Database.Accept(p)
}

func (p *PrintVisitor) VisitUsingExpr(u *UsingClause) error {
	builder := p.builder
	// Always parenthesise: without parentheses a following comma join
	// (`USING x, c`) is read as another USING column (#73).
	builder.WriteString("USING (")
	if err := u.Using.Accept(p); err != nil {
		return err
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitWhenExpr(w *WhenClause) error {
	builder := p.builder
	builder.WriteString("WHEN ")
	if err := w.When.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" THEN ")
	if err := w.Then.Accept(p); err != nil {
		return err
	}
	if w.Else != nil {
		builder.WriteString(" ELSE ")
		if err := w.Else.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitWhereExpr(w *WhereClause) error {
	builder := p.builder
	builder.WriteString("WHERE ")
	if err := w.Expr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitWindowExpr(w *WindowClause) error {
	builder := p.builder
	builder.WriteString("WINDOW ")
	for i, window := range w.Windows {
		if i > 0 {
			builder.WriteString(", ")
		}
		// WindowDefinition is not an AST node; print `name AS (...)` inline.
		if err := window.Name.Accept(p); err != nil {
			return err
		}
		builder.WriteString(" AS ")
		if err := window.Expr.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitWindowFrameParam(f *WindowFrameParam) error {
	if err := f.Param.Accept(p); err != nil {
		return err
	}
	p.builder.WriteByte(' ')
	p.builder.WriteString(f.Direction)
	return nil
}

func (p *PrintVisitor) VisitFill(f *Fill) error {
	p.builder.WriteString("WITH FILL")
	for _, part := range []struct {
		keyword string
		expr    Expr
	}{{" FROM ", f.From}, {" TO ", f.To}, {" STEP ", f.Step}, {" STALENESS ", f.Staleness}} {
		if part.expr == nil {
			continue
		}
		p.builder.WriteString(part.keyword)
		if err := part.expr.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitInterpolateItem(i *InterpolateItem) error {
	if err := i.Column.Accept(p); err != nil {
		return err
	}
	if i.Expr != nil {
		p.builder.WriteString(" AS ")
		if err := i.Expr.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitInterpolateClause(i *InterpolateClause) error {
	builder := p.builder
	builder.WriteString("INTERPOLATE")
	if len(i.Items) > 0 {
		builder.WriteString(" (")
		for idx, item := range i.Items {
			if idx > 0 {
				builder.WriteString(", ")
			}
			if err := item.Accept(p); err != nil {
				return err
			}
		}
		builder.WriteByte(')')
	}
	return nil
}

func (p *PrintVisitor) VisitWindowConditionExpr(w *WindowExpr) error {
	builder := p.builder
	builder.WriteByte('(')
	sep := false
	writeSep := func() {
		if sep {
			builder.WriteByte(' ')
		}
		sep = true
	}
	if w.WindowName != nil {
		writeSep()
		if err := w.WindowName.Accept(p); err != nil {
			return err
		}
	}
	if w.PartitionBy != nil {
		writeSep()
		if err := w.PartitionBy.Accept(p); err != nil {
			return err
		}
	}
	if w.OrderBy != nil {
		writeSep()
		if err := w.OrderBy.Accept(p); err != nil {
			return err
		}
	}
	if w.Frame != nil {
		writeSep()
		if err := w.Frame.Accept(p); err != nil {
			return err
		}
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitWindowFrameExpr(f *WindowFrameClause) error {
	builder := p.builder
	builder.WriteString(f.Type)
	builder.WriteString(" ")
	if err := f.Extend.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	p.builder.WriteString("CURRENT ROW")
	return nil
}

func (p *PrintVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	if err := f.Expr.Accept(p); err != nil {
		return err
	}
	if f.Direction != "" {
		p.builder.WriteByte(' ')
		p.builder.WriteString(f.Direction)
	}
	return nil
}

func (p *PrintVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	builder := p.builder
	if err := f.Number.Accept(p); err != nil {
		return err
	}
	builder.WriteByte(' ')
	builder.WriteString(f.Direction)
	return nil
}

func (p *PrintVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	p.builder.WriteString("UNBOUNDED ")
	p.builder.WriteString(f.Direction)
	return nil
}

func (p *PrintVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	builder := p.builder
	if err := w.Function.Accept(p); err != nil {
		return err
	}
	builder.WriteString(" OVER ")
	if err := w.OverExpr.Accept(p); err != nil {
		return err
	}
	return nil
}

func (p *PrintVisitor) VisitWithExpr(w *WithClause) error {
	builder := p.builder
	builder.WriteString("WITH ")
	if w.HasRecursive {
		builder.WriteString("RECURSIVE ")
	}
	for i, cte := range w.CTEs {
		if i > 0 {
			builder.WriteString(", ")
		}
		if err := cte.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) VisitWithTimeoutExpr(w *WithTimeoutClause) error {
	builder := p.builder
	builder.WriteString("WITH TIMEOUT")
	if w.Number != nil {
		builder.WriteByte(' ')
		if err := w.Number.Accept(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *PrintVisitor) enter(expr Expr) {}

func (p *PrintVisitor) leave(expr Expr) {}

// isNilExpr reports whether expr is nil or a typed nil pointer.
func isNilExpr(expr Expr) bool {
	if expr == nil {
		return true
	}
	v := reflect.ValueOf(expr)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
