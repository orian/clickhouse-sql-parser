package parser

type ASTVisitor interface {
	VisitOperationExpr(expr *OperationExpr) error
	VisitTernaryExpr(expr *TernaryOperation) error
	VisitBinaryExpr(expr *BinaryOperation) error
	VisitIndexOperation(expr *IndexOperation) error
	VisitAlterTable(expr *AlterTable) error
	VisitAlterTableAttachPartition(expr *AlterTableAttachPartition) error
	VisitAlterTableDetachPartition(expr *AlterTableDetachPartition) error
	VisitAlterTableDropPartition(expr *AlterTableDropPartition) error
	VisitAlterTableFreezePartition(expr *AlterTableFreezePartition) error
	VisitAlterTableAddColumn(expr *AlterTableAddColumn) error
	VisitAlterTableAddIndex(expr *AlterTableAddIndex) error
	VisitAlterTableAddProjection(expr *AlterTableAddProjection) error
	VisitTableProjection(expr *TableProjection) error
	VisitProjectionOrderBy(expr *ProjectionOrderByClause) error
	VisitProjectionSelect(expr *ProjectionSelectStmt) error
	VisitAlterTableDropColumn(expr *AlterTableDropColumn) error
	VisitAlterTableDropIndex(expr *AlterTableDropIndex) error
	VisitAlterTableDropProjection(expr *AlterTableDropProjection) error
	VisitAlterTableRemoveTTL(expr *AlterTableRemoveTTL) error
	VisitAlterTableClearColumn(expr *AlterTableClearColumn) error
	VisitAlterTableClearIndex(expr *AlterTableClearIndex) error
	VisitAlterTableClearProjection(expr *AlterTableClearProjection) error
	VisitAlterTableMaterializeIndex(expr *AlterTableMaterializeIndex) error
	VisitAlterTableMaterializeProjection(expr *AlterTableMaterializeProjection) error
	VisitAlterTableRenameColumn(expr *AlterTableRenameColumn) error
	VisitAlterTableModifyTTL(expr *AlterTableModifyTTL) error
	VisitAlterTableModifyColumn(expr *AlterTableModifyColumn) error
	VisitAlterTableReplacePartition(expr *AlterTableReplacePartition) error
	VisitRemovePropertyType(expr *RemovePropertyType) error
	VisitTableIndex(expr *TableIndex) error
	VisitIdent(expr *Ident) error
	VisitUUID(expr *UUID) error
	VisitCreateDatabase(expr *CreateDatabase) error
	VisitCreateTable(expr *CreateTable) error
	VisitCreateMaterializedView(expr *CreateMaterializedView) error
	VisitCreateView(expr *CreateView) error
	VisitCreateFunction(expr *CreateFunction) error
	VisitRoleName(expr *RoleName) error
	VisitSettingPair(expr *SettingPair) error
	VisitRoleSetting(expr *RoleSetting) error
	VisitCreateRole(expr *CreateRole) error
	VisitAlterRole(expr *AlterRole) error
	VisitRoleRenamePair(expr *RoleRenamePair) error
	VisitDestinationExpr(expr *DestinationClause) error
	VisitConstraintExpr(expr *ConstraintClause) error
	VisitNullLiteral(expr *NullLiteral) error
	VisitNotNullLiteral(expr *NotNullLiteral) error
	VisitNestedIdentifier(expr *NestedIdentifier) error
	VisitColumnIdentifier(expr *ColumnIdentifier) error
	VisitTableIdentifier(expr *TableIdentifier) error
	VisitTableSchemaExpr(expr *TableSchemaClause) error
	VisitTableArgListExpr(expr *TableArgListExpr) error
	VisitTableFunctionExpr(expr *TableFunctionExpr) error
	VisitOnClusterExpr(expr *ClusterClause) error
	VisitPartitionExpr(expr *PartitionClause) error
	VisitPartitionByExpr(expr *PartitionByClause) error
	VisitPrimaryKeyExpr(expr *PrimaryKeyClause) error
	VisitSampleByExpr(expr *SampleByClause) error
	VisitTTLExpr(expr *TTLExpr) error
	VisitTTLExprList(expr *TTLClause) error
	VisitOrderByExpr(expr *OrderExpr) error
	VisitOrderByListExpr(expr *OrderByClause) error
	VisitSettingsExpr(expr *SettingExprList) error
	VisitSettingsExprList(expr *SettingsClause) error
	VisitParamExprList(expr *ParamExprList) error
	VisitMapLiteral(expr *MapLiteral) error
	VisitArrayParamList(expr *ArrayParamList) error
	VisitQueryParam(expr *QueryParam) error
	VisitObjectParams(expr *ObjectParams) error
	VisitFunctionExpr(expr *FunctionExpr) error
	VisitWindowFunctionExpr(expr *WindowFunctionExpr) error
	VisitColumnDef(expr *ColumnDef) error
	VisitColumnExpr(expr *ColumnExpr) error
	VisitScalarType(expr *ScalarType) error
	VisitJSONType(expr *JSONType) error
	VisitPropertyType(expr *PropertyType) error
	VisitTypeWithParams(expr *TypeWithParams) error
	VisitComplexType(expr *ComplexType) error
	VisitNestedType(expr *NestedType) error
	VisitCompressionCodec(expr *CompressionCodec) error
	VisitNumberLiteral(expr *NumberLiteral) error
	VisitStringLiteral(expr *StringLiteral) error
	VisitRatioExpr(expr *RatioExpr) error
	VisitEnumValue(expr *EnumValue) error
	VisitEnumType(expr *EnumType) error
	VisitIntervalExpr(expr *IntervalExpr) error
	VisitEngineExpr(expr *EngineExpr) error
	VisitColumnTypeExpr(expr *ColumnTypeExpr) error
	VisitColumnArgList(expr *ColumnArgList) error
	VisitColumnExprList(expr *ColumnExprList) error
	VisitWhenExpr(expr *WhenClause) error
	VisitCaseExpr(expr *CaseExpr) error
	VisitCastExpr(expr *CastExpr) error
	VisitWithExpr(expr *WithClause) error
	VisitTopExpr(expr *TopClause) error
	VisitCreateLiveView(expr *CreateLiveView) error
	VisitWithTimeoutExpr(expr *WithTimeoutClause) error
	VisitTableExpr(expr *TableExpr) error
	VisitOnExpr(expr *OnClause) error
	VisitUsingExpr(expr *UsingClause) error
	VisitJoinExpr(expr *JoinExpr) error
	VisitJoinConstraintExpr(expr *JoinConstraintClause) error
	VisitJoinTableExpr(expr *JoinTableExpr) error
	VisitFromExpr(expr *FromClause) error
	VisitIsNullExpr(expr *IsNullExpr) error
	VisitIsNotNullExpr(expr *IsNotNullExpr) error
	VisitAliasExpr(expr *AliasExpr) error
	VisitWhereExpr(expr *WhereClause) error
	VisitPrewhereExpr(expr *PrewhereClause) error
	VisitGroupByExpr(expr *GroupByClause) error
	VisitHavingExpr(expr *HavingClause) error
	VisitLimitExpr(expr *LimitClause) error
	VisitLimitByExpr(expr *LimitByClause) error
	VisitWindowConditionExpr(expr *WindowExpr) error
	VisitWindowExpr(expr *WindowClause) error
	VisitWindowFrameExpr(expr *WindowFrameClause) error
	VisitWindowFrameExtendExpr(expr *WindowFrameExtendExpr) error
	VisitBetweenClause(expr *BetweenClause) error
	VisitWindowFrameCurrentRow(expr *WindowFrameCurrentRow) error
	VisitWindowFrameUnbounded(expr *WindowFrameUnbounded) error
	VisitWindowFrameNumber(expr *WindowFrameNumber) error
	VisitArrayJoinExpr(expr *ArrayJoinClause) error
	VisitSelectQuery(expr *SelectQuery) error
	VisitSubQueryExpr(expr *SubQuery) error
	VisitNotExpr(expr *NotExpr) error
	VisitNegateExpr(expr *NegateExpr) error
	VisitGlobalInExpr(expr *GlobalInOperation) error
	VisitExtractExpr(expr *ExtractExpr) error
	VisitDropDatabase(expr *DropDatabase) error
	VisitDropStmt(expr *DropStmt) error
	VisitDropUserOrRole(expr *DropUserOrRole) error
	VisitUseExpr(expr *UseStmt) error
	VisitCTEExpr(expr *CTEStmt) error
	VisitSetExpr(expr *SetStmt) error
	VisitFormatExpr(expr *FormatClause) error
	VisitOptimizeExpr(expr *OptimizeStmt) error
	VisitDeduplicateExpr(expr *DeduplicateClause) error
	VisitSystemExpr(expr *SystemStmt) error
	VisitSystemFlushExpr(expr *SystemFlushExpr) error
	VisitSystemReloadExpr(expr *SystemReloadExpr) error
	VisitSystemSyncExpr(expr *SystemSyncExpr) error
	VisitSystemCtrlExpr(expr *SystemCtrlExpr) error
	VisitSystemDropExpr(expr *SystemDropExpr) error
	VisitTruncateTable(expr *TruncateTable) error
	VisitSampleRatioExpr(expr *SampleClause) error
	VisitPlaceHolderExpr(expr *PlaceHolder) error
	VisitDeleteFromExpr(expr *DeleteClause) error
	VisitColumnNamesExpr(expr *ColumnNamesExpr) error
	VisitValuesExpr(expr *AssignmentValues) error
	VisitInsertExpr(expr *InsertStmt) error
	VisitCheckExpr(expr *CheckStmt) error
	VisitUnaryExpr(expr *UnaryExpr) error
	VisitRenameStmt(expr *RenameStmt) error
	VisitExplainExpr(expr *ExplainStmt) error
	VisitPrivilegeExpr(expr *PrivilegeClause) error
	VisitGrantPrivilegeExpr(expr *GrantPrivilegeStmt) error
	VisitSelectItem(expr *SelectItem) error

	enter(expr Expr)
	leave(expr Expr)
}
