package parser

type VisitFunc func(expr Expr) error

type DefaultASTVisitor struct {
	Visit VisitFunc
}

func (v *DefaultASTVisitor) VisitOperationExpr(expr *OperationExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTernaryExpr(expr *TernaryOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitBinaryExpr(expr *BinaryOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitIndexOperation(expr *IndexOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitJoinTableExpr(expr *JoinTableExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTable(expr *AlterTable) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableAttachPartition(expr *AlterTableAttachPartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableDetachPartition(expr *AlterTableDetachPartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableDropPartition(expr *AlterTableDropPartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableFreezePartition(expr *AlterTableFreezePartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableAddColumn(expr *AlterTableAddColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableAddIndex(expr *AlterTableAddIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableAddProjection(expr *AlterTableAddProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitProjectionOrderBy(expr *ProjectionOrderByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitProjectionSelect(expr *ProjectionSelectStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableProjection(expr *TableProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableDropColumn(expr *AlterTableDropColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableDropIndex(expr *AlterTableDropIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableDropProjection(expr *AlterTableDropProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableRemoveTTL(expr *AlterTableRemoveTTL) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableClearColumn(expr *AlterTableClearColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableClearIndex(expr *AlterTableClearIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableClearProjection(expr *AlterTableClearProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableMaterializeProjection(expr *AlterTableMaterializeProjection) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableMaterializeIndex(expr *AlterTableMaterializeIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableRenameColumn(expr *AlterTableRenameColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableModifyTTL(expr *AlterTableModifyTTL) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableModifyColumn(expr *AlterTableModifyColumn) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterTableReplacePartition(expr *AlterTableReplacePartition) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitRemovePropertyType(expr *RemovePropertyType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableIndex(expr *TableIndex) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitIdent(expr *Ident) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitUUID(expr *UUID) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateDatabase(expr *CreateDatabase) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateTable(expr *CreateTable) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateMaterializedView(expr *CreateMaterializedView) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateView(expr *CreateView) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateFunction(expr *CreateFunction) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitRoleName(expr *RoleName) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSettingPair(expr *SettingPair) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitRoleSetting(expr *RoleSetting) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateRole(expr *CreateRole) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAlterRole(expr *AlterRole) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitRoleRenamePair(expr *RoleRenamePair) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitDestinationExpr(expr *DestinationClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitConstraintExpr(expr *ConstraintClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNullLiteral(expr *NullLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNotNullLiteral(expr *NotNullLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNestedIdentifier(expr *NestedIdentifier) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnIdentifier(expr *ColumnIdentifier) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableIdentifier(expr *TableIdentifier) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableSchemaExpr(expr *TableSchemaClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableArgListExpr(expr *TableArgListExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableFunctionExpr(expr *TableFunctionExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitOnClusterExpr(expr *ClusterClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPartitionExpr(expr *PartitionClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPartitionByExpr(expr *PartitionByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPrimaryKeyExpr(expr *PrimaryKeyClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSampleByExpr(expr *SampleByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTTLExpr(expr *TTLExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTTLExprList(expr *TTLClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitOrderByExpr(expr *OrderExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitOrderByListExpr(expr *OrderByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSettingsExpr(expr *SettingExprList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSettingsExprList(expr *SettingsClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitParamExprList(expr *ParamExprList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitArrayParamList(expr *ArrayParamList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitQueryParam(expr *QueryParam) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitMapLiteral(expr *MapLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitObjectParams(expr *ObjectParams) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitFunctionExpr(expr *FunctionExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowFunctionExpr(expr *WindowFunctionExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnDef(expr *ColumnDef) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnExpr(expr *ColumnExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitScalarType(expr *ScalarType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitJSONType(expr *JSONType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPropertyType(expr *PropertyType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTypeWithParams(expr *TypeWithParams) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitComplexType(expr *ComplexType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNestedType(expr *NestedType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCompressionCodec(expr *CompressionCodec) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNumberLiteral(expr *NumberLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitStringLiteral(expr *StringLiteral) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitRatioExpr(expr *RatioExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitEnumValue(expr *EnumValue) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitEnumType(expr *EnumType) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitIntervalExpr(expr *IntervalExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitEngineExpr(expr *EngineExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnTypeExpr(expr *ColumnTypeExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnArgList(expr *ColumnArgList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnExprList(expr *ColumnExprList) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWhenExpr(expr *WhenClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCaseExpr(expr *CaseExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCastExpr(expr *CastExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWithExpr(expr *WithClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTopExpr(expr *TopClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCreateLiveView(expr *CreateLiveView) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWithTimeoutExpr(expr *WithTimeoutClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTableExpr(expr *TableExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitOnExpr(expr *OnClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitUsingExpr(expr *UsingClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitJoinExpr(expr *JoinExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitJoinConstraintExpr(expr *JoinConstraintClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitFromExpr(expr *FromClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitIsNullExpr(expr *IsNullExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitIsNotNullExpr(expr *IsNotNullExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitAliasExpr(expr *AliasExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWhereExpr(expr *WhereClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPrewhereExpr(expr *PrewhereClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitGroupByExpr(expr *GroupByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitHavingExpr(expr *HavingClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitLimitExpr(expr *LimitClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitLimitByExpr(expr *LimitByClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowConditionExpr(expr *WindowExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowExpr(expr *WindowClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowFrameExpr(expr *WindowFrameClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowFrameExtendExpr(expr *WindowFrameExtendExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitBetweenClause(expr *BetweenClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowFrameCurrentRow(expr *WindowFrameCurrentRow) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowFrameUnbounded(expr *WindowFrameUnbounded) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitWindowFrameNumber(expr *WindowFrameNumber) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitArrayJoinExpr(expr *ArrayJoinClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSelectQuery(expr *SelectQuery) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSubQueryExpr(expr *SubQuery) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNotExpr(expr *NotExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitNegateExpr(expr *NegateExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitGlobalInExpr(expr *GlobalInOperation) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitExtractExpr(expr *ExtractExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitDropDatabase(expr *DropDatabase) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitDropStmt(expr *DropStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitDropUserOrRole(expr *DropUserOrRole) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitUseExpr(expr *UseStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCTEExpr(expr *CTEStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSetExpr(expr *SetStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitFormatExpr(expr *FormatClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitOptimizeExpr(expr *OptimizeStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitDeduplicateExpr(expr *DeduplicateClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSystemExpr(expr *SystemStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSystemFlushExpr(expr *SystemFlushExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSystemReloadExpr(expr *SystemReloadExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSystemSyncExpr(expr *SystemSyncExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSystemCtrlExpr(expr *SystemCtrlExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSystemDropExpr(expr *SystemDropExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitTruncateTable(expr *TruncateTable) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSampleRatioExpr(expr *SampleClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPlaceHolderExpr(expr *PlaceHolder) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitDeleteFromExpr(expr *DeleteClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitColumnNamesExpr(expr *ColumnNamesExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitValuesExpr(expr *AssignmentValues) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitInsertExpr(expr *InsertStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitCheckExpr(expr *CheckStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitUnaryExpr(expr *UnaryExpr) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitRenameStmt(expr *RenameStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitExplainExpr(expr *ExplainStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitPrivilegeExpr(expr *PrivilegeClause) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitGrantPrivilegeExpr(expr *GrantPrivilegeStmt) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) VisitSelectItem(expr *SelectItem) error {
	if v.Visit != nil {
		return v.Visit(expr)
	}
	return nil
}

func (v *DefaultASTVisitor) enter(expr Expr) {}

func (v *DefaultASTVisitor) leave(expr Expr) {}
