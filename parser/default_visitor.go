package parser

type DefaultVisitor struct {
	self ASTVisitor
}

func (visitor DefaultVisitor) VisitAliasExpr(a *AliasExpr) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if err := a.Alias.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterRole(a *AlterRole) error {
	visitor.enter(a)
	defer visitor.leave(a)
	for _, roleRenamePair := range a.RoleRenamePairs {
		if err := roleRenamePair.Accept(visitor.self); err != nil {
			return err
		}
	}
	for _, setting := range a.Settings {
		if err := setting.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTable(a *AlterTable) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if a.OnCluster != nil {
		if err := a.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}

	for _, expr := range a.AlterExprs {
		if err := expr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableAddColumn(a *AlterTableAddColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Column.Accept(visitor.self); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableAddIndex(a *AlterTableAddIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Index.Accept(visitor.self); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableAddProjection(a *AlterTableAddProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.TableProjection.Accept(visitor.self); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableAttachPartition(a *AlterTableAttachPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.self); err != nil {
		return err
	}
	if a.From != nil {
		if err := a.From.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableClearColumn(a *AlterTableClearColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor.self); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableClearIndex(a *AlterTableClearIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor.self); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableClearProjection(a *AlterTableClearProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor.self); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableDetachPartition(a *AlterTableDetachPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.self); err != nil {
		return err
	}
	if a.Settings != nil {
		if err := a.Settings.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableDropColumn(a *AlterTableDropColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableMaterializeIndex(a *AlterTableMaterializeIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor.self); err != nil {
		return err
	}
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableMaterializeProjection(a *AlterTableMaterializeProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor.self); err != nil {
		return err
	}
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableModifyColumn(a *AlterTableModifyColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Column.Accept(visitor.self); err != nil {
		return err
	}
	if a.RemovePropertyType != nil {
		if err := a.RemovePropertyType.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.TTL.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableRenameColumn(a *AlterTableRenameColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.OldColumnName.Accept(visitor.self); err != nil {
		return err
	}
	if err := a.NewColumnName.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.self); err != nil {
		return err
	}
	if err := a.Table.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitArrayJoinExpr(a *ArrayJoinClause) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitArrayParamList(a *ArrayParamList) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Items.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

// nolint: funlen

// START, STOP
// REPLICATED, DISTRIBUTED

func (visitor DefaultVisitor) VisitValuesExpr(v *AssignmentValues) error {
	visitor.enter(v)
	defer visitor.leave(v)
	for _, value := range v.Values {
		if err := value.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitBetweenClause(f *BetweenClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Between.Accept(visitor.self); err != nil {
		return err
	}
	if err := f.And.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitBinaryExpr(p *BinaryOperation) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.LeftExpr.Accept(visitor.self); err != nil {
		return err
	}
	if err := p.RightExpr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitCTEExpr(c *CTEStmt) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if err := c.Alias.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitCaseExpr(c *CaseExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.Expr != nil {
		if err := c.Expr.Accept(visitor.self); err != nil {
			return err
		}
	}
	for _, when := range c.Whens {
		if err := when.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Else != nil {
		if err := c.Else.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitCastExpr(c *CastExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if err := c.AsType.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitCheckExpr(c *CheckStmt) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Table.Accept(visitor.self); err != nil {
		return err
	}
	if c.Partition != nil {
		if err := c.Partition.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitOnClusterExpr(o *ClusterClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitColumnArgList(c *ColumnArgList) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitColumnDef(c *ColumnDef) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	if c.Type != nil {
		if err := c.Type.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.NotNull != nil {
		if err := c.NotNull.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Nullable != nil {
		if err := c.Nullable.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.DefaultExpr != nil {
		if err := c.DefaultExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.MaterializedExpr != nil {
		if err := c.MaterializedExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.AliasExpr != nil {
		if err := c.AliasExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Codec != nil {
		if err := c.Codec.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.TTL != nil {
		if err := c.TTL.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		if err := c.Comment.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitColumnExpr(c *ColumnExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if c.Alias != nil {
		if err := c.Alias.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitColumnExprList(c *ColumnExprList) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitColumnIdentifier(c *ColumnIdentifier) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.Database != nil {
		if err := c.Database.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Table != nil {
		if err := c.Table.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := c.Column.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitColumnNamesExpr(c *ColumnNamesExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for i := range c.ColumnNames {
		if err := c.ColumnNames[i].Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitComplexType(c *ComplexType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	for _, param := range c.Params {
		if err := param.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitCompressionCodec(c *CompressionCodec) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Type.Accept(visitor.self); err != nil {
		return err
	}
	if c.TypeLevel != nil {
		if err := c.TypeLevel.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	if c.Level != nil {
		if err := c.Level.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitConstraintExpr(c *ConstraintClause) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Constraint.Accept(visitor.self); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitCreateDatabase(c *CreateDatabase) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitCreateFunction(c *CreateFunction) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.FunctionName.Accept(visitor.self); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := c.Params.Accept(visitor.self); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitCreateLiveView(c *CreateLiveView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		if err := c.Destination.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.WithTimeout != nil {
		if err := c.WithTimeout.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitCreateMaterializedView(c *CreateMaterializedView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		if err := c.Destination.Accept(visitor.self); err != nil {
			return err
		}
		if c.Destination.TableSchema != nil {
			if err := c.Destination.TableSchema.Accept(visitor.self); err != nil {
				return err
			}
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitCreateRole(c *CreateRole) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, roleName := range c.RoleNames {
		if err := roleName.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.AccessStorageType != nil {
		if err := c.AccessStorageType.Accept(visitor.self); err != nil {
			return err
		}
	}
	for _, setting := range c.Settings {
		if err := setting.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitCreateTable(c *CreateTable) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitCreateView(c *CreateView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor.self); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitDeduplicateExpr(d *DeduplicateClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if d.By != nil {
		if err := d.By.Accept(visitor.self); err != nil {
			return err
		}
	}
	if d.Except != nil {
		if err := d.Except.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitDeleteFromExpr(d *DeleteClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Table.Accept(visitor.self); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if d.WhereExpr != nil {
		if err := d.WhereExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitDestinationExpr(d *DestinationClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.TableIdentifier.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitDropDatabase(d *DropDatabase) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Name.Accept(visitor.self); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitDropStmt(d *DropStmt) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Name.Accept(visitor.self); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil

}

func (visitor DefaultVisitor) VisitDropUserOrRole(d *DropUserOrRole) error {
	visitor.enter(d)
	defer visitor.leave(d)
	for _, name := range d.Names {
		if err := name.Accept(visitor.self); err != nil {
			return err
		}
	}
	if d.From != nil {
		if err := d.From.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitEngineExpr(e *EngineExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if e.Params != nil {
		if err := e.Params.Accept(visitor.self); err != nil {
			return err
		}
	}
	if e.PrimaryKey != nil {
		if err := e.PrimaryKey.Accept(visitor.self); err != nil {
			return err
		}
	}
	if e.PartitionBy != nil {
		if err := e.PartitionBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if e.SampleBy != nil {
		if err := e.SampleBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if e.TTL != nil {
		if err := e.TTL.Accept(visitor.self); err != nil {
			return err
		}
	}
	if e.Settings != nil {
		if err := e.Settings.Accept(visitor.self); err != nil {
			return err
		}
	}
	if e.OrderBy != nil {
		if err := e.OrderBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitEnumType(e *EnumType) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Name.Accept(visitor.self); err != nil {
		return err
	}
	for i := range e.Values {
		if err := e.Values[i].Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitEnumValue(e *EnumValue) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Name.Accept(visitor.self); err != nil {
		return err
	}
	if err := e.Value.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitExplainExpr(e *ExplainStmt) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Statement.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitExtractExpr(e *ExtractExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.FromExpr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitFormatExpr(f *FormatClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Format.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitFromExpr(f *FromClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Name.Accept(visitor.self); err != nil {
		return err
	}
	if err := f.Params.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitGlobalInExpr(g *GlobalInOperation) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitGrantPrivilegeExpr(g *GrantPrivilegeStmt) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if g.OnCluster != nil {
		if err := g.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	for _, privilege := range g.Privileges {
		if err := privilege.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := g.On.Accept(visitor.self); err != nil {
		return err
	}
	for _, role := range g.To {
		if err := role.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitGroupByExpr(g *GroupByClause) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitHavingExpr(h *HavingClause) error {
	visitor.enter(h)
	defer visitor.leave(h)
	if err := h.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitIdent(i *Ident) error {
	visitor.enter(i)
	defer visitor.leave(i)
	return nil
}
func (visitor DefaultVisitor) VisitIndexOperation(i *IndexOperation) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if err := i.Object.Accept(visitor.self); err != nil {
		return err
	}
	if err := i.Index.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitInsertExpr(i *InsertStmt) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if i.Format != nil {
		if err := i.Format.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := i.Table.Accept(visitor.self); err != nil {
		return err
	}
	if i.ColumnNames != nil {
		if err := i.ColumnNames.Accept(visitor.self); err != nil {
			return err
		}
	}
	for _, value := range i.Values {
		if err := value.Accept(visitor.self); err != nil {
			return err
		}
	}
	if i.SelectExpr != nil {
		if err := i.SelectExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitIntervalExpr(i *IntervalExpr) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if err := i.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if err := i.Unit.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitJSONType(j *JSONType) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Name.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitJoinConstraintExpr(j *JoinConstraintClause) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if j.On != nil {
		if err := j.On.Accept(visitor.self); err != nil {
			return err
		}
	}
	if j.Using != nil {
		if err := j.Using.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitJoinExpr(j *JoinExpr) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Left.Accept(visitor.self); err != nil {
		return err
	}
	if j.Right != nil {
		if err := j.Right.Accept(visitor.self); err != nil {
			return err
		}
	}
	if j.Constraints != nil {
		if err := j.Constraints.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitJoinTableExpr(j *JoinTableExpr) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Table.Accept(visitor.self); err != nil {
		return err
	}
	if j.SampleRatio != nil {
		return j.SampleRatio.Accept(visitor.self)
	}
	return nil
}

func (visitor DefaultVisitor) VisitLimitByExpr(l *LimitByClause) error {
	visitor.enter(l)
	defer visitor.leave(l)
	if l.Limit != nil {
		if err := l.Limit.Accept(visitor.self); err != nil {
			return err
		}
	}
	if l.ByExpr != nil {
		if err := l.ByExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitLimitExpr(l *LimitClause) error {
	visitor.enter(l)
	defer visitor.leave(l)
	if err := l.Limit.Accept(visitor.self); err != nil {
		return err
	}
	if l.Offset != nil {
		if err := l.Offset.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitMapLiteral(m *MapLiteral) error {
	visitor.enter(m)
	defer visitor.leave(m)
	for _, kv := range m.KeyValues {
		if err := kv.Key.Accept(visitor.self); err != nil {
			return err
		}
		if err := kv.Value.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitNegateExpr(n *NegateExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Ident.Accept(visitor.self); err != nil {
		return err
	}
	if n.DotIdent != nil {
		if err := n.DotIdent.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitNestedType(n *NestedType) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Name.Accept(visitor.self); err != nil {
		return err
	}
	for _, column := range n.Columns {
		if err := column.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitNotExpr(n *NotExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.NullLiteral.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitNullLiteral(n *NullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}

func (visitor DefaultVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}
func (visitor DefaultVisitor) VisitObjectParams(o *ObjectParams) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Object.Accept(visitor.self); err != nil {
		return err
	}
	if err := o.Params.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitOnExpr(o *OnClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.On.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitOperationExpr(o *OperationExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	return nil
}

func (visitor DefaultVisitor) VisitOptimizeExpr(o *OptimizeStmt) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Table.Accept(visitor.self); err != nil {
		return err
	}
	if o.OnCluster != nil {
		if err := o.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	if o.Partition != nil {
		if err := o.Partition.Accept(visitor.self); err != nil {
			return err
		}
	}
	if o.Deduplicate != nil {
		if err := o.Deduplicate.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitOrderByListExpr(o *OrderByClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	for _, item := range o.Items {
		if err := item.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitOrderByExpr(o *OrderExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if o.Alias != nil {
		if err := o.Alias.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitParamExprList(f *ParamExprList) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Items.Accept(visitor.self); err != nil {
		return err
	}
	if f.ColumnArgList != nil {
		if err := f.ColumnArgList.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitPartitionByExpr(p *PartitionByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitPartitionExpr(p *PartitionClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.Expr != nil {
		if err := p.Expr.Accept(visitor.self); err != nil {
			return err
		}
	}
	if p.ID != nil {
		if err := p.ID.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitPlaceHolderExpr(p *PlaceHolder) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

func (visitor DefaultVisitor) VisitPrewhereExpr(w *PrewhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitPrimaryKeyExpr(p *PrimaryKeyClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitPrivilegeExpr(p *PrivilegeClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.Params != nil {
		if err := p.Params.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitProjectionOrderBy(p *ProjectionOrderByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

func (visitor DefaultVisitor) VisitProjectionSelect(p *ProjectionSelectStmt) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.With != nil {
		if err := p.With.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := p.SelectColumns.Accept(visitor.self); err != nil {
		return err
	}
	if p.GroupBy != nil {
		if err := p.GroupBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if p.OrderBy != nil {
		if err := p.OrderBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitPropertyType(c *PropertyType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitQueryParam(q *QueryParam) error {
	visitor.enter(q)
	defer visitor.leave(q)
	if err := q.Name.Accept(visitor.self); err != nil {
		return err
	}
	if err := q.Type.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitRatioExpr(r *RatioExpr) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.Numerator.Accept(visitor.self); err != nil {
		return err
	}
	if r.Denominator != nil {
		if err := r.Denominator.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitRemovePropertyType(a *RemovePropertyType) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.PropertyType.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitRenameStmt(r *RenameStmt) error {
	visitor.enter(r)
	defer visitor.leave(r)
	for _, pair := range r.TargetPairList {
		if err := pair.Old.Accept(visitor.self); err != nil {
			return err
		}
		if err := pair.New.Accept(visitor.self); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		if err := r.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitRoleName(r *RoleName) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.Name.Accept(visitor.self); err != nil {
		return err
	}
	if r.Scope != nil {
		if err := r.Scope.Accept(visitor.self); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		if err := r.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitRoleRenamePair(r *RoleRenamePair) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.RoleName.Accept(visitor.self); err != nil {
		return err
	}
	if r.NewName != nil {
		if err := r.NewName.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitRoleSetting(r *RoleSetting) error {
	visitor.enter(r)
	defer visitor.leave(r)
	for _, settingPair := range r.SettingPairs {
		if err := settingPair.Accept(visitor.self); err != nil {
			return err
		}
	}
	if r.Modifier != nil {
		if err := r.Modifier.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSampleByExpr(s *SampleByClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitSampleRatioExpr(s *SampleClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Ratio.Accept(visitor.self); err != nil {
		return err
	}
	if s.Offset != nil {
		if err := s.Offset.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitScalarType(s *ScalarType) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitSelectItem(s *SelectItem) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor.self); err != nil {
		return err
	}
	for _, modifier := range s.Modifiers {
		if err := modifier.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Alias != nil {
		if err := s.Alias.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSelectQuery(s *SelectQuery) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.With != nil {
		if err := s.With.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Top != nil {
		if err := s.Top.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.SelectItems != nil {
		for _, item := range s.SelectItems {
			if err := item.Accept(visitor.self); err != nil {
				return err
			}
		}
	}
	if s.From != nil {
		if err := s.From.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.ArrayJoin != nil {
		if err := s.ArrayJoin.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Window != nil {
		if err := s.Window.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Prewhere != nil {
		if err := s.Prewhere.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Where != nil {
		if err := s.Where.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Having != nil {
		if err := s.Having.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.OrderBy != nil {
		if err := s.OrderBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.LimitBy != nil {
		if err := s.LimitBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Limit != nil {
		if err := s.Limit.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Format != nil {
		if err := s.Format.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.UnionAll != nil {
		if err := s.UnionAll.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.UnionDistinct != nil {
		if err := s.UnionDistinct.Accept(visitor.self); err != nil {
			return err
		}
	}
	if s.Except != nil {
		if err := s.Except.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSetExpr(s *SetStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Settings.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitSettingsExpr(s *SettingExprList) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.self); err != nil {
		return err
	}
	if err := s.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitSettingPair(s *SettingPair) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.self); err != nil {
		return err
	}
	if s.Value != nil {
		if err := s.Value.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSettingsExprList(s *SettingsClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	for _, item := range s.Items {
		if err := item.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitStringLiteral(s *StringLiteral) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}

func (visitor DefaultVisitor) VisitSubQueryExpr(s *SubQuery) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Select != nil {
		if err := s.Select.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSystemCtrlExpr(s *SystemCtrlExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Cluster != nil {
		if err := s.Cluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}
func (visitor DefaultVisitor) VisitSystemFlushExpr(s *SystemFlushExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Distributed != nil {
		if err := s.Distributed.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitSystemReloadExpr(s *SystemReloadExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Dictionary != nil {
		if err := s.Dictionary.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitSystemExpr(s *SystemStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitSystemSyncExpr(s *SystemSyncExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Cluster.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitTTLExprList(t *TTLClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, item := range t.Items {
		if err := item.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitTTLExpr(t *TTLExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitTableArgListExpr(t *TableArgListExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, arg := range t.Args {
		if err := arg.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitTableExpr(t *TableExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Expr.Accept(visitor.self); err != nil {
		return err
	}
	if t.Alias != nil {
		if err := t.Alias.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitTableFunctionExpr(t *TableFunctionExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Name.Accept(visitor.self); err != nil {
		return err
	}
	if err := t.Args.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitTableIdentifier(t *TableIdentifier) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if t.Database != nil {
		if err := t.Database.Accept(visitor.self); err != nil {
			return err
		}
	}
	if err := t.Table.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitTableIndex(a *TableIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Name.Accept(visitor.self); err != nil {
		return err
	}
	if err := a.ColumnExpr.Accept(visitor.self); err != nil {
		return err
	}
	if err := a.ColumnType.Accept(visitor.self); err != nil {
		return err
	}
	if err := a.Granularity.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitTableProjection(t *TableProjection) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Identifier.Accept(visitor.self); err != nil {
		return err
	}
	if err := t.Select.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitTableSchemaExpr(t *TableSchemaClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, column := range t.Columns {
		if err := column.Accept(visitor.self); err != nil {
			return err
		}
	}
	if t.AliasTable != nil {
		if err := t.AliasTable.Accept(visitor.self); err != nil {
			return err
		}
	}
	if t.TableFunction != nil {
		if err := t.TableFunction.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitTernaryExpr(t *TernaryOperation) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.TrueExpr.Accept(visitor.self); err != nil {
		return err
	}
	if err := t.FalseExpr.Accept(visitor.self); err != nil {
		return err
	}
	if err := t.Condition.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitTopExpr(t *TopClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Number.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitTruncateTable(t *TruncateTable) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Name.Accept(visitor.self); err != nil {
		return err
	}
	if t.OnCluster != nil {
		if err := t.OnCluster.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitTypeWithParams(s *TypeWithParams) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.self); err != nil {
		return err
	}
	for _, param := range s.Params {
		if err := param.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitUUID(u *UUID) error {
	visitor.enter(u)
	defer visitor.leave(u)
	return nil
}

func (visitor DefaultVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitUseExpr(u *UseStmt) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Database.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitUsingExpr(u *UsingClause) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Using.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitWhenExpr(w *WhenClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.When.Accept(visitor.self); err != nil {
		return err
	}
	if err := w.Then.Accept(visitor.self); err != nil {
		return err
	}
	if w.Else != nil {
		if err := w.Else.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitWhereExpr(w *WhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitWindowExpr(w *WindowClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if w.WindowExpr != nil {
		if err := w.WindowExpr.Accept(visitor.self); err != nil {
			return err
		}
	}
	if w.Name != nil {
		if err := w.Name.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultVisitor) VisitWindowConditionExpr(w *WindowExpr) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if w.PartitionBy != nil {
		if err := w.PartitionBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if w.OrderBy != nil {
		if err := w.OrderBy.Accept(visitor.self); err != nil {
			return err
		}
	}
	if w.Frame != nil {
		if err := w.Frame.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitWindowFrameExpr(f *WindowFrameClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Extend.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}
func (visitor DefaultVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Number.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}
func (visitor DefaultVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Function.Accept(visitor.self); err != nil {
		return err
	}
	if err := w.OverExpr.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) VisitWithExpr(w *WithClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	for _, cte := range w.CTEs {
		if err := cte.Accept(visitor.self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultVisitor) VisitWithTimeoutExpr(w *WithTimeoutClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Number.Accept(visitor.self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultVisitor) enter(expr Expr) {}

func (visitor DefaultVisitor) leave(expr Expr) {}
