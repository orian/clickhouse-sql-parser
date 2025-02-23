package parser

import (
	"strings"
)

type MyDefaultVisitor struct{}

func (visitor MyDefaultVisitor) VisitAliasExpr(a *AliasExpr) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor); err != nil {
		return err
	}
	if err := a.Alias.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterRole(a *AlterRole) error {
	visitor.enter(a)
	defer visitor.leave(a)
	for _, roleRenamePair := range a.RoleRenamePairs {
		if err := roleRenamePair.Accept(visitor); err != nil {
			return err
		}
	}
	for _, setting := range a.Settings {
		if err := setting.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTable(a *AlterTable) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if a.OnCluster != nil {
		if err := a.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}

	for _, expr := range a.AlterExprs {
		if err := expr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableAddColumn(a *AlterTableAddColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Column.Accept(visitor); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableAddIndex(a *AlterTableAddIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Index.Accept(visitor); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableAddProjection(a *AlterTableAddProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.TableProjection.Accept(visitor); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableAttachPartition(a *AlterTableAttachPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor); err != nil {
		return err
	}
	if a.From != nil {
		if err := a.From.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableClearColumn(a *AlterTableClearColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableClearIndex(a *AlterTableClearIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableClearProjection(a *AlterTableClearProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableDetachPartition(a *AlterTableDetachPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor); err != nil {
		return err
	}
	if a.Settings != nil {
		if err := a.Settings.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableDropColumn(a *AlterTableDropColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableMaterializeIndex(a *AlterTableMaterializeIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor); err != nil {
		return err
	}
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableMaterializeProjection(a *AlterTableMaterializeProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor); err != nil {
		return err
	}
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableModifyColumn(a *AlterTableModifyColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Column.Accept(visitor); err != nil {
		return err
	}
	if a.RemovePropertyType != nil {
		if err := a.RemovePropertyType.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.TTL.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableRenameColumn(a *AlterTableRenameColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.OldColumnName.Accept(visitor); err != nil {
		return err
	}
	if err := a.NewColumnName.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor); err != nil {
		return err
	}
	if err := a.Table.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitArrayJoinClause(a *ArrayJoinClause) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitArrayParamList(a *ArrayParamList) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Items.Accept(visitor); err != nil {
		return err
	}
	return nil
}

// nolint: funlen

// START, STOP
// REPLICATED, DISTRIBUTED

func (visitor MyDefaultVisitor) VisitAssignmentValues(v *AssignmentValues) error {
	visitor.enter(v)
	defer visitor.leave(v)
	for _, value := range v.Values {
		if err := value.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitBetweenClause(f *BetweenClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Between.Accept(visitor); err != nil {
		return err
	}
	if err := f.And.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitBinaryOperation(p *BinaryOperation) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.LeftExpr.Accept(visitor); err != nil {
		return err
	}
	if err := p.RightExpr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCTEStmt(c *CTEStmt) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor); err != nil {
		return err
	}
	if err := c.Alias.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitCaseExpr(c *CaseExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.Expr != nil {
		if err := c.Expr.Accept(visitor); err != nil {
			return err
		}
	}
	for _, when := range c.Whens {
		if err := when.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Else != nil {
		if err := c.Else.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCastExpr(c *CastExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor); err != nil {
		return err
	}
	if err := c.AsType.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCheckStmt(c *CheckStmt) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Table.Accept(visitor); err != nil {
		return err
	}
	if c.Partition != nil {
		if err := c.Partition.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitClusterClause(o *ClusterClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitColumnArgList(c *ColumnArgList) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitColumnDef(c *ColumnDef) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	if c.Type != nil {
		if err := c.Type.Accept(visitor); err != nil {
			return err
		}
	}
	if c.NotNull != nil {
		if err := c.NotNull.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Nullable != nil {
		if err := c.Nullable.Accept(visitor); err != nil {
			return err
		}
	}
	if c.DefaultExpr != nil {
		if err := c.DefaultExpr.Accept(visitor); err != nil {
			return err
		}
	}
	if c.MaterializedExpr != nil {
		if err := c.MaterializedExpr.Accept(visitor); err != nil {
			return err
		}
	}
	if c.AliasExpr != nil {
		if err := c.AliasExpr.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Codec != nil {
		if err := c.Codec.Accept(visitor); err != nil {
			return err
		}
	}
	if c.TTL != nil {
		if err := c.TTL.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		if err := c.Comment.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitColumnExpr(c *ColumnExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor); err != nil {
		return err
	}
	if c.Alias != nil {
		if err := c.Alias.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitColumnExprList(c *ColumnExprList) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitColumnIdentifier(c *ColumnIdentifier) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.Database != nil {
		if err := c.Database.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Table != nil {
		if err := c.Table.Accept(visitor); err != nil {
			return err
		}
	}
	if err := c.Column.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitColumnNamesExpr(c *ColumnNamesExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for i := range c.ColumnNames {
		if err := c.ColumnNames[i].Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitComplexType(c *ComplexType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	for _, param := range c.Params {
		if err := param.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCompressionCodec(c *CompressionCodec) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Type.Accept(visitor); err != nil {
		return err
	}
	if c.TypeLevel != nil {
		if err := c.TypeLevel.Accept(visitor); err != nil {
			return err
		}
	}
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	if c.Level != nil {
		if err := c.Level.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitConstraintClause(c *ConstraintClause) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Constraint.Accept(visitor); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitCreateDatabase(c *CreateDatabase) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCreateFunction(c *CreateFunction) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.FunctionName.Accept(visitor); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if err := c.Params.Accept(visitor); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCreateLiveView(c *CreateLiveView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		if err := c.Destination.Accept(visitor); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor); err != nil {
			return err
		}
	}
	if c.WithTimeout != nil {
		if err := c.WithTimeout.Accept(visitor); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitCreateMaterializedView(c *CreateMaterializedView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		if err := c.Destination.Accept(visitor); err != nil {
			return err
		}
		if c.Destination.TableSchema != nil {
			if err := c.Destination.TableSchema.Accept(visitor); err != nil {
				return err
			}
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCreateRole(c *CreateRole) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, roleName := range c.RoleNames {
		if err := roleName.Accept(visitor); err != nil {
			return err
		}
	}
	if c.AccessStorageType != nil {
		if err := c.AccessStorageType.Accept(visitor); err != nil {
			return err
		}
	}
	for _, setting := range c.Settings {
		if err := setting.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitCreateTable(c *CreateTable) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitCreateView(c *CreateView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitDeduplicateClause(d *DeduplicateClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if d.By != nil {
		if err := d.By.Accept(visitor); err != nil {
			return err
		}
	}
	if d.Except != nil {
		if err := d.Except.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitDeleteClause(d *DeleteClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Table.Accept(visitor); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if d.WhereExpr != nil {
		if err := d.WhereExpr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitDestinationClause(d *DestinationClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.TableIdentifier.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitDropDatabase(d *DropDatabase) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Name.Accept(visitor); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitDropStmt(d *DropStmt) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Name.Accept(visitor); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	return nil

}

func (visitor MyDefaultVisitor) VisitDropUserOrRole(d *DropUserOrRole) error {
	visitor.enter(d)
	defer visitor.leave(d)
	for _, name := range d.Names {
		if err := name.Accept(visitor); err != nil {
			return err
		}
	}
	if d.From != nil {
		if err := d.From.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitEngineExpr(e *EngineExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if e.Params != nil {
		if err := e.Params.Accept(visitor); err != nil {
			return err
		}
	}
	if e.PrimaryKey != nil {
		if err := e.PrimaryKey.Accept(visitor); err != nil {
			return err
		}
	}
	if e.PartitionBy != nil {
		if err := e.PartitionBy.Accept(visitor); err != nil {
			return err
		}
	}
	if e.SampleBy != nil {
		if err := e.SampleBy.Accept(visitor); err != nil {
			return err
		}
	}
	if e.TTL != nil {
		if err := e.TTL.Accept(visitor); err != nil {
			return err
		}
	}
	if e.Settings != nil {
		if err := e.Settings.Accept(visitor); err != nil {
			return err
		}
	}
	if e.OrderBy != nil {
		if err := e.OrderBy.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitEnumType(e *EnumType) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Name.Accept(visitor); err != nil {
		return err
	}
	for i := range e.Values {
		if err := e.Values[i].Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitEnumValue(e *EnumValue) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Name.Accept(visitor); err != nil {
		return err
	}
	if err := e.Value.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitExplainStmt(e *ExplainStmt) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Statement.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitExtractExpr(e *ExtractExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.FromExpr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitFormatClause(f *FormatClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Format.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitFromClause(f *FromClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Name.Accept(visitor); err != nil {
		return err
	}
	if err := f.Params.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitGlobalInOperation(g *GlobalInOperation) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitGrantPrivilegeStmt(g *GrantPrivilegeStmt) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if g.OnCluster != nil {
		if err := g.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	for _, privilege := range g.Privileges {
		if err := privilege.Accept(visitor); err != nil {
			return err
		}
	}
	if err := g.On.Accept(visitor); err != nil {
		return err
	}
	for _, role := range g.To {
		if err := role.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitGroupByClause(g *GroupByClause) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitHavingClause(h *HavingClause) error {
	visitor.enter(h)
	defer visitor.leave(h)
	if err := h.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitIdent(i *Ident) error {
	visitor.enter(i)
	defer visitor.leave(i)
	return nil
}
func (visitor MyDefaultVisitor) VisitIndexOperation(i *IndexOperation) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if err := i.Object.Accept(visitor); err != nil {
		return err
	}
	if err := i.Index.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitInsertStmt(i *InsertStmt) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if i.Format != nil {
		if err := i.Format.Accept(visitor); err != nil {
			return err
		}
	}
	if err := i.Table.Accept(visitor); err != nil {
		return err
	}
	if i.ColumnNames != nil {
		if err := i.ColumnNames.Accept(visitor); err != nil {
			return err
		}
	}
	for _, value := range i.Values {
		if err := value.Accept(visitor); err != nil {
			return err
		}
	}
	if i.SelectExpr != nil {
		if err := i.SelectExpr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitIntervalExpr(i *IntervalExpr) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if err := i.Expr.Accept(visitor); err != nil {
		return err
	}
	if err := i.Unit.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitJSONType(j *JSONType) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitJoinConstraintClause(j *JoinConstraintClause) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if j.On != nil {
		if err := j.On.Accept(visitor); err != nil {
			return err
		}
	}
	if j.Using != nil {
		if err := j.Using.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitJoinExpr(j *JoinExpr) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Left.Accept(visitor); err != nil {
		return err
	}
	if j.Right != nil {
		if err := j.Right.Accept(visitor); err != nil {
			return err
		}
	}
	if j.Constraints != nil {
		if err := j.Constraints.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitJoinTableExpr(j *JoinTableExpr) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Table.Accept(visitor); err != nil {
		return err
	}
	if j.SampleRatio != nil {
		return j.SampleRatio.Accept(visitor)
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitLimitByClause(l *LimitByClause) error {
	visitor.enter(l)
	defer visitor.leave(l)
	if l.Limit != nil {
		if err := l.Limit.Accept(visitor); err != nil {
			return err
		}
	}
	if l.ByExpr != nil {
		if err := l.ByExpr.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitLimitClause(l *LimitClause) error {
	visitor.enter(l)
	defer visitor.leave(l)
	if err := l.Limit.Accept(visitor); err != nil {
		return err
	}
	if l.Offset != nil {
		if err := l.Offset.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitMapLiteral(m *MapLiteral) error {
	visitor.enter(m)
	defer visitor.leave(m)
	for _, kv := range m.KeyValues {
		if err := kv.Key.Accept(visitor); err != nil {
			return err
		}
		if err := kv.Value.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitNegateExpr(n *NegateExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Ident.Accept(visitor); err != nil {
		return err
	}
	if n.DotIdent != nil {
		if err := n.DotIdent.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitNestedType(n *NestedType) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Name.Accept(visitor); err != nil {
		return err
	}
	for _, column := range n.Columns {
		if err := column.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitNotExpr(n *NotExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.NullLiteral.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitNullLiteral(n *NullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}

func (visitor MyDefaultVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}
func (visitor MyDefaultVisitor) VisitObjectParams(o *ObjectParams) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Object.Accept(visitor); err != nil {
		return err
	}
	if err := o.Params.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitOnClause(o *OnClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.On.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitOperationExpr(o *OperationExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	return nil
}

func (visitor MyDefaultVisitor) VisitOptimizeStmt(o *OptimizeStmt) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Table.Accept(visitor); err != nil {
		return err
	}
	if o.OnCluster != nil {
		if err := o.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	if o.Partition != nil {
		if err := o.Partition.Accept(visitor); err != nil {
			return err
		}
	}
	if o.Deduplicate != nil {
		if err := o.Deduplicate.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitOrderByClause(o *OrderByClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	for _, item := range o.Items {
		if err := item.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitOrderExpr(o *OrderExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Expr.Accept(visitor); err != nil {
		return err
	}
	if o.Alias != nil {
		if err := o.Alias.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitParamExprList(f *ParamExprList) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Items.Accept(visitor); err != nil {
		return err
	}
	if f.ColumnArgList != nil {
		if err := f.ColumnArgList.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitPartitionByClause(p *PartitionByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitPartitionClause(p *PartitionClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.Expr != nil {
		if err := p.Expr.Accept(visitor); err != nil {
			return err
		}
	}
	if p.ID != nil {
		if err := p.ID.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitPlaceHolder(p *PlaceHolder) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

func (visitor MyDefaultVisitor) VisitPrewhereClause(w *PrewhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitPrimaryKeyClause(p *PrimaryKeyClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitPrivilegeClause(p *PrivilegeClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.Params != nil {
		if err := p.Params.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitProjectionOrderByClause(p *ProjectionOrderByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

func (visitor MyDefaultVisitor) VisitProjectionSelectStmt(p *ProjectionSelectStmt) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.With != nil {
		if err := p.With.Accept(visitor); err != nil {
			return err
		}
	}
	if err := p.SelectColumns.Accept(visitor); err != nil {
		return err
	}
	if p.GroupBy != nil {
		if err := p.GroupBy.Accept(visitor); err != nil {
			return err
		}
	}
	if p.OrderBy != nil {
		if err := p.OrderBy.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitPropertyType(c *PropertyType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitQueryParam(q *QueryParam) error {
	visitor.enter(q)
	defer visitor.leave(q)
	if err := q.Name.Accept(visitor); err != nil {
		return err
	}
	if err := q.Type.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitRatioExpr(r *RatioExpr) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.Numerator.Accept(visitor); err != nil {
		return err
	}
	if r.Denominator != nil {
		if err := r.Denominator.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitRemovePropertyType(a *RemovePropertyType) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.PropertyType.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitRenameStmt(r *RenameStmt) error {
	visitor.enter(r)
	defer visitor.leave(r)
	for _, pair := range r.TargetPairList {
		if err := pair.Old.Accept(visitor); err != nil {
			return err
		}
		if err := pair.New.Accept(visitor); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		if err := r.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitRoleName(r *RoleName) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.Name.Accept(visitor); err != nil {
		return err
	}
	if r.Scope != nil {
		if err := r.Scope.Accept(visitor); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		if err := r.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitRoleRenamePair(r *RoleRenamePair) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.RoleName.Accept(visitor); err != nil {
		return err
	}
	if r.NewName != nil {
		if err := r.NewName.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitRoleSetting(r *RoleSetting) error {
	visitor.enter(r)
	defer visitor.leave(r)
	for _, settingPair := range r.SettingPairs {
		if err := settingPair.Accept(visitor); err != nil {
			return err
		}
	}
	if r.Modifier != nil {
		if err := r.Modifier.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSampleByClause(s *SampleByClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSampleClause(s *SampleClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Ratio.Accept(visitor); err != nil {
		return err
	}
	if s.Offset != nil {
		if err := s.Offset.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitScalarType(s *ScalarType) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitSelectItem(s *SelectItem) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor); err != nil {
		return err
	}
	for _, modifier := range s.Modifiers {
		if err := modifier.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Alias != nil {
		if err := s.Alias.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSelectQuery(s *SelectQuery) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.With != nil {
		if err := s.With.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Top != nil {
		if err := s.Top.Accept(visitor); err != nil {
			return err
		}
	}
	if s.SelectItems != nil {
		for _, item := range s.SelectItems {
			if err := item.Accept(visitor); err != nil {
				return err
			}
		}
	}
	if s.From != nil {
		if err := s.From.Accept(visitor); err != nil {
			return err
		}
	}
	if s.ArrayJoin != nil {
		if err := s.ArrayJoin.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Window != nil {
		if err := s.Window.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Prewhere != nil {
		if err := s.Prewhere.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Where != nil {
		if err := s.Where.Accept(visitor); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Having != nil {
		if err := s.Having.Accept(visitor); err != nil {
			return err
		}
	}
	if s.OrderBy != nil {
		if err := s.OrderBy.Accept(visitor); err != nil {
			return err
		}
	}
	if s.LimitBy != nil {
		if err := s.LimitBy.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Limit != nil {
		if err := s.Limit.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Format != nil {
		if err := s.Format.Accept(visitor); err != nil {
			return err
		}
	}
	if s.UnionAll != nil {
		if err := s.UnionAll.Accept(visitor); err != nil {
			return err
		}
	}
	if s.UnionDistinct != nil {
		if err := s.UnionDistinct.Accept(visitor); err != nil {
			return err
		}
	}
	if s.Except != nil {
		if err := s.Except.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSetStmt(s *SetStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Settings.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitSettingExprList(s *SettingExprList) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor); err != nil {
		return err
	}
	if err := s.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitSettingPair(s *SettingPair) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor); err != nil {
		return err
	}
	if s.Value != nil {
		if err := s.Value.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSettingsClause(s *SettingsClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	for _, item := range s.Items {
		if err := item.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitStringLiteral(s *StringLiteral) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}

func (visitor MyDefaultVisitor) VisitSubQuery(s *SubQuery) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Select != nil {
		if err := s.Select.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSystemCtrlExpr(s *SystemCtrlExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Cluster != nil {
		if err := s.Cluster.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}
func (visitor MyDefaultVisitor) VisitSystemFlushExpr(s *SystemFlushExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Distributed != nil {
		if err := s.Distributed.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSystemReloadExpr(s *SystemReloadExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Dictionary != nil {
		if err := s.Dictionary.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitSystemStmt(s *SystemStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitSystemSyncExpr(s *SystemSyncExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Cluster.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTTLClause(t *TTLClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, item := range t.Items {
		if err := item.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTTLExpr(t *TTLExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTableArgListExpr(t *TableArgListExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, arg := range t.Args {
		if err := arg.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitTableExpr(t *TableExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Expr.Accept(visitor); err != nil {
		return err
	}
	if t.Alias != nil {
		if err := t.Alias.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTableFunctionExpr(t *TableFunctionExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Name.Accept(visitor); err != nil {
		return err
	}
	if err := t.Args.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTableIdentifier(t *TableIdentifier) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if t.Database != nil {
		if err := t.Database.Accept(visitor); err != nil {
			return err
		}
	}
	if err := t.Table.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTableIndex(a *TableIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Name.Accept(visitor); err != nil {
		return err
	}
	if err := a.ColumnExpr.Accept(visitor); err != nil {
		return err
	}
	if err := a.ColumnType.Accept(visitor); err != nil {
		return err
	}
	if err := a.Granularity.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTableProjection(t *TableProjection) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Identifier.Accept(visitor); err != nil {
		return err
	}
	if err := t.Select.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitTableSchemaClause(t *TableSchemaClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, column := range t.Columns {
		if err := column.Accept(visitor); err != nil {
			return err
		}
	}
	if t.AliasTable != nil {
		if err := t.AliasTable.Accept(visitor); err != nil {
			return err
		}
	}
	if t.TableFunction != nil {
		if err := t.TableFunction.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTernaryExpr(t *TernaryOperation) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.TrueExpr.Accept(visitor); err != nil {
		return err
	}
	if err := t.FalseExpr.Accept(visitor); err != nil {
		return err
	}
	if err := t.Condition.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitTopClause(t *TopClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Number.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitTruncateTable(t *TruncateTable) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Name.Accept(visitor); err != nil {
		return err
	}
	if t.OnCluster != nil {
		if err := t.OnCluster.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitTypeWithParams(s *TypeWithParams) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor); err != nil {
		return err
	}
	for _, param := range s.Params {
		if err := param.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitUUID(u *UUID) error {
	visitor.enter(u)
	defer visitor.leave(u)
	return nil
}

func (visitor MyDefaultVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitUseStmt(u *UseStmt) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Database.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitUsingClause(u *UsingClause) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Using.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitWhenClause(w *WhenClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.When.Accept(visitor); err != nil {
		return err
	}
	if err := w.Then.Accept(visitor); err != nil {
		return err
	}
	if w.Else != nil {
		if err := w.Else.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWhereClause(w *WhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWindowClause(w *WindowClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if w.WindowExpr != nil {
		if err := w.WindowExpr.Accept(visitor); err != nil {
			return err
		}
	}
	if w.Name != nil {
		if err := w.Name.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitWindowExpr(w *WindowExpr) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if w.PartitionBy != nil {
		if err := w.PartitionBy.Accept(visitor); err != nil {
			return err
		}
	}
	if w.OrderBy != nil {
		if err := w.OrderBy.Accept(visitor); err != nil {
			return err
		}
	}
	if w.Frame != nil {
		if err := w.Frame.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWindowFrameClause(f *WindowFrameClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Extend.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}
func (visitor MyDefaultVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Number.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}
func (visitor MyDefaultVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Function.Accept(visitor); err != nil {
		return err
	}
	if err := w.OverExpr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWithClause(w *WithClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	for _, cte := range w.CTEs {
		if err := cte.Accept(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitWithTimeoutClause(w *WithTimeoutClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Number.Accept(visitor); err != nil {
		return err
	}
	return nil
}
func (visitor MyDefaultVisitor) enter(expr Expr) {
}
func (visitor MyDefaultVisitor) leave(expr Expr) {

}
