package parser

type DefaultASTVisitor struct {
	Self ASTVisitor
}

func (visitor DefaultASTVisitor) VisitAliasExpr(a *AliasExpr) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.Alias.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterRole(a *AlterRole) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)

	if a.TableIdentifier != nil {
		if err := a.TableIdentifier.Accept(visitor.Self); err != nil {
			return err
		}
	}

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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.ColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.IndexName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.Partition.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.ProjectionName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if a.Partition != nil {
		if err := a.Partition.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableMaterializeIndex(a *AlterTableMaterializeIndex) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
func (visitor DefaultASTVisitor) VisitAlterTableModifyQuery(a *AlterTableModifyQuery) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.SelectExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableModifySetting(a *AlterTableModifySetting) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	for _, setting := range a.Settings {
		if err := setting.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableResetSetting(a *AlterTableResetSetting) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	for _, setting := range a.Settings {
		if err := setting.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.TTL.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableRenameColumn(a *AlterTableRenameColumn) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.OldColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.NewColumnName.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.Partition.Accept(visitor.Self); err != nil {
		return err
	}
	if err := a.Table.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitArrayJoinExpr(a *ArrayJoinClause) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitArrayParamList(a *ArrayParamList) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.Items.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

// nolint: funlen

// START, STOP
// REPLICATED, DISTRIBUTED

func (visitor DefaultASTVisitor) VisitValuesExpr(v *AssignmentValues) error {
	visitor.Enter(v)
	defer visitor.Leave(v)
	for _, value := range v.Values {
		if err := value.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitBetweenClause(f *BetweenClause) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	if f.Expr != nil {
		if err := f.Expr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if err := f.Between.Accept(visitor.Self); err != nil {
		return err
	}
	if err := f.And.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitBinaryExpr(p *BinaryOperation) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
	if err := p.LeftExpr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := p.RightExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCTEExpr(c *CTEStmt) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.Alias.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitCaseExpr(c *CaseExpr) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.AsType.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCheckExpr(c *CheckStmt) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(o)
	defer visitor.Leave(o)
	if err := o.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitColumnArgList(c *ColumnArgList) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnDef(c *ColumnDef) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
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

func (visitor DefaultASTVisitor) VisitRefreshExpr(r *RefreshExpr) error {
	visitor.Enter(r)
	defer visitor.Leave(r)
	if r.Interval != nil {
		if err := r.Interval.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if r.Offset != nil {
		if err := r.Offset.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTypedPlaceholder(t *TypedPlaceholder) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	if err := t.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.Type.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitColumnExprList(c *ColumnExprList) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
	for _, item := range c.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnIdentifier(c *ColumnIdentifier) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
	for i := range c.ColumnNames {
		if err := c.ColumnNames[i].Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitComplexType(c *ComplexType) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
	if err := c.Constraint.Accept(visitor.Self); err != nil {
		return err
	}
	if err := c.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitCreateDatabase(c *CreateDatabase) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	if c.Comment != nil {
		if err := c.Comment.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateFunction(c *CreateFunction) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if c.OnCluster != nil {
		if err := c.OnCluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Refresh != nil {
		if err := c.Refresh.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.RandomizeFor != nil {
		if err := c.RandomizeFor.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, dep := range c.DependsOn {
		if err := dep.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Settings != nil {
		if err := c.Settings.Accept(visitor.Self); err != nil {
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
	if c.Definer != nil {
		if err := c.Definer.Accept(visitor.Self); err != nil {
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

func (visitor DefaultASTVisitor) VisitAuthenticationClause(a *AuthenticationClause) error {
	visitor.Enter(a)
	defer visitor.Leave(a)
	if a.AuthValue != nil {
		if err := a.AuthValue.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if a.LdapServer != nil {
		if err := a.LdapServer.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if a.KerberosRealm != nil {
		if err := a.KerberosRealm.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitHostClause(h *HostClause) error {
	visitor.Enter(h)
	defer visitor.Leave(h)
	if h.HostValue != nil {
		if err := h.HostValue.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDefaultRoleClause(d *DefaultRoleClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	for _, role := range d.Roles {
		if err := role.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitGranteesClause(g *GranteesClause) error {
	visitor.Enter(g)
	defer visitor.Leave(g)
	for _, grantee := range g.Grantees {
		if err := grantee.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, except := range g.ExceptUsers {
		if err := except.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateUser(c *CreateUser) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
	for _, userName := range c.UserNames {
		if err := userName.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Authentication != nil {
		if err := c.Authentication.Accept(visitor.Self); err != nil {
			return err
		}
	}
	for _, host := range c.Hosts {
		if err := host.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.DefaultRole != nil {
		if err := c.DefaultRole.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.DefaultDatabase != nil {
		if err := c.DefaultDatabase.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if c.Grantees != nil {
		if err := c.Grantees.Accept(visitor.Self); err != nil {
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

func (visitor DefaultASTVisitor) VisitCreateDictionary(c *CreateDictionary) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	if c.Schema != nil {
		if err := c.Schema.Accept(visitor.Self); err != nil {
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

func (visitor DefaultASTVisitor) VisitDictionarySchemaClause(d *DictionarySchemaClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	for _, attr := range d.Attributes {
		if err := attr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryAttribute(d *DictionaryAttribute) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := d.Type.Accept(visitor.Self); err != nil {
		return err
	}
	if d.Default != nil {
		if err := d.Default.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Expression != nil {
		if err := d.Expression.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryEngineClause(d *DictionaryEngineClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if d.PrimaryKey != nil {
		if err := d.PrimaryKey.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Source != nil {
		if err := d.Source.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Lifetime != nil {
		if err := d.Lifetime.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Layout != nil {
		if err := d.Layout.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Range != nil {
		if err := d.Range.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Settings != nil {
		if err := d.Settings.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryPrimaryKeyClause(d *DictionaryPrimaryKeyClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Keys.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionarySourceClause(d *DictionarySourceClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Source.Accept(visitor.Self); err != nil {
		return err
	}
	for _, arg := range d.Args {
		if err := arg.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryArgExpr(d *DictionaryArgExpr) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := d.Value.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryLifetimeClause(d *DictionaryLifetimeClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if d.Value != nil {
		if err := d.Value.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Min != nil {
		if err := d.Min.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if d.Max != nil {
		if err := d.Max.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryLayoutClause(d *DictionaryLayoutClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Layout.Accept(visitor.Self); err != nil {
		return err
	}
	for _, arg := range d.Args {
		if err := arg.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDictionaryRangeClause(d *DictionaryRangeClause) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Min.Accept(visitor.Self); err != nil {
		return err
	}
	if err := d.Max.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitShowExpr(s *ShowStmt) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if s.Target != nil {
		if err := s.Target.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.LikePattern != nil {
		if err := s.LikePattern.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Limit != nil {
		if err := s.Limit.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.OutFile != nil {
		if err := s.OutFile.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if s.Format != nil {
		if err := s.Format.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDescribeExpr(d *DescribeStmt) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.Target.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitPath(p *Path) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
	for _, ident := range p.Fields {
		if err := ident.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTargetPairExpr(t *TargetPair) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	if err := t.Old.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.New.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitCreateRole(c *CreateRole) error {
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
	if c.Name != nil {
		if err := c.Name.Accept(visitor.Self); err != nil {
			return err
		}
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
	visitor.Enter(c)
	defer visitor.Leave(c)
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
	visitor.Enter(d)
	defer visitor.Leave(d)
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
	visitor.Enter(d)
	defer visitor.Leave(d)
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
	visitor.Enter(d)
	defer visitor.Leave(d)
	if err := d.TableIdentifier.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitDropDatabase(d *DropDatabase) error {
	visitor.Enter(d)
	defer visitor.Leave(d)
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
	visitor.Enter(d)
	defer visitor.Leave(d)
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
	visitor.Enter(d)
	defer visitor.Leave(d)
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
	visitor.Enter(e)
	defer visitor.Leave(e)
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
	visitor.Enter(e)
	defer visitor.Leave(e)
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
	visitor.Enter(e)
	defer visitor.Leave(e)
	if err := e.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := e.Value.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitExplainExpr(e *ExplainStmt) error {
	visitor.Enter(e)
	defer visitor.Leave(e)
	if err := e.Statement.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitExtractExpr(e *ExtractExpr) error {
	visitor.Enter(e)
	defer visitor.Leave(e)
	if err := e.FromExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitFormatExpr(f *FormatClause) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	if err := f.Format.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitFromExpr(f *FromClause) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	if err := f.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	if err := f.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := f.Params.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitGlobalInExpr(g *GlobalInOperation) error {
	visitor.Enter(g)
	defer visitor.Leave(g)
	if err := g.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitGrantPrivilegeExpr(g *GrantPrivilegeStmt) error {
	visitor.Enter(g)
	defer visitor.Leave(g)
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
	visitor.Enter(g)
	defer visitor.Leave(g)
	if g.Expr != nil {
		if err := g.Expr.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitHavingExpr(h *HavingClause) error {
	visitor.Enter(h)
	defer visitor.Leave(h)
	if err := h.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitIdent(i *Ident) error {
	visitor.Enter(i)
	defer visitor.Leave(i)
	return nil
}
func (visitor DefaultASTVisitor) VisitIndexOperation(i *IndexOperation) error {
	visitor.Enter(i)
	defer visitor.Leave(i)
	if err := i.Object.Accept(visitor.Self); err != nil {
		return err
	}
	if err := i.Index.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitInsertExpr(i *InsertStmt) error {
	visitor.Enter(i)
	defer visitor.Leave(i)
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
	visitor.Enter(i)
	defer visitor.Leave(i)
	if err := i.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if err := i.Unit.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitJSONType(j *JSONType) error {
	visitor.Enter(j)
	defer visitor.Leave(j)
	if err := j.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitJoinConstraintExpr(j *JoinConstraintClause) error {
	visitor.Enter(j)
	defer visitor.Leave(j)
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
	visitor.Enter(j)
	defer visitor.Leave(j)
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
	visitor.Enter(j)
	defer visitor.Leave(j)
	if err := j.Table.Accept(visitor.Self); err != nil {
		return err
	}
	if j.SampleRatio != nil {
		return j.SampleRatio.Accept(visitor.Self)
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitLimitByExpr(l *LimitByClause) error {
	visitor.Enter(l)
	defer visitor.Leave(l)
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
	visitor.Enter(l)
	defer visitor.Leave(l)
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
	visitor.Enter(m)
	defer visitor.Leave(m)
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
	visitor.Enter(n)
	defer visitor.Leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
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
	visitor.Enter(n)
	defer visitor.Leave(n)
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
	visitor.Enter(n)
	defer visitor.Leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
	if err := n.NullLiteral.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitNullLiteral(n *NullLiteral) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
	return nil
}

func (visitor DefaultASTVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
	return nil
}
func (visitor DefaultASTVisitor) VisitObjectParams(o *ObjectParams) error {
	visitor.Enter(o)
	defer visitor.Leave(o)
	if err := o.Object.Accept(visitor.Self); err != nil {
		return err
	}
	if err := o.Params.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitOnExpr(o *OnClause) error {
	visitor.Enter(o)
	defer visitor.Leave(o)
	if err := o.On.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitOperationExpr(o *OperationExpr) error {
	visitor.Enter(o)
	defer visitor.Leave(o)
	return nil
}

func (visitor DefaultASTVisitor) VisitOptimizeExpr(o *OptimizeStmt) error {
	visitor.Enter(o)
	defer visitor.Leave(o)
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
	visitor.Enter(o)
	defer visitor.Leave(o)
	for _, item := range o.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitOrderByExpr(o *OrderExpr) error {
	visitor.Enter(o)
	defer visitor.Leave(o)
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
	visitor.Enter(f)
	defer visitor.Leave(f)
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
	visitor.Enter(p)
	defer visitor.Leave(p)
	if err := p.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitPartitionExpr(p *PartitionClause) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
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
	visitor.Enter(p)
	defer visitor.Leave(p)
	return nil
}

func (visitor DefaultASTVisitor) VisitPrewhereExpr(w *PrewhereClause) error {
	visitor.Enter(w)
	defer visitor.Leave(w)
	if err := w.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitPrimaryKeyExpr(p *PrimaryKeyClause) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
	if err := p.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitPrivilegeExpr(p *PrivilegeClause) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
	if p.Params != nil {
		if err := p.Params.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitProjectionOrderBy(p *ProjectionOrderByClause) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
	return nil
}

func (visitor DefaultASTVisitor) VisitProjectionSelect(p *ProjectionSelectStmt) error {
	visitor.Enter(p)
	defer visitor.Leave(p)
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
	visitor.Enter(c)
	defer visitor.Leave(c)
	if err := c.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitQueryParam(q *QueryParam) error {
	visitor.Enter(q)
	defer visitor.Leave(q)
	if err := q.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := q.Type.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitRatioExpr(r *RatioExpr) error {
	visitor.Enter(r)
	defer visitor.Leave(r)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
	if err := a.PropertyType.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitRenameStmt(r *RenameStmt) error {
	visitor.Enter(r)
	defer visitor.Leave(r)
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
	visitor.Enter(r)
	defer visitor.Leave(r)
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
	visitor.Enter(r)
	defer visitor.Leave(r)
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
	visitor.Enter(r)
	defer visitor.Leave(r)
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
	visitor.Enter(s)
	defer visitor.Leave(s)
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSampleRatioExpr(s *SampleClause) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
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
	visitor.Enter(s)
	defer visitor.Leave(s)
	if err := s.Name.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSelectItem(s *SelectItem) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
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
	visitor.Enter(s)
	defer visitor.Leave(s)
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
	visitor.Enter(s)
	defer visitor.Leave(s)
	if err := s.Settings.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSettingsExpr(s *SettingExpr) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if err := s.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSettingPair(s *SettingPair) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
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
	visitor.Enter(s)
	defer visitor.Leave(s)
	for _, item := range s.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitStringLiteral(s *StringLiteral) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	return nil
}

func (visitor DefaultASTVisitor) VisitSubQueryExpr(s *SubQuery) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if s.Select != nil {
		if err := s.Select.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemCtrlExpr(s *SystemCtrlExpr) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if s.Cluster != nil {
		if err := s.Cluster.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	return nil
}
func (visitor DefaultASTVisitor) VisitSystemFlushExpr(s *SystemFlushExpr) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if s.Distributed != nil {
		if err := s.Distributed.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemReloadExpr(s *SystemReloadExpr) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if s.Dictionary != nil {
		if err := s.Dictionary.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitSystemExpr(s *SystemStmt) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if err := s.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitSystemSyncExpr(s *SystemSyncExpr) error {
	visitor.Enter(s)
	defer visitor.Leave(s)
	if err := s.Cluster.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTTLExprList(t *TTLClause) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	for _, item := range t.Items {
		if err := item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTTLExpr(t *TTLExpr) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	if err := t.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	if t.Policy != nil {
		if err := t.Policy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTTLPolicy(t *TTLPolicy) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	if t.Item != nil {
		if err := t.Item.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if t.Where != nil {
		if err := t.Where.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if t.GroupBy != nil {
		if err := t.GroupBy.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTTLPolicyRule(t *TTLPolicyRule) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	if t.ToVolume != nil {
		if err := t.ToVolume.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if t.ToDisk != nil {
		if err := t.ToDisk.Accept(visitor.Self); err != nil {
			return err
		}
	}
	if t.Action != nil {
		if err := t.Action.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTTLPolicyItemAction(t *TTLPolicyRuleAction) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	if t.Codec != nil {
		if err := t.Codec.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableArgListExpr(t *TableArgListExpr) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
	for _, arg := range t.Args {
		if err := arg.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTableExpr(t *TableExpr) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
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
	visitor.Enter(t)
	defer visitor.Leave(t)
	if err := t.Name.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.Args.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitTableIdentifier(t *TableIdentifier) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
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
	visitor.Enter(a)
	defer visitor.Leave(a)
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
	visitor.Enter(t)
	defer visitor.Leave(t)
	if err := t.Identifier.Accept(visitor.Self); err != nil {
		return err
	}
	if err := t.Select.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTableSchemaExpr(t *TableSchemaClause) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
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
	visitor.Enter(t)
	defer visitor.Leave(t)
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
	visitor.Enter(t)
	defer visitor.Leave(t)
	if err := t.Number.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitTruncateTable(t *TruncateTable) error {
	visitor.Enter(t)
	defer visitor.Leave(t)
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
	visitor.Enter(s)
	defer visitor.Leave(s)
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
	visitor.Enter(u)
	defer visitor.Leave(u)
	return nil
}

func (visitor DefaultASTVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	visitor.Enter(n)
	defer visitor.Leave(n)
	if err := n.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitUseExpr(u *UseStmt) error {
	visitor.Enter(u)
	defer visitor.Leave(u)
	if err := u.Database.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitUsingExpr(u *UsingClause) error {
	visitor.Enter(u)
	defer visitor.Leave(u)
	if err := u.Using.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitWhenExpr(w *WhenClause) error {
	visitor.Enter(w)
	defer visitor.Leave(w)
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
	visitor.Enter(w)
	defer visitor.Leave(w)
	if err := w.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowExpr(w *WindowClause) error {
	visitor.Enter(w)
	defer visitor.Leave(w)
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
	visitor.Enter(w)
	defer visitor.Leave(w)
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
	visitor.Enter(f)
	defer visitor.Leave(f)
	if err := f.Extend.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	if err := f.Expr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	if err := f.Number.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	visitor.Enter(f)
	defer visitor.Leave(f)
	return nil
}
func (visitor DefaultASTVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	visitor.Enter(w)
	defer visitor.Leave(w)
	if err := w.Function.Accept(visitor.Self); err != nil {
		return err
	}
	if err := w.OverExpr.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWithExpr(w *WithClause) error {
	visitor.Enter(w)
	defer visitor.Leave(w)
	for _, cte := range w.CTEs {
		if err := cte.Accept(visitor.Self); err != nil {
			return err
		}
	}
	return nil
}

func (visitor DefaultASTVisitor) VisitWithTimeoutExpr(w *WithTimeoutClause) error {
	visitor.Enter(w)
	defer visitor.Leave(w)
	if err := w.Number.Accept(visitor.Self); err != nil {
		return err
	}
	return nil
}

func (visitor DefaultASTVisitor) Enter(_ Expr) {}

func (visitor DefaultASTVisitor) Leave(_ Expr) {}
