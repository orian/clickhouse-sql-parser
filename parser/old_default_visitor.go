package parser

type VisitFunc func(expr Expr) error

type DeprecatedDefaultASTVisitor struct {
	Visit VisitFunc
}

func (v *DeprecatedDefaultASTVisitor) VisitOperationExpr(expr *OperationExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTernaryExpr(expr *TernaryOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitBinaryExpr(expr *BinaryOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitIndexOperation(expr *IndexOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitJoinTableExpr(expr *JoinTableExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTable(expr *AlterTable) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableAttachPartition(expr *AlterTableAttachPartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableDetachPartition(expr *AlterTableDetachPartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableDropPartition(expr *AlterTableDropPartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableFreezePartition(expr *AlterTableFreezePartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableAddColumn(expr *AlterTableAddColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableAddIndex(expr *AlterTableAddIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableAddProjection(expr *AlterTableAddProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitProjectionOrderBy(expr *ProjectionOrderByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitProjectionSelect(expr *ProjectionSelectStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableProjection(expr *TableProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableDropColumn(expr *AlterTableDropColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableDropIndex(expr *AlterTableDropIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableDropProjection(expr *AlterTableDropProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableRemoveTTL(expr *AlterTableRemoveTTL) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableClearColumn(expr *AlterTableClearColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableClearIndex(expr *AlterTableClearIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableClearProjection(expr *AlterTableClearProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableMaterializeProjection(expr *AlterTableMaterializeProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableMaterializeIndex(expr *AlterTableMaterializeIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableRenameColumn(expr *AlterTableRenameColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableModifyTTL(expr *AlterTableModifyTTL) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableModifyColumn(expr *AlterTableModifyColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterTableReplacePartition(expr *AlterTableReplacePartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitRemovePropertyType(expr *RemovePropertyType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableIndex(expr *TableIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitIdent(expr *Ident) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitUUID(expr *UUID) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateDatabase(expr *CreateDatabase) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateTable(expr *CreateTable) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateMaterializedView(expr *CreateMaterializedView) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateView(expr *CreateView) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateFunction(expr *CreateFunction) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitRoleName(expr *RoleName) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSettingPair(expr *SettingPair) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitRoleSetting(expr *RoleSetting) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateRole(expr *CreateRole) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAlterRole(expr *AlterRole) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitRoleRenamePair(expr *RoleRenamePair) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitDestinationExpr(expr *DestinationClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitConstraintExpr(expr *ConstraintClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNullLiteral(expr *NullLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNotNullLiteral(expr *NotNullLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNestedIdentifier(expr *NestedIdentifier) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnIdentifier(expr *ColumnIdentifier) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableIdentifier(expr *TableIdentifier) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableSchemaExpr(expr *TableSchemaClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableArgListExpr(expr *TableArgListExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableFunctionExpr(expr *TableFunctionExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitOnClusterExpr(expr *ClusterClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPartitionExpr(expr *PartitionClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPartitionByExpr(expr *PartitionByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPrimaryKeyExpr(expr *PrimaryKeyClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSampleByExpr(expr *SampleByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTTLExpr(expr *TTLExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTTLExprList(expr *TTLClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitOrderByExpr(expr *OrderExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitOrderByListExpr(expr *OrderByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSettingsExpr(expr *SettingExprList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSettingsExprList(expr *SettingsClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitParamExprList(expr *ParamExprList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitArrayParamList(expr *ArrayParamList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitQueryParam(expr *QueryParam) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitMapLiteral(expr *MapLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitObjectParams(expr *ObjectParams) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitFunctionExpr(expr *FunctionExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowFunctionExpr(expr *WindowFunctionExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnDef(expr *ColumnDef) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnExpr(expr *ColumnExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitScalarType(expr *ScalarType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitJSONType(expr *JSONType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPropertyType(expr *PropertyType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTypeWithParams(expr *TypeWithParams) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitComplexType(expr *ComplexType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNestedType(expr *NestedType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCompressionCodec(expr *CompressionCodec) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNumberLiteral(expr *NumberLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitStringLiteral(expr *StringLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitRatioExpr(expr *RatioExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitEnumValue(expr *EnumValue) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitEnumType(expr *EnumType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitIntervalExpr(expr *IntervalExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitEngineExpr(expr *EngineExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnTypeExpr(expr *ColumnTypeExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnArgList(expr *ColumnArgList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnExprList(expr *ColumnExprList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWhenExpr(expr *WhenClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCaseExpr(expr *CaseExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCastExpr(expr *CastExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWithExpr(expr *WithClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTopExpr(expr *TopClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCreateLiveView(expr *CreateLiveView) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWithTimeoutExpr(expr *WithTimeoutClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTableExpr(expr *TableExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitOnExpr(expr *OnClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitUsingExpr(expr *UsingClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitJoinExpr(expr *JoinExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitJoinConstraintExpr(expr *JoinConstraintClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitFromExpr(expr *FromClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitIsNullExpr(expr *IsNullExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitIsNotNullExpr(expr *IsNotNullExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitAliasExpr(expr *AliasExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWhereExpr(expr *WhereClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPrewhereExpr(expr *PrewhereClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitGroupByExpr(expr *GroupByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitHavingExpr(expr *HavingClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitLimitExpr(expr *LimitClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitLimitByExpr(expr *LimitByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowConditionExpr(expr *WindowExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowExpr(expr *WindowClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowFrameExpr(expr *WindowFrameClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowFrameExtendExpr(expr *WindowFrameExtendExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitBetweenClause(expr *BetweenClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowFrameCurrentRow(expr *WindowFrameCurrentRow) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowFrameUnbounded(expr *WindowFrameUnbounded) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitWindowFrameNumber(expr *WindowFrameNumber) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitArrayJoinExpr(expr *ArrayJoinClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSelectQuery(expr *SelectQuery) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSubQueryExpr(expr *SubQuery) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNotExpr(expr *NotExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitNegateExpr(expr *NegateExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitGlobalInExpr(expr *GlobalInOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitExtractExpr(expr *ExtractExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitDropDatabase(expr *DropDatabase) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitDropStmt(expr *DropStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitDropUserOrRole(expr *DropUserOrRole) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitUseExpr(expr *UseStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCTEExpr(expr *CTEStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSetExpr(expr *SetStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitFormatExpr(expr *FormatClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitOptimizeExpr(expr *OptimizeStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitDeduplicateExpr(expr *DeduplicateClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSystemExpr(expr *SystemStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSystemFlushExpr(expr *SystemFlushExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSystemReloadExpr(expr *SystemReloadExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSystemSyncExpr(expr *SystemSyncExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSystemCtrlExpr(expr *SystemCtrlExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSystemDropExpr(expr *SystemDropExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitTruncateTable(expr *TruncateTable) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSampleRatioExpr(expr *SampleClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPlaceHolderExpr(expr *PlaceHolder) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitDeleteFromExpr(expr *DeleteClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitColumnNamesExpr(expr *ColumnNamesExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitValuesExpr(expr *AssignmentValues) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitInsertExpr(expr *InsertStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitCheckExpr(expr *CheckStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitUnaryExpr(expr *UnaryExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitRenameStmt(expr *RenameStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitExplainExpr(expr *ExplainStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitPrivilegeExpr(expr *PrivilegeClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitGrantPrivilegeExpr(expr *GrantPrivilegeStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) VisitSelectItem(expr *SelectItem) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DeprecatedDefaultASTVisitor) enter(expr Expr) {}

func (v *DeprecatedDefaultASTVisitor) leave(expr Expr) {}
