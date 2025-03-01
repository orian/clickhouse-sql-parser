package parser

import (
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
		p.builder.WriteString(a.Expr.String())
		p.builder.WriteByte(')')
	} else {
		p.builder.WriteString(a.Expr.String())
	}
	p.builder.WriteString(" AS ")
	p.builder.WriteString(a.Alias.String())
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
		builder.WriteString(roleRenamePair.String())
	}
	if len(a.Settings) > 0 {
		builder.WriteString(" SETTINGS ")
		for i, setting := range a.Settings {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(setting.String())
		}
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTable(a *AlterTable) error {
	builder := p.builder
	builder.WriteString("ALTER TABLE ")
	builder.WriteString(a.TableIdentifier.String())
	if a.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(a.OnCluster.String())
	}
	for i, expr := range a.AlterExprs {
		builder.WriteString(" ")
		builder.WriteString(expr.String())
		if i != len(a.AlterExprs)-1 {
			builder.WriteString(",")
		}
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableAddColumn(a *AlterTableAddColumn) error {
	builder := p.builder
	builder.WriteString("ADD COLUMN ")
	builder.WriteString(a.Column.String())
	if a.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if a.After != nil {
		builder.WriteString(" AFTER ")
		builder.WriteString(a.After.String())
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableAddIndex(a *AlterTableAddIndex) error {
	builder := p.builder
	builder.WriteString("ADD ")
	builder.WriteString(a.Index.String())
	if a.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	if a.After != nil {
		builder.WriteString(" AFTER ")
		builder.WriteString(a.After.String())
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableAddProjection(a *AlterTableAddProjection) error {
	builder := p.builder
	builder.WriteString("ADD PROJECTION ")
	if a.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(a.TableProjection.String())
	if a.After != nil {
		builder.WriteString(" AFTER ")
		builder.WriteString(a.After.String())
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableAttachPartition(a *AlterTableAttachPartition) error {
	builder := p.builder
	builder.WriteString("ATTACH ")
	builder.WriteString(a.Partition.String())
	if a.From != nil {
		builder.WriteString(" FROM ")
		builder.WriteString(a.From.String())
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableClearColumn(a *AlterTableClearColumn) error {
	builder := p.builder
	builder.WriteString("CLEAR COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(a.ColumnName.String())
	if a.PartitionExpr != nil {
		builder.WriteString(" IN ")
		builder.WriteString(a.PartitionExpr.String())
	}

	return nil
}

func (p *PrintVisitor) VisitAlterTableClearIndex(a *AlterTableClearIndex) error {
	builder := p.builder
	builder.WriteString("CLEAR INDEX ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(a.IndexName.String())
	if a.PartitionExpr != nil {
		builder.WriteString(" IN ")
		builder.WriteString(a.PartitionExpr.String())
	}

	return nil
}

func (p *PrintVisitor) VisitAlterTableClearProjection(a *AlterTableClearProjection) error {
	builder := p.builder
	builder.WriteString("CLEAR PROJECTION ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(a.ProjectionName.String())
	if a.PartitionExpr != nil {
		builder.WriteString(" IN ")
		builder.WriteString(a.PartitionExpr.String())
	}

	return nil
}
func (p *PrintVisitor) VisitAlterTableDetachPartition(a *AlterTableDetachPartition) error {
	builder := p.builder
	builder.WriteString("DETACH ")
	builder.WriteString(a.Partition.String())
	if a.Settings != nil {
		builder.WriteByte(' ')
		builder.WriteString(a.Settings.String())
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableDropColumn(a *AlterTableDropColumn) error {
	builder := p.builder
	builder.WriteString("DROP COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(a.ColumnName.String())
	return nil
}

func (p *PrintVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	builder := p.builder
	builder.WriteString("DROP INDEX ")
	builder.WriteString(a.IndexName.String())
	if a.IfExists {
		builder.WriteString(" IF EXISTS")
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	builder := p.builder
	builder.WriteString("DROP ")
	if a.HasDetached {
		builder.WriteString("DETACHED ")
	}
	builder.WriteString(a.Partition.String())
	if a.Settings != nil {
		builder.WriteByte(' ')
		builder.WriteString(a.Settings.String())
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	builder := p.builder
	builder.WriteString("DROP PROJECTION ")
	builder.WriteString(a.ProjectionName.String())
	if a.IfExists {
		builder.WriteString(" IF EXISTS")
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	builder := p.builder
	builder.WriteString("FREEZE")
	if a.Partition != nil {
		builder.WriteByte(' ')
		builder.WriteString(a.Partition.String())
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
	builder.WriteString(a.IndexName.String())
	if a.Partition != nil {
		builder.WriteString(" IN ")
		builder.WriteString(a.Partition.String())
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
	builder.WriteString(a.ProjectionName.String())
	if a.Partition != nil {
		builder.WriteString(" IN ")
		builder.WriteString(a.Partition.String())
	}
	return nil
}

func (p *PrintVisitor) VisitAlterTableModifyColumn(a *AlterTableModifyColumn) error {
	builder := p.builder
	builder.WriteString("MODIFY COLUMN ")
	if a.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(a.Column.String())
	if a.RemovePropertyType != nil {
		builder.WriteString(a.RemovePropertyType.String())
	}
	return nil
}
func (p *PrintVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	builder := p.builder
	builder.WriteString("MODIFY ")
	builder.WriteString("TTL ")
	builder.WriteString(a.TTL.String())
	return nil
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
	builder.WriteString(a.OldColumnName.String())
	builder.WriteString(" TO ")
	builder.WriteString(a.NewColumnName.String())
	return nil
}

func (p *PrintVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	builder := p.builder
	builder.WriteString("REPLACE ")
	builder.WriteString(a.Partition.String())
	builder.WriteString(" FROM ")
	builder.WriteString(a.Table.String())
	return nil
}

func (p *PrintVisitor) VisitArrayJoinExpr(a *ArrayJoinClause) error {
	p.builder.WriteString(a.Type + " ARRAY JOIN " + a.Expr.String())
	return nil
}
func (p *PrintVisitor) VisitArrayParamList(a *ArrayParamList) error {
	builder := p.builder
	builder.WriteString("[")
	for i, item := range a.Items.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(item.String())
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
		builder.WriteString(value.String())
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitBetweenClause(f *BetweenClause) error {
	builder := p.builder
	builder.WriteString(f.Expr.String())
	builder.WriteString(" BETWEEN ")
	builder.WriteString(f.Between.String())
	builder.WriteString(" AND ")
	builder.WriteString(f.And.String())
	return nil
}
func (pv *PrintVisitor) VisitBinaryExpr(p *BinaryOperation) error {
	builder := pv.builder
	builder.WriteString(p.LeftExpr.String())
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
	builder.WriteString(p.RightExpr.String())
	return nil
}

func (p *PrintVisitor) VisitCTEExpr(c *CTEStmt) error {
	builder := p.builder
	builder.WriteString(c.Expr.String())
	builder.WriteString(" AS ")
	if _, isSelect := c.Alias.(*SelectQuery); isSelect {
		builder.WriteByte('(')
		builder.WriteString(c.Alias.String())
		builder.WriteByte(')')
	} else {
		builder.WriteString(c.Alias.String())
	}
	return nil
}
func (p *PrintVisitor) VisitCaseExpr(c *CaseExpr) error {
	builder := p.builder
	builder.WriteString("CASE ")
	if c.Expr != nil {
		builder.WriteString(c.Expr.String())
	}
	for _, when := range c.Whens {
		builder.WriteString(when.String())
	}
	if c.Else != nil {
		builder.WriteString(" ELSE ")
		builder.WriteString(c.Else.String())
	}
	builder.WriteString(" END")
	return nil
}

func (p *PrintVisitor) VisitCastExpr(c *CastExpr) error {
	builder := p.builder
	builder.WriteString("CAST(")
	builder.WriteString(c.Expr.String())
	if c.Separator == "," {
		builder.WriteString(", ")
	} else {
		builder.WriteString(" AS ")
	}
	builder.WriteString(c.AsType.String())
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitCheckExpr(c *CheckStmt) error {
	builder := p.builder
	builder.WriteString("CHECK TABLE ")
	builder.WriteString(c.Table.String())
	if c.Partition != nil {
		builder.WriteString(" ")
		builder.WriteString(c.Partition.String())
	}
	return nil
}
func (p *PrintVisitor) VisitOnClusterExpr(o *ClusterClause) error {
	builder := p.builder
	builder.WriteString("ON CLUSTER ")
	builder.WriteString(o.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitColumnArgList(c *ColumnArgList) error {
	builder := p.builder
	builder.WriteByte('(')
	for i, item := range c.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(item.String())
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitColumnDef(c *ColumnDef) error {
	builder := p.builder
	builder.WriteString(c.Name.String())
	if c.Type != nil {
		builder.WriteByte(' ')
		builder.WriteString(c.Type.String())
	}
	if c.NotNull != nil {
		builder.WriteString(" NOT NULL")
	} else if c.Nullable != nil {
		builder.WriteString(" NULL")
	}
	if c.DefaultExpr != nil {
		builder.WriteString(" DEFAULT ")
		builder.WriteString(c.DefaultExpr.String())
	}
	if c.MaterializedExpr != nil {
		builder.WriteString(" MATERIALIZED ")
		builder.WriteString(c.MaterializedExpr.String())
	}
	if c.AliasExpr != nil {
		builder.WriteString(" ALIAS ")
		builder.WriteString(c.AliasExpr.String())
	}
	if c.Codec != nil {
		builder.WriteByte(' ')
		builder.WriteString(c.Codec.String())
	}
	if c.TTL != nil {
		builder.WriteByte(' ')
		builder.WriteString(c.TTL.String())
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		builder.WriteString(c.Comment.String())
	}
	return nil
}
func (p *PrintVisitor) VisitColumnExpr(c *ColumnExpr) error {
	builder := p.builder
	builder.WriteString(c.Expr.String())
	if c.Alias != nil {
		builder.WriteString(" AS ")
		builder.WriteString(c.Alias.String())
	}
	return nil
}

func (p *PrintVisitor) VisitColumnExprList(c *ColumnExprList) error {
	builder := p.builder
	if c.HasDistinct {
		builder.WriteString("DISTINCT ")
	}
	for i, item := range c.Items {
		builder.WriteString(item.String())
		if i != len(c.Items)-1 {
			builder.WriteString(", ")
		}
	}
	return nil
}
func (p *PrintVisitor) VisitColumnIdentifier(c *ColumnIdentifier) error {
	if c.Database != nil {
		p.builder.WriteString(c.Database.String() + "." + c.Table.String() + "." + c.Column.String())
	} else if c.Table != nil {
		p.builder.WriteString(c.Table.String() + "." + c.Column.String())
	} else {
		p.builder.WriteString(c.Column.String())
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
		builder.WriteString(column.String())
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	p.builder.WriteString(c.Name.String())
	return nil
}
func (p *PrintVisitor) VisitComplexType(c *ComplexType) error {
	builder := p.builder
	builder.WriteString(c.Name.String())
	builder.WriteByte('(')
	for i, param := range c.Params {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(param.String())
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitCompressionCodec(c *CompressionCodec) error {
	builder := p.builder
	builder.WriteString("CODEC(")
	if c.Type != nil {
		builder.WriteString(c.Type.String())
		if c.TypeLevel != nil {
			builder.WriteByte('(')
			builder.WriteString(c.TypeLevel.String())
			builder.WriteByte(')')
		}
		builder.WriteByte(',')
		builder.WriteByte(' ')
	}
	builder.WriteString(c.Name.String())
	if c.Level != nil {
		builder.WriteByte('(')
		builder.WriteString(c.Level.String())
		builder.WriteByte(')')
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitConstraintExpr(c *ConstraintClause) error {
	builder := p.builder
	builder.WriteString(c.Constraint.String())
	builder.WriteByte(' ')
	builder.WriteString(c.Expr.String())
	return nil
}
func (p *PrintVisitor) VisitCreateDatabase(c *CreateDatabase) error {
	builder := p.builder
	builder.WriteString("CREATE DATABASE ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(c.Name.String())
	if c.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(c.OnCluster.String())
	}
	if c.Engine != nil {
		builder.WriteString(" ")
		builder.WriteString(c.Engine.String())
	}
	return nil
}

func (p *PrintVisitor) VisitCreateFunction(c *CreateFunction) error {
	builder := p.builder
	builder.WriteString("CREATE FUNCTION ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(c.FunctionName.String())
	if c.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(c.OnCluster.String())
	}
	builder.WriteString(" AS ")
	builder.WriteString(c.Params.String())
	builder.WriteString(" -> ")
	builder.WriteString(c.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitCreateLiveView(c *CreateLiveView) error {
	builder := p.builder
	builder.WriteString("CREATE LIVE VIEW ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(c.Name.String())

	if c.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(c.OnCluster.String())
	}

	if c.WithTimeout != nil {
		builder.WriteString(" ")
		builder.WriteString(c.WithTimeout.String())
	}

	if c.Destination != nil {
		builder.WriteString(" ")
		builder.WriteString(c.Destination.String())
	}

	if c.TableSchema != nil {
		builder.WriteString(" ")
		builder.WriteString(c.TableSchema.String())
	}

	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		builder.WriteString(c.SubQuery.String())
	}

	return nil
}
func (p *PrintVisitor) VisitCreateMaterializedView(c *CreateMaterializedView) error {
	builder := p.builder
	builder.WriteString("CREATE MATERIALIZED VIEW ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(c.Name.String())
	if c.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(c.OnCluster.String())
	}
	if c.Engine != nil {
		builder.WriteString(c.Engine.String())
	}
	if c.Destination != nil {
		builder.WriteString(" ")
		builder.WriteString(c.Destination.String())
		if c.Destination.TableSchema != nil {
			builder.WriteString(" ")
			builder.WriteString(c.Destination.TableSchema.String())
		}
	}
	if c.Populate {
		builder.WriteString(" POPULATE ")
	}
	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		builder.WriteString(c.SubQuery.String())
	}

	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		builder.WriteString(c.Comment.String())
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
		builder.WriteString(roleName.String())
	}
	if c.AccessStorageType != nil {
		builder.WriteString(" IN ")
		builder.WriteString(c.AccessStorageType.String())
	}
	if len(c.Settings) > 0 {
		builder.WriteString(" SETTINGS ")
		for i, setting := range c.Settings {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(setting.String())
		}
	}
	return nil
}
func (p *PrintVisitor) VisitCreateTable(c *CreateTable) error {
	builder := p.builder
	builder.WriteString("CREATE")
	if c.HasTemporary {
		builder.WriteString(" TEMPORARY")
	}
	builder.WriteString(" TABLE ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(c.Name.String())
	if c.UUID != nil {
		builder.WriteString(" ")
		builder.WriteString(c.UUID.String())
	}
	if c.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(c.OnCluster.String())
	}
	if c.TableSchema != nil {
		builder.WriteString(" ")
		builder.WriteString(c.TableSchema.String())
	}
	if c.Engine != nil {
		builder.WriteString(c.Engine.String())
	}
	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		builder.WriteString(c.SubQuery.String())
	}
	if c.Comment != nil {
		builder.WriteString(" COMMENT ")
		builder.WriteString(c.Comment.String())
	}
	return nil
}

func (p *PrintVisitor) VisitCreateView(c *CreateView) error {
	builder := p.builder
	builder.WriteString("CREATE VIEW ")
	if c.IfNotExists {
		builder.WriteString("IF NOT EXISTS ")
	}
	builder.WriteString(c.Name.String())
	if c.UUID != nil {
		builder.WriteString(" ")
		builder.WriteString(c.UUID.String())
	}

	if c.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(c.OnCluster.String())
	}

	if c.TableSchema != nil {
		builder.WriteString(" ")
		builder.WriteString(c.TableSchema.String())
	}

	if c.SubQuery != nil {
		builder.WriteString(" AS ")
		builder.WriteString(c.SubQuery.String())
	}
	return nil
}

func (p *PrintVisitor) VisitDeduplicateExpr(d *DeduplicateClause) error {
	builder := p.builder
	builder.WriteString(" DEDUPLICATE")
	if d.By != nil {
		builder.WriteString(" BY ")
		builder.WriteString(d.By.String())
	}
	if d.Except != nil {
		builder.WriteString(" EXCEPT ")
		builder.WriteString(d.Except.String())
	}
	return nil
}

func (p *PrintVisitor) VisitDeleteFromExpr(d *DeleteClause) error {
	builder := p.builder
	builder.WriteString("DELETE FROM ")
	builder.WriteString(d.Table.String())
	if d.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(d.OnCluster.String())
	}
	if d.WhereExpr != nil {
		builder.WriteString(" WHERE ")
		builder.WriteString(d.WhereExpr.String())
	}
	return nil
}
func (p *PrintVisitor) VisitDestinationExpr(d *DestinationClause) error {
	builder := p.builder
	builder.WriteString("TO ")
	builder.WriteString(d.TableIdentifier.String())
	return nil
}

func (p *PrintVisitor) VisitDropDatabase(d *DropDatabase) error {
	builder := p.builder
	builder.WriteString("DROP DATABASE ")
	if d.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(d.Name.String())
	if d.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(d.OnCluster.String())
	}
	return nil
}

func (p *PrintVisitor) VisitDropStmt(d *DropStmt) error {
	builder := p.builder
	builder.WriteString("DROP ")
	if d.IsTemporary {
		builder.WriteString("TEMPORARY ")
	}
	builder.WriteString(d.DropTarget + " ")
	if d.IfExists {
		builder.WriteString("IF EXISTS ")
	}
	builder.WriteString(d.Name.String())
	if d.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(d.OnCluster.String())
	}
	if len(d.Modifier) != 0 {
		builder.WriteString(" " + d.Modifier)
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
		builder.WriteString(name.String())
	}
	if len(d.Modifier) != 0 {
		builder.WriteString(" " + d.Modifier)
	}
	if d.From != nil {
		builder.WriteString(" FROM ")
		builder.WriteString(d.From.String())
	}
	return nil
}
func (p *PrintVisitor) VisitEngineExpr(e *EngineExpr) error {

	builder := p.builder
	builder.WriteString(" ENGINE = ")
	builder.WriteString(e.Name)
	if e.Params != nil {
		builder.WriteString(e.Params.String())
	}
	if e.PrimaryKey != nil {
		builder.WriteString(" ")
		builder.WriteString(e.PrimaryKey.String())
	}
	if e.PartitionBy != nil {
		builder.WriteString(" ")
		builder.WriteString(e.PartitionBy.String())
	}
	if e.SampleBy != nil {
		builder.WriteString(" ")
		builder.WriteString(e.SampleBy.String())
	}
	if e.TTL != nil {
		builder.WriteString(" ")
		builder.WriteString(e.TTL.String())
	}
	if e.Settings != nil {
		builder.WriteString(" ")
		builder.WriteString(e.Settings.String())
	}
	if e.OrderBy != nil {
		builder.WriteString(" ")
		builder.WriteString(e.OrderBy.String())
	}
	return nil
}
func (p *PrintVisitor) VisitEnumType(e *EnumType) error {
	builder := p.builder
	builder.WriteString(e.Name.String())
	builder.WriteByte('(')
	for i, enum := range e.Values {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(enum.String())
	}
	builder.WriteByte(')')
	return nil
}
func (p *PrintVisitor) VisitEnumValue(e *EnumValue) error {
	builder := p.builder
	builder.WriteString(e.Name.String())
	builder.WriteByte('=')
	builder.WriteString(e.Value.String())
	return nil
}

func (p *PrintVisitor) VisitExplainExpr(e *ExplainStmt) error {
	builder := p.builder
	builder.WriteString("EXPLAIN ")
	builder.WriteString(e.Type)
	builder.WriteByte(' ')
	builder.WriteString(e.Statement.String())
	return nil
}
func (p *PrintVisitor) VisitExtractExpr(e *ExtractExpr) error {
	builder := p.builder
	builder.WriteString("EXTRACT(")
	builder.WriteString(e.Interval.String())
	builder.WriteString(" FROM ")
	builder.WriteString(e.FromExpr.String())
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitFormatExpr(f *FormatClause) error {
	p.builder.WriteString("FORMAT " + f.Format.String())
	return nil
}

func (p *PrintVisitor) VisitFromExpr(f *FromClause) error {
	builder := p.builder
	builder.WriteString("FROM ")
	builder.WriteString(f.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	builder := p.builder
	builder.WriteString(f.Name.String())
	builder.WriteString(f.Params.String())
	return nil
}

func (p *PrintVisitor) VisitGlobalInExpr(g *GlobalInOperation) error {
	p.builder.WriteString("GLOBAL " + g.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitGrantPrivilegeExpr(g *GrantPrivilegeStmt) error {
	builder := p.builder
	builder.WriteString("GRANT ")
	if g.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(g.OnCluster.String())
	}
	for i, privilege := range g.Privileges {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(privilege.String())
	}
	builder.WriteString(" ON ")
	builder.WriteString(g.On.String())
	builder.WriteString(" TO ")
	for i, role := range g.To {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(role.String())
	}
	for _, option := range g.WithOptions {
		builder.WriteString(" WITH " + option + " OPTION")
	}

	return nil
}
func (p *PrintVisitor) VisitGroupByExpr(g *GroupByClause) error {
	builder := p.builder
	builder.WriteString("GROUP BY ")
	if g.AggregateType != "" {
		builder.WriteString(g.AggregateType)
	}
	builder.WriteString(g.Expr.String())
	if g.WithCube {
		builder.WriteString(" WITH CUBE")
	}
	if g.WithRollup {
		builder.WriteString(" WITH ROLLUP")
	}
	if g.WithTotals {
		builder.WriteString(" WITH TOTALS")
	}
	return nil
}

func (p *PrintVisitor) VisitHavingExpr(h *HavingClause) error {
	p.builder.WriteString("HAVING " + h.Expr.String())
	return nil
}
func (p *PrintVisitor) VisitIdent(i *Ident) error {
	if i.QuoteType == BackTicks {
		p.builder.WriteString("`" + i.Name + "`")
	} else if i.QuoteType == DoubleQuote {
		p.builder.WriteString(`"` + i.Name + `"`)
	} else {
		p.builder.WriteString(i.Name)
	}
	return nil
}
func (p *PrintVisitor) VisitIndexOperation(i *IndexOperation) error {
	builder := p.builder
	builder.WriteString(i.Object.String())
	builder.WriteString(string(i.Operation))
	builder.WriteString(i.Index.String())
	return nil
}

func (p *PrintVisitor) VisitInsertExpr(i *InsertStmt) error {
	builder := p.builder
	builder.WriteString("INSERT INTO TABLE ")
	builder.WriteString(i.Table.String())
	if i.ColumnNames != nil {
		builder.WriteString(" ")
		builder.WriteString(i.ColumnNames.String())
	}
	if i.Format != nil {
		builder.WriteString(" ")
		builder.WriteString(i.Format.String())
	}

	builder.WriteString(" ")
	if i.SelectExpr != nil {
		builder.WriteString(i.SelectExpr.String())
	} else {
		builder.WriteString("VALUES ")
		for j, value := range i.Values {
			if j > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(value.String())
		}
	}
	return nil
}
func (p *PrintVisitor) VisitIntervalExpr(i *IntervalExpr) error {
	builder := p.builder
	builder.WriteString("INTERVAL ")
	builder.WriteString(i.Expr.String())
	builder.WriteByte(' ')
	builder.WriteString(i.Unit.String())
	return nil
}

func (p *PrintVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	builder := p.builder
	builder.WriteString(n.Expr.String())
	builder.WriteString(" IS NOT NULL")
	return nil
}
func (p *PrintVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	builder := p.builder
	builder.WriteString(n.Expr.String())
	builder.WriteString(" IS NULL")
	return nil
}
func (p *PrintVisitor) VisitJSONOption(j *JSONOption) error {
	builder := p.builder
	if j.SkipPath != nil {
		builder.WriteString("SKIP ")
		builder.WriteString(j.SkipPath.String())
	}
	if j.SkipRegex != nil {
		builder.WriteString(" SKIP REGEXP ")
		builder.WriteString(j.SkipRegex.String())
	}
	return nil
}

func (p *PrintVisitor) VisitJSONOptions(j *JSONOptions) error {
	builder := p.builder
	builder.WriteByte('(')
	for i, item := range j.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(item.String())
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
		builder.WriteString(ident.String())
	}
	return nil
}

func (p *PrintVisitor) VisitJSONType(j *JSONType) error {
	builder := p.builder
	builder.WriteString(j.Name.String())
	if j.Options != nil {
		builder.WriteString(j.Options.String())
	}
	return nil
}

func (p *PrintVisitor) VisitJoinConstraintExpr(j *JoinConstraintClause) error {
	builder := p.builder
	if j.On != nil {
		builder.WriteString("ON ")
		builder.WriteString(j.On.String())
	} else {
		builder.WriteString("USING ")
		builder.WriteString(j.Using.String())
	}
	return nil
}
func (p *PrintVisitor) VisitJoinExpr(j *JoinExpr) error {
	builder := p.builder
	builder.WriteString(j.Left.String())
	if j.Right != nil {
		buildJoinString(builder, j.Right)
	}
	return nil
}
func (p *PrintVisitor) VisitJoinTableExpr(j *JoinTableExpr) error {
	builder := p.builder
	builder.WriteString(j.Table.String())
	if j.SampleRatio != nil {
		builder.WriteByte(' ')
		builder.WriteString(j.SampleRatio.String())
	}
	if j.HasFinal {
		builder.WriteString(" FINAL")
	}
	return nil
}

func (p *PrintVisitor) VisitLimitByExpr(l *LimitByClause) error {
	builder := p.builder
	if l.Limit != nil {
		builder.WriteString(l.Limit.String())
	}
	if l.ByExpr != nil {
		builder.WriteString(" BY ")
		builder.WriteString(l.ByExpr.String())
	}
	return nil
}
func (p *PrintVisitor) VisitLimitExpr(l *LimitClause) error {
	builder := p.builder
	builder.WriteString("LIMIT ")
	builder.WriteString(l.Limit.String())
	if l.Offset != nil {
		builder.WriteString(" OFFSET ")
		builder.WriteString(l.Offset.String())
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
		builder.WriteString(value.Key.String())
		builder.WriteString(": ")
		builder.WriteString(value.Value.String())
	}
	builder.WriteString("}")
	return nil
}

func (p *PrintVisitor) VisitNegateExpr(n *NegateExpr) error {
	p.builder.WriteString("-" + n.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	if n.DotIdent != nil {
		p.builder.WriteString(n.Ident.String() + "." + n.DotIdent.String())
		return nil
	}
	p.builder.WriteString(n.Ident.String())
	return nil
}

func (p *PrintVisitor) VisitNestedType(n *NestedType) error {
	builder := p.builder

	builder.WriteString(n.Name.String())
	builder.WriteByte('(')
	for i, column := range n.Columns {
		builder.WriteString(column.String())
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
	builder.WriteString(o.On.String())
	return nil
}
func (p *PrintVisitor) VisitOperationExpr(o *OperationExpr) error {
	p.builder.WriteString(strings.ToUpper(string(o.Kind)))
	return nil
}

func (p *PrintVisitor) VisitOptimizeExpr(o *OptimizeStmt) error {
	builder := p.builder
	builder.WriteString("OPTIMIZE TABLE ")
	builder.WriteString(o.Table.String())
	if o.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(o.OnCluster.String())
	}
	if o.Partition != nil {
		builder.WriteString(" ")
		builder.WriteString(o.Partition.String())
	}
	if o.HasFinal {
		builder.WriteString(" FINAL")
	}
	if o.Deduplicate != nil {
		builder.WriteString(o.Deduplicate.String())
	}
	return nil
}
func (p *PrintVisitor) VisitOrderByListExpr(o *OrderByClause) error {
	builder := p.builder
	builder.WriteString("ORDER BY ")
	for i, item := range o.Items {
		builder.WriteString(item.String())
		if i != len(o.Items)-1 {
			builder.WriteByte(',')
			builder.WriteByte(' ')
		}
	}
	return nil
}
func (p *PrintVisitor) VisitOrderByExpr(o *OrderExpr) error {
	builder := p.builder
	builder.WriteString(o.Expr.String())
	if o.Alias != nil {
		builder.WriteString(" AS ")
		builder.WriteString(o.Alias.String())
	}
	if o.Direction != OrderDirectionNone {
		builder.WriteByte(' ')
		builder.WriteString(string(o.Direction))
	}
	return nil
}

func (p *PrintVisitor) VisitParamExprList(f *ParamExprList) error {
	builder := p.builder
	builder.WriteString("(")
	for i, item := range f.Items.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(item.String())
	}
	builder.WriteString(")")
	return nil
}
func (p *PrintVisitor) VisitPartitionByExpr(part *PartitionByClause) error {
	p.builder.WriteString("PARTITION BY ")
	return part.Expr.Accept(p)
}
func (p *PrintVisitor) VisitPartitionExpr(part *PartitionClause) error {
	p.builder.WriteString("PARTITION ")
	if part.ID != nil {
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
		return pc.Accept(p)
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
	p.builder.WriteString(r.Numerator.String())
	if r.Denominator != nil {
		p.builder.WriteString("/")
		p.builder.WriteString(r.Denominator.String())
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
		builder.WriteString(pair.Old.String())
		builder.WriteString(" TO ")
		builder.WriteString(pair.New.String())
	}
	if r.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(r.OnCluster.String())
	}
	return nil
}
func (p *PrintVisitor) VisitRoleName(r *RoleName) error {
	builder := p.builder
	builder.WriteString(r.Name.String())
	if r.Scope != nil {
		builder.WriteString("@")
		builder.WriteString(r.Scope.String())
	}
	if r.OnCluster != nil {
		builder.WriteByte(' ')
		builder.WriteString(r.OnCluster.String())
	}
	return nil
}

func (p *PrintVisitor) VisitRoleRenamePair(r *RoleRenamePair) error {
	builder := p.builder
	builder.WriteString(r.RoleName.String())
	if r.NewName != nil {
		builder.WriteString(" RENAME TO ")
		builder.WriteString(r.NewName.String())
	}
	return nil
}
func (p *PrintVisitor) VisitRoleSetting(r *RoleSetting) error {
	builder := p.builder
	for i, settingPair := range r.SettingPairs {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(settingPair.String())
	}
	if r.Modifier != nil {
		if len(r.SettingPairs) > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(r.Modifier.String())
	}
	return nil
}

func (p *PrintVisitor) VisitSampleByExpr(s *SampleByClause) error {
	builder := p.builder
	builder.WriteString("SAMPLE BY ")
	builder.WriteString(s.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitSampleRatioExpr(s *SampleClause) error {
	builder := p.builder
	builder.WriteString("SAMPLE ")
	builder.WriteString(s.Ratio.String())
	if s.Offset != nil {
		builder.WriteString(" OFFSET ")
		builder.WriteString(s.Offset.String())
	}
	return nil
}
func (p *PrintVisitor) VisitScalarType(s *ScalarType) error {
	return s.Name.Accept(p)
}
func (p *PrintVisitor) VisitSelectItem(s *SelectItem) error {
	builder := p.builder
	builder.WriteString(s.Expr.String())
	for _, modifier := range s.Modifiers {
		builder.WriteByte(' ')
		builder.WriteString(modifier.String())
	}
	if s.Alias != nil {
		builder.WriteString(" AS ")
		builder.WriteString(s.Alias.String())
	}
	return nil
}

func (p *PrintVisitor) VisitSelectQuery(s *SelectQuery) error {
	builder := p.builder
	if s.With != nil {
		builder.WriteString("WITH")
		for i, cte := range s.With.CTEs {
			builder.WriteString(" ")
			builder.WriteString(cte.String())
			if i != len(s.With.CTEs)-1 {
				builder.WriteByte(',')
			}
		}
		builder.WriteString(" ")
	}
	builder.WriteString("SELECT ")
	if s.Top != nil {
		builder.WriteString(s.Top.String())
		builder.WriteString(" ")
	}
	for i, selectItem := range s.SelectItems {
		builder.WriteString(selectItem.String())
		if i != len(s.SelectItems)-1 {
			builder.WriteString(", ")
		}
	}
	if s.From != nil {
		builder.WriteString(" ")
		builder.WriteString(s.From.String())
	}
	if s.ArrayJoin != nil {
		builder.WriteString(" ")
		builder.WriteString(s.ArrayJoin.String())
	}
	if s.Window != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Window.String())
	}
	if s.Prewhere != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Prewhere.String())
	}
	if s.Where != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Where.String())
	}
	if s.GroupBy != nil {
		builder.WriteString(" ")
		builder.WriteString(s.GroupBy.String())
	}
	if s.Having != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Having.String())
	}
	if s.OrderBy != nil {
		builder.WriteString(" ")
		builder.WriteString(s.OrderBy.String())
	}
	if s.LimitBy != nil {
		builder.WriteString(" ")
		builder.WriteString(s.LimitBy.String())
	}
	if s.Limit != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Limit.String())
	}
	if s.Settings != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Settings.String())
	}
	if s.Format != nil {
		builder.WriteString(" ")
		builder.WriteString(s.Format.String())
	}
	if s.UnionAll != nil {
		builder.WriteString(" UNION ALL ")
		builder.WriteString(s.UnionAll.String())
	} else if s.UnionDistinct != nil {
		builder.WriteString(" UNION DISTINCT ")
		builder.WriteString(s.UnionDistinct.String())
	} else if s.Except != nil {
		builder.WriteString(" EXCEPT ")
		builder.WriteString(s.Except.String())
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
		builder.WriteString(item.String())
	}
	return nil
}
func (p *PrintVisitor) VisitSettingsExpr(s *SettingExprList) error {
	builder := p.builder
	builder.WriteString(s.Name.String())
	builder.WriteByte('=')
	builder.WriteString(s.Expr.String())
	return nil
}
func (p *PrintVisitor) VisitSettingPair(s *SettingPair) error {
	builder := p.builder
	builder.WriteString(s.Name.String())
	if s.Value != nil {
		if s.Operation == TokenKindSingleEQ {
			builder.WriteString(string(s.Operation))
		} else {
			builder.WriteByte(' ')
		}
		builder.WriteString(s.Value.String())
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
		builder.WriteString(item.String())
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
		p.builder.WriteString(s.Select.String())
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
	if s.Cluster != nil {
		builder.WriteByte(' ')
		builder.WriteString(s.Cluster.String())
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
	if s.Logs {
		builder.WriteString("LOGS")
	} else {
		builder.WriteString(s.Distributed.String())
	}
	return nil
}

func (p *PrintVisitor) VisitSystemReloadExpr(s *SystemReloadExpr) error {
	builder := p.builder
	builder.WriteString("RELOAD ")
	builder.WriteString(s.Type)
	if s.Dictionary != nil {
		builder.WriteByte(' ')
		builder.WriteString(s.Dictionary.String())
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
	builder.WriteString(s.Cluster.String())
	return nil
}
func (p *PrintVisitor) VisitTTLExprList(t *TTLClause) error {
	builder := p.builder
	builder.WriteString("TTL ")
	for i, item := range t.Items {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(item.String())
	}
	return nil
}
func (p *PrintVisitor) VisitTTLExpr(t *TTLExpr) error {
	builder := p.builder
	builder.WriteString(t.Expr.String())
	return nil
}
func (p *PrintVisitor) VisitTableArgListExpr(t *TableArgListExpr) error {
	builder := p.builder
	builder.WriteByte('(')
	for i, arg := range t.Args {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(arg.String())
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitTableExpr(t *TableExpr) error {
	builder := p.builder
	builder.WriteString(t.Expr.String())
	if t.Alias != nil {
		builder.WriteByte(' ')
		builder.WriteString(t.Alias.String())
	}
	if t.HasFinal {
		builder.WriteString(" FINAL")
	}
	return nil
}
func (p *PrintVisitor) VisitTableFunctionExpr(t *TableFunctionExpr) error {
	builder := p.builder
	builder.WriteString(t.Name.String())
	builder.WriteString(t.Args.String())
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
	builder := p.builder
	builder.WriteString("INDEX")
	builder.WriteByte(' ')
	builder.WriteString(a.Name.String())

	if _, ok := a.ColumnExpr.Expr.(*Ident); ok {
		builder.WriteByte(' ')
	}
	builder.WriteString(a.ColumnExpr.String())
	builder.WriteByte(' ')
	builder.WriteString("TYPE")
	builder.WriteByte(' ')
	builder.WriteString(a.ColumnType.String())
	builder.WriteByte(' ')
	builder.WriteString("GRANULARITY")
	builder.WriteByte(' ')
	builder.WriteString(a.Granularity.String())
	return nil
}
func (p *PrintVisitor) VisitTableProjection(t *TableProjection) error {
	builder := p.builder
	builder.WriteString(t.Identifier.String())
	builder.WriteString(" ")
	builder.WriteString(t.Select.String())
	return nil
}

func (p *PrintVisitor) VisitTableSchemaExpr(t *TableSchemaClause) error {
	builder := p.builder
	if len(t.Columns) > 0 {
		builder.WriteString("(")
		for i, column := range t.Columns {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(column.String())
		}
		builder.WriteByte(')')
	}
	if t.AliasTable != nil {
		builder.WriteString(" AS ")
		builder.WriteString(t.AliasTable.String())
	}
	if t.TableFunction != nil {
		builder.WriteByte(' ')
		builder.WriteString(t.TableFunction.String())
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
	builder.WriteString(t.Condition.String())
	builder.WriteString(" ? ")
	builder.WriteString(t.TrueExpr.String())
	builder.WriteString(" : ")
	builder.WriteString(t.FalseExpr.String())
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
	builder.WriteString(t.Name.String())
	if t.OnCluster != nil {
		builder.WriteString(" ")
		builder.WriteString(t.OnCluster.String())
	}
	return nil
}

func (p *PrintVisitor) VisitTypeWithParams(s *TypeWithParams) error {
	builder := p.builder
	builder.WriteString(s.Name.String())
	builder.WriteByte('(')
	for i, size := range s.Params {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(size.String())
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitUUID(u *UUID) error {
	p.builder.WriteString("UUID ")
	return u.Value.Accept(p)
}

func (p *PrintVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	p.builder.WriteString(string(n.Kind))
	p.builder.WriteByte(' ')
	return n.Expr.Accept(p)
}

func (p *PrintVisitor) VisitUseExpr(u *UseStmt) error {
	p.builder.WriteString("USE ")
	return u.Database.Accept(p)
}

func (p *PrintVisitor) VisitUsingExpr(u *UsingClause) error {
	builder := p.builder
	builder.WriteString("USING ")
	builder.WriteString(u.Using.String())
	return nil
}

func (p *PrintVisitor) VisitWhenExpr(w *WhenClause) error {
	builder := p.builder
	builder.WriteString("WHEN ")
	builder.WriteString(w.When.String())
	builder.WriteString(" THEN ")
	builder.WriteString(w.Then.String())
	if w.Else != nil {
		builder.WriteString(" ELSE ")
		builder.WriteString(w.Else.String())
	}
	return nil
}

func (p *PrintVisitor) VisitWhereExpr(w *WhereClause) error {
	builder := p.builder
	builder.WriteString("WHERE ")
	builder.WriteString(w.Expr.String())
	return nil
}

func (p *PrintVisitor) VisitWindowExpr(w *WindowClause) error {
	builder := p.builder
	builder.WriteString("WINDOW ")
	builder.WriteString(w.Name.String())
	builder.WriteString(" ")
	builder.WriteString(w.WindowExpr.String())
	return nil
}

func (p *PrintVisitor) VisitWindowConditionExpr(w *WindowExpr) error {
	builder := p.builder
	builder.WriteByte('(')
	if w.PartitionBy != nil {
		builder.WriteString(" ")
		builder.WriteString(w.PartitionBy.String())
	}
	if w.OrderBy != nil {
		builder.WriteString(" ")
		builder.WriteString(w.OrderBy.String())
	}
	if w.Frame != nil {
		builder.WriteString(" ")
		builder.WriteString(w.Frame.String())
	}
	builder.WriteByte(')')
	return nil
}

func (p *PrintVisitor) VisitWindowFrameExpr(f *WindowFrameClause) error {
	builder := p.builder
	builder.WriteString(f.Type)
	builder.WriteString(" ")
	builder.WriteString(f.Extend.String())
	return nil
}

func (p *PrintVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	p.builder.WriteString("CURRENT ROW")
	return nil
}

func (p *PrintVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	return f.Expr.Accept(p)
}

func (p *PrintVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	builder := p.builder
	builder.WriteString(f.Number.String())
	builder.WriteByte(' ')
	builder.WriteString(f.Direction)
	return nil
}

func (p *PrintVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	p.builder.WriteString(f.Direction)
	p.builder.WriteString(" UNBOUNDED")
	return nil
}

func (p *PrintVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	builder := p.builder
	builder.WriteString(w.Function.String())
	builder.WriteString(" OVER ")
	builder.WriteString(w.OverExpr.String())
	return nil
}

func (p *PrintVisitor) VisitWithExpr(w *WithClause) error {
	builder := p.builder
	builder.WriteString("WITH ")
	for i, cte := range w.CTEs {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(cte.String())
	}
	return nil
}

func (p *PrintVisitor) VisitWithTimeoutExpr(w *WithTimeoutClause) error {
	builder := p.builder
	builder.WriteString("WITH TIMEOUT ")
	builder.WriteString(w.Number.String())
	return nil
}

func (p *PrintVisitor) enter(expr Expr) {}

func (p *PrintVisitor) leave(expr Expr) {}
