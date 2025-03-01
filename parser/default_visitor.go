package parser

type DefaultASTVisitor struct {
	Self ASTVisitor
}

func (visitor DefaultASTVisitor) VisitAliasExpr(a *AliasExpr) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.Alias.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterRole(a *AlterRole) error {
	visitor.enter(a)
	defer visitor.leave(a)
	for _, roleRenamePair := range a.RoleRenamePairs {
		if err := roleRenamePair.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, setting := range a.Settings {
		if err := setting.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTable(a *AlterTable) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if a.OnCluster != nil {
		if err := a.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}

	for _, expr := range a.AlterExprs {
		if err := expr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableAddColumn(a *AlterTableAddColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Column.Accept(visitor.Self); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableAddIndex(a *AlterTableAddIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Index.Accept(visitor.Self); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableAddProjection(a *AlterTableAddProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.TableProjection.Accept(visitor.Self); err != nil {
		return err
	}
	if a.After != nil {
		if err := a.After.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableAttachPartition(a *AlterTableAttachPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.Self); err != nil {
		return err
	}
	if a.From != nil {
		if err := a.From.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableClearColumn(a *AlterTableClearColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableClearIndex(a *AlterTableClearIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor.Self); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableClearProjection(a *AlterTableClearProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor.Self); err != nil {
		return err
	}
	if a.PartitionExpr != nil {
		if err := a.PartitionExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableDetachPartition(a *AlterTableDetachPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.Self); err != nil {
		return err
	}
	if a.Settings != nil {
		if err := a.Settings.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableDropColumn(a *AlterTableDropColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableMaterializeIndex(a *AlterTableMaterializeIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor.Self); err != nil {
		return err
	}
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableMaterializeProjection(a *AlterTableMaterializeProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor.Self); err != nil {
		return err
	}
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableModifyColumn(a *AlterTableModifyColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Column.Accept(visitor.Self); err != nil {
		return err
	}
	if a.RemovePropertyType != nil {
		if err := a.RemovePropertyType.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.TTL.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableRenameColumn(a *AlterTableRenameColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.OldColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.NewColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.Table.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitArrayJoinExpr(a *ArrayJoinClause) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitArrayParamList(a *ArrayParamList) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Items.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

// nolint: funlen

// START, STOP
// REPLICATED, DISTRIBUTED

func (visitor DefaultASTVisitor) VisitValuesExpr(v *AssignmentValues) error {
	visitor.enter(v)
	defer visitor.leave(v)
	for _, value := range v.Values {
		if err := value.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitBetweenClause(f *BetweenClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Between.Accept(visitor.Self); err != nil {
		return err
	}
	if err := f.And.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitBinaryExpr(p *BinaryOperation) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.LeftExpr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := p.RightExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCTEExpr(c *CTEStmt) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.Alias.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitCaseExpr(c *CaseExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.Expr != nil {
		if err := c.Expr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, when := range c.Whens {
		if err := when.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Else != nil {
		if err := c.Else.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCastExpr(c *CastExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.AsType.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCheckExpr(c *CheckStmt) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Table.Accept(visitor.Self); err != nil {
		return err
	}
	if c.Partition != nil {
		if err := c.Partition.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitOnClusterExpr(o *ClusterClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitColumnArgList(c *ColumnArgList) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnDef(c *ColumnDef) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if c.Type != nil {
		if err := c.Type.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.NotNull != nil {
		if err := c.NotNull.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Nullable != nil {
		if err := c.Nullable.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.DefaultExpr != nil {
		if err := c.DefaultExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.MaterializedExpr != nil {
		if err := c.MaterializedExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.AliasExpr != nil {
		if err := c.AliasExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Codec != nil {
		if err := c.Codec.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.TTL != nil {
		if err := c.TTL.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Comment != nil {
		if err := c.Comment.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnExpr(c *ColumnExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if c.Alias != nil {
		if err := c.Alias.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitColumnExprList(c *ColumnExprList) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnIdentifier(c *ColumnIdentifier) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.Database != nil {
		if err := c.Database.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Table != nil {
		if err := c.Table.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := c.Column.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitColumnNamesExpr(c *ColumnNamesExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for i := range c.ColumnNames {
		if err := c.ColumnNames[i].Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitComplexType(c *ComplexType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	for _, param := range c.Params {
		if err := param.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCompressionCodec(c *CompressionCodec) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Type.Accept(visitor.Self); err != nil {
		return err
	}
	if c.TypeLevel != nil {
		if err := c.TypeLevel.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if c.Level != nil {
		if err := c.Level.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitConstraintExpr(c *ConstraintClause) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Constraint.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitCreateDatabase(c *CreateDatabase) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateFunction(c *CreateFunction) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.FunctionName.Accept(visitor.Self); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := c.Params.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateLiveView(c *CreateLiveView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		if err := c.Destination.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.WithTimeout != nil {
		if err := c.WithTimeout.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitCreateMaterializedView(c *CreateMaterializedView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Destination != nil {
		if err := c.Destination.Accept(visitor.Self); err != nil {
			return err
		}
		if c.Destination.TableSchema != nil {
			if err := c.Destination.TableSchema.Accept(visitor.Self); err != nil {
				return err
			}
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateRole(c *CreateRole) error {
	visitor.enter(c)
	defer visitor.leave(c)
	for _, roleName := range c.RoleNames {
		if err := roleName.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.AccessStorageType != nil {
		if err := c.AccessStorageType.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, setting := range c.Settings {
		if err := setting.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitCreateTable(c *CreateTable) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Engine != nil {
		if err := c.Engine.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateView(c *CreateView) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if c.UUID != nil {
		if err := c.UUID.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.TableSchema != nil {
		if err := c.TableSchema.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.SubQuery != nil {
		if err := c.SubQuery.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDeduplicateExpr(d *DeduplicateClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if d.By != nil {
		if err := d.By.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Except != nil {
		if err := d.Except.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDeleteFromExpr(d *DeleteClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Table.Accept(visitor.Self); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.WhereExpr != nil {
		if err := d.WhereExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitDestinationExpr(d *DestinationClause) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.TableIdentifier.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDropDatabase(d *DropDatabase) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDropStmt(d *DropStmt) error {
	visitor.enter(d)
	defer visitor.leave(d)
	if err := d.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if d.OnCluster != nil {
		if err := d.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil

}

func (visitor DefaultASTVisitor) VisitDropUserOrRole(d *DropUserOrRole) error {
	visitor.enter(d)
	defer visitor.leave(d)
	for _, name := range d.Names {
		if err := name.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.From != nil {
		if err := d.From.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitEngineExpr(e *EngineExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if e.Params != nil {
		if err := e.Params.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if e.PrimaryKey != nil {
		if err := e.PrimaryKey.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if e.PartitionBy != nil {
		if err := e.PartitionBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if e.SampleBy != nil {
		if err := e.SampleBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if e.TTL != nil {
		if err := e.TTL.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if e.Settings != nil {
		if err := e.Settings.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if e.OrderBy != nil {
		if err := e.OrderBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitEnumType(e *EnumType) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Name.Accept(visitor.Self); err != nil {
		return err
	}
	for i := range e.Values {
		if err := e.Values[i].Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitEnumValue(e *EnumValue) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := e.Value.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitExplainExpr(e *ExplainStmt) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.Statement.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitExtractExpr(e *ExtractExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.FromExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitFormatExpr(f *FormatClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Format.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitFromExpr(f *FromClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := f.Params.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitGlobalInExpr(g *GlobalInOperation) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitGrantPrivilegeExpr(g *GrantPrivilegeStmt) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if g.OnCluster != nil {
		if err := g.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, privilege := range g.Privileges {
		if err := privilege.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := g.On.Accept(visitor.Self); err != nil {
		return err
	}
	for _, role := range g.To {
		if err := role.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitGroupByExpr(g *GroupByClause) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitHavingExpr(h *HavingClause) error {
	visitor.enter(h)
	defer visitor.leave(h)
	if err := h.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitIdent(i *Ident) error {
	visitor.enter(i)
	defer visitor.leave(i)
	return nil
}
func (visitor DefaultASTVisitor) VisitIndexOperation(i *IndexOperation) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if err := i.Object.Accept(visitor.Self); err != nil {
		return err
	}
	if err := i.Index.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitInsertExpr(i *InsertStmt) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if i.Format != nil {
		if err := i.Format.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := i.Table.Accept(visitor.Self); err != nil {
		return err
	}
	if i.ColumnNames != nil {
		if err := i.ColumnNames.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, value := range i.Values {
		if err := value.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if i.SelectExpr != nil {
		if err := i.SelectExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitIntervalExpr(i *IntervalExpr) error {
	visitor.enter(i)
	defer visitor.leave(i)
	if err := i.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := i.Unit.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitJSONType(j *JSONType) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitJoinConstraintExpr(j *JoinConstraintClause) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if j.On != nil {
		if err := j.On.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if j.Using != nil {
		if err := j.Using.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitJoinExpr(j *JoinExpr) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Left.Accept(visitor.Self); err != nil {
		return err
	}
	if j.Right != nil {
		if err := j.Right.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if j.Constraints != nil {
		if err := j.Constraints.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitJoinTableExpr(j *JoinTableExpr) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Table.Accept(visitor.Self); err != nil {
		return err
	}
	if j.SampleRatio != nil {
		return j.SampleRatio.Accept(visitor.Self)
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitLimitByExpr(l *LimitByClause) error {
	visitor.enter(l)
	defer visitor.leave(l)
	if l.Limit != nil {
		if err := l.Limit.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if l.ByExpr != nil {
		if err := l.ByExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitLimitExpr(l *LimitClause) error {
	visitor.enter(l)
	defer visitor.leave(l)
	if err := l.Limit.Accept(visitor.Self); err != nil {
		return err
	}
	if l.Offset != nil {
		if err := l.Offset.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitMapLiteral(m *MapLiteral) error {
	visitor.enter(m)
	defer visitor.leave(m)
	for _, kv := range m.KeyValues {
		if err := kv.Key.Accept(visitor.Self); err != nil {
			return err
		}
		if err := kv.Value.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitNegateExpr(n *NegateExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Ident.Accept(visitor.Self); err != nil {
		return err
	}
	if n.DotIdent != nil {
		if err := n.DotIdent.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitNestedType(n *NestedType) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Name.Accept(visitor.Self); err != nil {
		return err
	}
	for _, column := range n.Columns {
		if err := column.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitNotExpr(n *NotExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.NullLiteral.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitNullLiteral(n *NullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}

func (visitor DefaultASTVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}
func (visitor DefaultASTVisitor) VisitObjectParams(o *ObjectParams) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Object.Accept(visitor.Self); err != nil {
		return err
	}
	if err := o.Params.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitOnExpr(o *OnClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.On.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitOperationExpr(o *OperationExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	return nil
}

func (visitor DefaultASTVisitor) VisitOptimizeExpr(o *OptimizeStmt) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Table.Accept(visitor.Self); err != nil {
		return err
	}
	if o.OnCluster != nil {
		if err := o.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if o.Partition != nil {
		if err := o.Partition.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if o.Deduplicate != nil {
		if err := o.Deduplicate.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitOrderByListExpr(o *OrderByClause) error {
	visitor.enter(o)
	defer visitor.leave(o)
	for _, item := range o.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitOrderByExpr(o *OrderExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	if err := o.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if o.Alias != nil {
		if err := o.Alias.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitParamExprList(f *ParamExprList) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Items.Accept(visitor.Self); err != nil {
		return err
	}
	if f.ColumnArgList != nil {
		if err := f.ColumnArgList.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitPartitionByExpr(p *PartitionByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitPartitionExpr(p *PartitionClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.Expr != nil {
		if err := p.Expr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if p.ID != nil {
		if err := p.ID.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitPlaceHolderExpr(p *PlaceHolder) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

func (visitor DefaultASTVisitor) VisitPrewhereExpr(w *PrewhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitPrimaryKeyExpr(p *PrimaryKeyClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if err := p.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitPrivilegeExpr(p *PrivilegeClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.Params != nil {
		if err := p.Params.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitProjectionOrderBy(p *ProjectionOrderByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

func (visitor DefaultASTVisitor) VisitProjectionSelect(p *ProjectionSelectStmt) error {
	visitor.enter(p)
	defer visitor.leave(p)
	if p.With != nil {
		if err := p.With.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := p.SelectColumns.Accept(visitor.Self); err != nil {
		return err
	}
	if p.GroupBy != nil {
		if err := p.GroupBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if p.OrderBy != nil {
		if err := p.OrderBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitPropertyType(c *PropertyType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitQueryParam(q *QueryParam) error {
	visitor.enter(q)
	defer visitor.leave(q)
	if err := q.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := q.Type.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitRatioExpr(r *RatioExpr) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.Numerator.Accept(visitor.Self); err != nil {
		return err
	}
	if r.Denominator != nil {
		if err := r.Denominator.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitRemovePropertyType(a *RemovePropertyType) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.PropertyType.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitRenameStmt(r *RenameStmt) error {
	visitor.enter(r)
	defer visitor.leave(r)
	for _, pair := range r.TargetPairList {
		if err := pair.Old.Accept(visitor.Self); err != nil {
			return err
		}
		if err := pair.New.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		if err := r.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitRoleName(r *RoleName) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if r.Scope != nil {
		if err := r.Scope.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if r.OnCluster != nil {
		if err := r.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitRoleRenamePair(r *RoleRenamePair) error {
	visitor.enter(r)
	defer visitor.leave(r)
	if err := r.RoleName.Accept(visitor.Self); err != nil {
		return err
	}
	if r.NewName != nil {
		if err := r.NewName.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitRoleSetting(r *RoleSetting) error {
	visitor.enter(r)
	defer visitor.leave(r)
	for _, settingPair := range r.SettingPairs {
		if err := settingPair.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if r.Modifier != nil {
		if err := r.Modifier.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSampleByExpr(s *SampleByClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSampleRatioExpr(s *SampleClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Ratio.Accept(visitor.Self); err != nil {
		return err
	}
	if s.Offset != nil {
		if err := s.Offset.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitScalarType(s *ScalarType) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSelectItem(s *SelectItem) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	for _, modifier := range s.Modifiers {
		if err := modifier.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Alias != nil {
		if err := s.Alias.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSelectQuery(s *SelectQuery) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.With != nil {
		if err := s.With.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Top != nil {
		if err := s.Top.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.SelectItems != nil {
		for _, item := range s.SelectItems {
			if err := item.Accept(visitor.Self); err != nil {
				return err
			}
		}
	}
	if s.From != nil {
		if err := s.From.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.ArrayJoin != nil {
		if err := s.ArrayJoin.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Window != nil {
		if err := s.Window.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Prewhere != nil {
		if err := s.Prewhere.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Where != nil {
		if err := s.Where.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Having != nil {
		if err := s.Having.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.OrderBy != nil {
		if err := s.OrderBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.LimitBy != nil {
		if err := s.LimitBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Limit != nil {
		if err := s.Limit.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Format != nil {
		if err := s.Format.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.UnionAll != nil {
		if err := s.UnionAll.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.UnionDistinct != nil {
		if err := s.UnionDistinct.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Except != nil {
		if err := s.Except.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSetExpr(s *SetStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Settings.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSettingsExpr(s *SettingExprList) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSettingPair(s *SettingPair) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if s.Value != nil {
		if err := s.Value.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSettingsExprList(s *SettingsClause) error {
	visitor.enter(s)
	defer visitor.leave(s)
	for _, item := range s.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitStringLiteral(s *StringLiteral) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}

func (visitor DefaultASTVisitor) VisitSubQueryExpr(s *SubQuery) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Select != nil {
		if err := s.Select.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemCtrlExpr(s *SystemCtrlExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Cluster != nil {
		if err := s.Cluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}
func (visitor DefaultASTVisitor) VisitSystemFlushExpr(s *SystemFlushExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Distributed != nil {
		if err := s.Distributed.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemReloadExpr(s *SystemReloadExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if s.Dictionary != nil {
		if err := s.Dictionary.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSystemExpr(s *SystemStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemSyncExpr(s *SystemSyncExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Cluster.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTTLExprList(t *TTLClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, item := range t.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTTLExpr(t *TTLExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableArgListExpr(t *TableArgListExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, arg := range t.Args {
		if err := arg.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTableExpr(t *TableExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if t.Alias != nil {
		if err := t.Alias.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableFunctionExpr(t *TableFunctionExpr) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.Args.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableIdentifier(t *TableIdentifier) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if t.Database != nil {
		if err := t.Database.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := t.Table.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableIndex(a *TableIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.ColumnExpr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.ColumnType.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.Granularity.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableProjection(t *TableProjection) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Identifier.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.Select.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTableSchemaExpr(t *TableSchemaClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	for _, column := range t.Columns {
		if err := column.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if t.AliasTable != nil {
		if err := t.AliasTable.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if t.TableFunction != nil {
		if err := t.TableFunction.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTernaryExpr(t *TernaryOperation) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.TrueExpr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.FalseExpr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.Condition.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTopExpr(t *TopClause) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Number.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTruncateTable(t *TruncateTable) error {
	visitor.enter(t)
	defer visitor.leave(t)
	if err := t.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if t.OnCluster != nil {
		if err := t.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTypeWithParams(s *TypeWithParams) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor.Self); err != nil {
		return err
	}
	for _, param := range s.Params {
		if err := param.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitUUID(u *UUID) error {
	visitor.enter(u)
	defer visitor.leave(u)
	return nil
}

func (visitor DefaultASTVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitUseExpr(u *UseStmt) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Database.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitUsingExpr(u *UsingClause) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Using.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitWhenExpr(w *WhenClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.When.Accept(visitor.Self); err != nil {
		return err
	}
	if err := w.Then.Accept(visitor.Self); err != nil {
		return err
	}
	if w.Else != nil {
		if err := w.Else.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWhereExpr(w *WhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowExpr(w *WindowClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if w.WindowExpr != nil {
		if err := w.WindowExpr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if w.Name != nil {
		if err := w.Name.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowConditionExpr(w *WindowExpr) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if w.PartitionBy != nil {
		if err := w.PartitionBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if w.OrderBy != nil {
		if err := w.OrderBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if w.Frame != nil {
		if err := w.Frame.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowFrameExpr(f *WindowFrameClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Extend.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Number.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Function.Accept(visitor.Self); err != nil {
		return err
	}
	if err := w.OverExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWithExpr(w *WithClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	for _, cte := range w.CTEs {
		if err := cte.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWithTimeoutExpr(w *WithTimeoutClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Number.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) enter(_ Expr) {}

func (visitor DefaultASTVisitor) leave(_ Expr) {}
