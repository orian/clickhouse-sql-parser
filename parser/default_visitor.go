package parser

import (
	"strings"
)

type PrintVisitor struct {
	builder strings.Builder
}
type MyDefaultVisitor struct{}

func (visitor MyDefaultVisitor) enter(expr Expr) {
}
func (visitor MyDefaultVisitor) leave(expr Expr) {
	//TODO implement me

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

func (p *PrintVisitor) VisitOperationExpr(o *OperationExpr) error {
	p.builder.WriteString(strings.ToUpper(string(o.Kind)))
	return nil
}

func (visitor MyDefaultVisitor) VisitOperationExpr(o *OperationExpr) error {
	visitor.enter(o)
	defer visitor.leave(o)
	return nil
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

func (p *PrintVisitor) VisitIndexOperation(i *IndexOperation) error {
	builder := p.builder
	builder.WriteString(i.Object.String())
	builder.WriteString(string(i.Operation))
	builder.WriteString(i.Index.String())
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

func (visitor MyDefaultVisitor) VisitAlterTableDropPartition(a *AlterTableDropPartition) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Partition.Accept(visitor); err != nil {
		return err
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

func (p *PrintVisitor) VisitAlterTableFreezePartition(a *AlterTableFreezePartition) error {
	builder := p.builder
	builder.WriteString("FREEZE")
	if a.Partition != nil {
		builder.WriteByte(' ')
		builder.WriteString(a.Partition.String())
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

type AlterTableAddIndex struct {
	AddPos       Pos
	StatementEnd Pos

	Index       *TableIndex
	IfNotExists bool
	After       *NestedIdentifier
}

func (a *AlterTableAddIndex) Pos() Pos {
	return a.AddPos
}

func (a *AlterTableAddIndex) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableAddIndex) AlterType() error {
	return "ADD_INDEX"
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

type ProjectionOrderByClause struct {
	OrderByPos Pos
	Columns    *ColumnExprList
}

func (p *ProjectionOrderByClause) Pos() Pos {
	return p.OrderByPos
}

func (p *ProjectionOrderByClause) End() Pos {
	return p.Columns.End()
}

func (p *PrintVisitor) VisitProjectionOrderByClause(p *ProjectionOrderByClause) error {
	builder := p.builder
	builder.WriteString("ORDER BY ")
	builder.WriteString(p.Columns.String())
	return nil
}

func (visitor MyDefaultVisitor) VisitProjectionOrderByClause(p *ProjectionOrderByClause) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

type ProjectionSelectStmt struct {
	LeftParenPos  Pos
	RightParenPos Pos
	With          *WithClause
	SelectColumns *ColumnExprList
	GroupBy       *GroupByClause
	OrderBy       *ProjectionOrderByClause
}

func (p *ProjectionSelectStmt) Pos() Pos {
	return p.LeftParenPos

}

func (p *ProjectionSelectStmt) End() Pos {
	return p.RightParenPos
}

func (p *PrintVisitor) VisitProjectionSelectStmt(p *ProjectionSelectStmt) error {
	builder := p.builder
	builder.WriteString("(")
	if p.With != nil {
		builder.WriteString(p.With.String())
		builder.WriteByte(' ')
	}
	builder.WriteString("SELECT ")
	builder.WriteString(p.SelectColumns.String())
	if p.GroupBy != nil {
		builder.WriteString(" ")
		builder.WriteString(p.GroupBy.String())
	}
	if p.OrderBy != nil {
		builder.WriteString(" ")
		builder.WriteString(p.OrderBy.String())
	}
	builder.WriteString(")")
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

type TableProjection struct {
	ProjectionPos Pos
	Identifier    *NestedIdentifier
	Select        *ProjectionSelectStmt
}

func (t *TableProjection) Pos() Pos {
	return t.ProjectionPos
}

func (t *TableProjection) End() Pos {
	return t.Select.End()
}

func (p *PrintVisitor) VisitTableProjection(t *TableProjection) error {
	builder := p.builder
	builder.WriteString(t.Identifier.String())
	builder.WriteString(" ")
	builder.WriteString(t.Select.String())
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

type AlterTableAddProjection struct {
	AddPos       Pos
	StatementEnd Pos

	IfNotExists     bool
	TableProjection *TableProjection
	After           *NestedIdentifier
}

func (a *AlterTableAddProjection) Pos() Pos {
	return a.AddPos
}

func (a *AlterTableAddProjection) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableAddProjection) AlterType() error {
	return "ADD_PROJECTION"
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

type AlterTableDropColumn struct {
	DropPos    Pos
	ColumnName *NestedIdentifier
	IfExists   bool
}

func (a *AlterTableDropColumn) Pos() Pos {
	return a.DropPos
}

func (a *AlterTableDropColumn) End() Pos {
	return a.ColumnName.End()
}

func (a *AlterTableDropColumn) AlterType() error {
	return "DROP_COLUMN"
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

func (visitor MyDefaultVisitor) VisitAlterTableDropColumn(a *AlterTableDropColumn) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ColumnName.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type AlterTableDropIndex struct {
	DropPos   Pos
	IndexName *NestedIdentifier
	IfExists  bool
}

func (a *AlterTableDropIndex) Pos() Pos {
	return a.DropPos
}

func (a *AlterTableDropIndex) End() Pos {
	return a.IndexName.End()
}

func (a *AlterTableDropIndex) AlterType() error {
	return "DROP_INDEX"
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

func (visitor MyDefaultVisitor) VisitAlterTableDropIndex(a *AlterTableDropIndex) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.IndexName.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type AlterTableDropProjection struct {
	DropPos        Pos
	ProjectionName *NestedIdentifier
	IfExists       bool
}

func (a *AlterTableDropProjection) Pos() Pos {
	return a.DropPos
}

func (a *AlterTableDropProjection) End() Pos {
	return a.ProjectionName.End()
}

func (a *AlterTableDropProjection) AlterType() error {
	return "DROP_PROJECTION"
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

func (visitor MyDefaultVisitor) VisitAlterTableDropProjection(a *AlterTableDropProjection) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.ProjectionName.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type AlterTableRemoveTTL struct {
	RemovePos    Pos
	StatementEnd Pos
}

func (a *AlterTableRemoveTTL) Pos() Pos {
	return a.RemovePos
}

func (a *AlterTableRemoveTTL) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableRemoveTTL) AlterType() error {
	return "REMOVE_TTL"
}

func (p *PrintVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	return "REMOVE TTL"
}

func (visitor MyDefaultVisitor) VisitAlterTableRemoveTTL(a *AlterTableRemoveTTL) error {
	visitor.enter(a)
	defer visitor.leave(a)
	return nil
}

type AlterTableClearColumn struct {
	ClearPos     Pos
	StatementEnd Pos

	IfExists      bool
	ColumnName    *NestedIdentifier
	PartitionExpr *PartitionClause
}

func (a *AlterTableClearColumn) Pos() Pos {
	return a.ClearPos
}

func (a *AlterTableClearColumn) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableClearColumn) AlterType() error {
	return "CLEAR_COLUMN"
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

type AlterTableClearIndex struct {
	ClearPos     Pos
	StatementEnd Pos

	IfExists      bool
	IndexName     *NestedIdentifier
	PartitionExpr *PartitionClause
}

func (a *AlterTableClearIndex) Pos() Pos {
	return a.ClearPos
}

func (a *AlterTableClearIndex) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableClearIndex) AlterType() error {
	return "CLEAR_INDEX"
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

type AlterTableClearProjection struct {
	ClearPos     Pos
	StatementEnd Pos

	IfExists       bool
	ProjectionName *NestedIdentifier
	PartitionExpr  *PartitionClause
}

func (a *AlterTableClearProjection) Pos() Pos {
	return a.ClearPos
}

func (a *AlterTableClearProjection) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableClearProjection) AlterType() error {
	return "CLEAR_PROJECTION"
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

type AlterTableRenameColumn struct {
	RenamePos Pos

	IfExists      bool
	OldColumnName *NestedIdentifier
	NewColumnName *NestedIdentifier
}

func (a *AlterTableRenameColumn) Pos() Pos {
	return a.RenamePos
}

func (a *AlterTableRenameColumn) End() Pos {
	return a.NewColumnName.End()
}

func (a *AlterTableRenameColumn) AlterType() error {
	return "RENAME_COLUMN"
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

type AlterTableModifyTTL struct {
	ModifyPos    Pos
	StatementEnd Pos
	TTL          *TTLExpr
}

func (a *AlterTableModifyTTL) Pos() Pos {
	return a.ModifyPos
}

func (a *AlterTableModifyTTL) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableModifyTTL) AlterType() error {
	return "MODIFY_TTL"
}

func (p *PrintVisitor) VisitAlterTableModifyTTL(a *AlterTableModifyTTL) error {
	builder := p.builder
	builder.WriteString("MODIFY ")
	builder.WriteString("TTL ")
	builder.WriteString(a.TTL.String())
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

type AlterTableModifyColumn struct {
	ModifyPos    Pos
	StatementEnd Pos

	IfExists           bool
	Column             *ColumnDef
	RemovePropertyType *RemovePropertyType
}

func (a *AlterTableModifyColumn) Pos() Pos {
	return a.ModifyPos
}

func (a *AlterTableModifyColumn) End() Pos {
	return a.StatementEnd
}

func (a *AlterTableModifyColumn) AlterType() error {
	return "MODIFY_COLUMN"
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

type AlterTableReplacePartition struct {
	ReplacePos Pos
	Partition  *PartitionClause
	Table      *TableIdentifier
}

func (a *AlterTableReplacePartition) Pos() Pos {
	return a.ReplacePos
}

func (a *AlterTableReplacePartition) End() Pos {
	return a.Table.End()
}

func (a *AlterTableReplacePartition) AlterType() error {
	return "REPLACE_PARTITION"
}

func (p *PrintVisitor) VisitAlterTableReplacePartition(a *AlterTableReplacePartition) error {
	builder := p.builder
	builder.WriteString("REPLACE ")
	builder.WriteString(a.Partition.String())
	builder.WriteString(" FROM ")
	builder.WriteString(a.Table.String())
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

type RemovePropertyType struct {
	RemovePos Pos

	PropertyType Expr
}

func (a *RemovePropertyType) Pos() Pos {
	return a.RemovePos
}

func (a *RemovePropertyType) End() Pos {
	return a.PropertyType.End()
}

func (p *PrintVisitor) VisitRemovePropertyType(a *RemovePropertyType) error {
	builder := p.builder
	builder.WriteString(" REMOVE ")
	builder.WriteString(a.PropertyType.String())
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

func (p *PrintVisitor) VisitTableIndex(a *TableIndex) error {
	builder := p.builder
	builder.WriteString("INDEX")
	builder.WriteByte(' ')
	builder.WriteString(a.Name.String())
	// a.ColumnDef = *Name --- e.g. INDEX idx column TYPE ...
	// a.ColumnDef = *ParamExprList --- e.g. INDEX idx(column) TYPE ...
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

func (p *PrintVisitor) VisitIdent(i *Ident) error {
	if i.QuoteType == BackTicks {
		p.builder.WriteString("`" + i.Name + "`")
	} else if i.QuoteType == DoubleQuote {
		p.builder.WriteString(`"` + i.Name + `"`)
	}
	return nil
}

func (visitor MyDefaultVisitor) VisitIdent(i *Ident) error {
	visitor.enter(i)
	defer visitor.leave(i)
	return nil
}

type UUID struct {
	Value *StringLiteral
}

func (u *UUID) Pos() Pos {
	return u.Value.LiteralPos
}

func (u *UUID) End() Pos {
	return u.Value.LiteralEnd
}

func (p *PrintVisitor) VisitUUID(u *UUID) error {
	return "UUID " + u.Value.String()
}

func (visitor MyDefaultVisitor) VisitUUID(u *UUID) error {
	visitor.enter(u)
	defer visitor.leave(u)
	return nil
}

type CreateDatabase struct {
	CreatePos    Pos // position of CREATE keyword
	StatementEnd Pos
	Name         Expr
	IfNotExists  bool // true if 'IF NOT EXISTS' is specified
	OnCluster    *ClusterClause
	Engine       *EngineExpr
}

func (c *CreateDatabase) Pos() Pos {
	return c.CreatePos
}

func (c *CreateDatabase) End() Pos {
	return c.StatementEnd
}

func (c *CreateDatabase) Type() error {
	return "DATABASE"
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

type CreateTable struct {
	CreatePos    Pos // position of CREATE|ATTACH keyword
	StatementEnd Pos
	Name         *TableIdentifier
	IfNotExists  bool
	UUID         *UUID
	OnCluster    *ClusterClause
	TableSchema  *TableSchemaClause
	Engine       *EngineExpr
	SubQuery     *SubQuery
	HasTemporary bool
	Comment      *StringLiteral
}

func (c *CreateTable) Pos() Pos {
	return c.CreatePos
}

func (c *CreateTable) End() Pos {
	return c.StatementEnd
}

func (c *CreateTable) Type() error {
	return "CREATE TABLE"
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

type CreateMaterializedView struct {
	CreatePos    Pos // position of CREATE|ATTACH keyword
	StatementEnd Pos
	Name         *TableIdentifier
	IfNotExists  bool
	OnCluster    *ClusterClause
	Engine       *EngineExpr
	Destination  *DestinationClause
	SubQuery     *SubQuery
	Populate     bool
	Comment      *StringLiteral
}

func (c *CreateMaterializedView) Pos() Pos {
	return c.CreatePos
}

func (c *CreateMaterializedView) End() Pos {
	return c.StatementEnd
}

func (c *CreateMaterializedView) Type() error {
	return "MATERIALIZED_VIEW"
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

type CreateView struct {
	CreatePos    Pos // position of CREATE|ATTACH keyword
	StatementEnd Pos
	Name         *TableIdentifier
	IfNotExists  bool
	UUID         *UUID
	OnCluster    *ClusterClause
	TableSchema  *TableSchemaClause
	SubQuery     *SubQuery
}

func (c *CreateView) Pos() Pos {
	return c.CreatePos
}

func (c *CreateView) End() Pos {
	return c.StatementEnd
}

func (c *CreateView) Type() error {
	return "VIEW"
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

type CreateFunction struct {
	CreatePos    Pos
	IfNotExists  bool
	FunctionName *Ident
	OnCluster    *ClusterClause
	Params       *ParamExprList
	Expr         Expr
}

func (c *CreateFunction) Type() error {
	return "FUNCTION"
}

func (c *CreateFunction) Pos() Pos {
	return c.CreatePos
}

func (c *CreateFunction) End() Pos {
	return c.Expr.End()
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

type RoleName struct {
	Name      Expr
	Scope     *StringLiteral
	OnCluster *ClusterClause
}

func (r *RoleName) Pos() Pos {
	return r.Name.Pos()
}

func (r *RoleName) End() Pos {
	if r.Scope != nil {
		return r.Scope.End()
	}
	if r.OnCluster != nil {
		return r.OnCluster.End()
	}
	return r.Name.End()
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

type SettingPair struct {
	Name      *Ident
	Operation TokenKind
	Value     Expr
}

func (s *SettingPair) Pos() Pos {
	return s.Name.NamePos
}

func (s *SettingPair) End() Pos {
	return s.Value.End()
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

type RoleSetting struct {
	SettingPairs []*SettingPair
	Modifier     *Ident
}

func (r *RoleSetting) Pos() Pos {
	if len(r.SettingPairs) > 0 {
		return r.SettingPairs[0].Pos()
	}
	return r.Modifier.NamePos
}

func (r *RoleSetting) End() Pos {
	if r.Modifier != nil {
		return r.Modifier.NameEnd
	}
	return r.SettingPairs[len(r.SettingPairs)-1].End()
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

type CreateRole struct {
	CreatePos         Pos
	StatementEnd      Pos
	IfNotExists       bool
	OrReplace         bool
	RoleNames         []*RoleName
	AccessStorageType *Ident
	Settings          []*RoleSetting
}

func (c *CreateRole) Pos() Pos {
	return c.CreatePos
}

func (c *CreateRole) End() Pos {
	return c.StatementEnd
}

func (c *CreateRole) Type() error {
	return "ROLE"
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

type AlterRole struct {
	AlterPos        Pos
	StatementEnd    Pos
	IfExists        bool
	RoleRenamePairs []*RoleRenamePair
	Settings        []*RoleSetting
}

func (a *AlterRole) Pos() Pos {
	return a.AlterPos
}

func (a *AlterRole) End() Pos {
	return a.StatementEnd
}

func (a *AlterRole) Type() error {
	return "ROLE"
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

type RoleRenamePair struct {
	RoleName     *RoleName
	NewName      Expr
	StatementEnd Pos
}

func (r *RoleRenamePair) Pos() Pos {
	return r.RoleName.Pos()
}

func (r *RoleRenamePair) End() Pos {
	return r.StatementEnd
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

type DestinationClause struct {
	ToPos           Pos
	TableIdentifier *TableIdentifier
	TableSchema     *TableSchemaClause
}

func (d *DestinationClause) Pos() Pos {
	return d.ToPos
}

func (d *DestinationClause) End() Pos {
	return d.TableIdentifier.End()
}

func (p *PrintVisitor) VisitDestinationClause(d *DestinationClause) error {
	builder := p.builder
	builder.WriteString("TO ")
	builder.WriteString(d.TableIdentifier.String())
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

type ConstraintClause struct {
	ConstraintPos Pos
	Constraint    *Ident
	Expr          Expr
}

func (c *ConstraintClause) Pos() Pos {
	return c.ConstraintPos
}

func (c *ConstraintClause) End() Pos {
	return c.Expr.End()
}

func (p *PrintVisitor) VisitConstraintClause(c *ConstraintClause) error {
	builder := p.builder
	builder.WriteString(c.Constraint.String())
	builder.WriteByte(' ')
	builder.WriteString(c.Expr.String())
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

type NullLiteral struct {
	NullPos Pos
}

func (n *NullLiteral) Pos() Pos {
	return n.NullPos
}

func (n *NullLiteral) End() Pos {
	return n.NullPos + 4
}

func (p *PrintVisitor) VisitNullLiteral(n *NullLiteral) error {
	return "NULL"
}

func (visitor MyDefaultVisitor) VisitNullLiteral(n *NullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}

type NotNullLiteral struct {
	NotPos      Pos
	NullLiteral *NullLiteral
}

func (n *NotNullLiteral) Pos() Pos {
	return n.NotPos
}

func (n *NotNullLiteral) End() Pos {
	return n.NullLiteral.End()
}

func (p *PrintVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	return "NOT NULL"
}

func (visitor MyDefaultVisitor) VisitNotNullLiteral(n *NotNullLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.NullLiteral.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type NestedIdentifier struct {
	Ident    *Ident
	DotIdent *Ident
}

func (n *NestedIdentifier) Pos() Pos {
	return n.Ident.Pos()
}

func (n *NestedIdentifier) End() Pos {
	if n.DotIdent != nil {
		return n.DotIdent.End()
	}
	return n.Ident.End()
}

func (p *PrintVisitor) VisitNestedIdentifier(n *NestedIdentifier) error {
	if n.DotIdent != nil {
		return n.Ident.String() + "." + n.DotIdent.String()
	}
	return n.Ident.String()
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

type ColumnIdentifier struct {
	Database *Ident
	Table    *Ident
	Column   *Ident
}

func (c *ColumnIdentifier) Pos() Pos {
	if c.Database != nil {
		return c.Database.NamePos
	} else if c.Table != nil {
		return c.Table.NamePos
	} else {
		return c.Column.NamePos
	}
}

func (c *ColumnIdentifier) End() Pos {
	return c.Column.NameEnd
}

func (p *PrintVisitor) VisitColumnIdentifier(c *ColumnIdentifier) error {
	if c.Database != nil {
		return c.Database.String() + "." + c.Table.String() + "." + c.Column.String()
	} else if c.Table != nil {
		return c.Table.String() + "." + c.Column.String()
	} else {
		return c.Column.String()
	}
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

type TableIdentifier struct {
	Database *Ident
	Table    *Ident
}

func (t *TableIdentifier) Pos() Pos {
	if t.Database != nil {
		return t.Database.Pos()
	}
	return t.Table.Pos()
}

func (t *TableIdentifier) End() Pos {
	return t.Table.End()
}

func (p *PrintVisitor) VisitTableIdentifier(t *TableIdentifier) error {
	if t.Database != nil {
		return t.Database.String() + "." + t.Table.String()
	}
	return t.Table.String()
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

type TableSchemaClause struct {
	SchemaPos     Pos
	SchemaEnd     Pos
	Columns       []Expr
	AliasTable    *TableIdentifier
	TableFunction *TableFunctionExpr
}

func (t *TableSchemaClause) Pos() Pos {
	return t.SchemaPos
}

func (t *TableSchemaClause) End() Pos {
	return t.SchemaEnd
}

func (p *PrintVisitor) VisitTableSchemaClause(t *TableSchemaClause) error {
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

type TableArgListExpr struct {
	LeftParenPos  Pos
	RightParenPos Pos
	Args          []Expr
}

func (t *TableArgListExpr) Pos() Pos {
	return t.LeftParenPos
}

func (t *TableArgListExpr) End() Pos {
	return t.RightParenPos
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

type TableFunctionExpr struct {
	Name Expr
	Args *TableArgListExpr
}

func (t *TableFunctionExpr) Pos() Pos {
	return t.Name.Pos()
}

func (t *TableFunctionExpr) End() Pos {
	return t.Args.End()
}

func (p *PrintVisitor) VisitTableFunctionExpr(t *TableFunctionExpr) error {
	builder := p.builder
	builder.WriteString(t.Name.String())
	builder.WriteString(t.Args.String())
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

type ClusterClause struct {
	OnPos Pos
	Expr  Expr
}

func (o *ClusterClause) Pos() Pos {
	return o.OnPos
}

func (o *ClusterClause) End() Pos {
	return o.Expr.End()
}

func (p *PrintVisitor) VisitClusterClause(o *ClusterClause) error {
	builder := p.builder
	builder.WriteString("ON CLUSTER ")
	builder.WriteString(o.Expr.String())
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

type PartitionClause struct {
	PartitionPos Pos
	Expr         Expr
	ID           *StringLiteral
	All          bool
}

func (p *PartitionClause) Pos() Pos {
	return p.PartitionPos
}

func (p *PartitionClause) End() Pos {
	if p.ID != nil {
		return p.ID.LiteralEnd
	}
	return p.Expr.End()
}

func (p *PrintVisitor) VisitPartitionClause(p *PartitionClause) error {
	builder := p.builder
	builder.WriteString("PARTITION ")
	if p.ID != nil {
		builder.WriteString(p.ID.String())
	} else if p.All {
		builder.WriteString("ALL")
	} else {
		builder.WriteString(p.Expr.String())
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

type PartitionByClause struct {
	PartitionPos Pos
	Expr         Expr
}

func (p *PartitionByClause) Pos() Pos {
	return p.PartitionPos
}

func (p *PartitionByClause) End() Pos {
	return p.Expr.End()
}

func (p *PrintVisitor) VisitPartitionByClause(p *PartitionByClause) error {
	builder := p.builder
	builder.WriteString("PARTITION BY ")
	builder.WriteString(p.Expr.String())
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

type PrimaryKeyClause struct {
	PrimaryPos Pos
	Expr       Expr
}

func (p *PrimaryKeyClause) Pos() Pos {
	return p.PrimaryPos
}

func (p *PrimaryKeyClause) End() Pos {
	return p.Expr.End()
}

func (p *PrintVisitor) VisitPrimaryKeyClause(p *PrimaryKeyClause) error {
	builder := p.builder
	builder.WriteString("PRIMARY KEY ")
	builder.WriteString(p.Expr.String())
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

type SampleByClause struct {
	SamplePos Pos
	Expr      Expr
}

func (s *SampleByClause) Pos() Pos {
	return s.SamplePos
}

func (s *SampleByClause) End() Pos {
	return s.Expr.End()
}

func (p *PrintVisitor) VisitSampleByClause(s *SampleByClause) error {
	builder := p.builder
	builder.WriteString("SAMPLE BY ")
	builder.WriteString(s.Expr.String())
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

type TTLExpr struct {
	TTLPos Pos
	Expr   Expr
}

func (t *TTLExpr) Pos() Pos {
	return t.TTLPos
}

func (t *TTLExpr) End() Pos {
	return t.Expr.End()
}

func (p *PrintVisitor) VisitTTLExpr(t *TTLExpr) error {
	builder := p.builder
	builder.WriteString(t.Expr.String())
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

type TTLClause struct {
	TTLPos  Pos
	ListEnd Pos
	Items   []*TTLExpr
}

func (t *TTLClause) Pos() Pos {
	return t.TTLPos
}

func (t *TTLClause) End() Pos {
	return t.ListEnd
}

func (p *PrintVisitor) VisitTTLClause(t *TTLClause) error {
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

type OrderExpr struct {
	OrderPos  Pos
	Expr      Expr
	Alias     *Ident
	Direction OrderDirection
}

func (o *OrderExpr) Pos() Pos {
	return o.OrderPos
}

func (o *OrderExpr) End() Pos {
	if o.Alias != nil {
		return o.Alias.End()
	}
	return o.Expr.End()
}

func (p *PrintVisitor) VisitOrderExpr(o *OrderExpr) error {
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

type OrderByClause struct {
	OrderPos Pos
	ListEnd  Pos
	Items    []Expr
}

func (o *OrderByClause) Pos() Pos {
	return o.OrderPos
}

func (o *OrderByClause) End() Pos {
	return o.ListEnd
}

func (p *PrintVisitor) VisitOrderByClause(o *OrderByClause) error {
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

type SettingExprList struct {
	SettingsPos Pos
	Name        *Ident
	Expr        Expr
}

func (s *SettingExprList) Pos() Pos {
	return s.SettingsPos
}

func (s *SettingExprList) End() Pos {
	return s.Expr.End()
}

func (p *PrintVisitor) VisitSettingExprList(s *SettingExprList) error {
	builder := p.builder
	builder.WriteString(s.Name.String())
	builder.WriteByte('=')
	builder.WriteString(s.Expr.String())
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

type SettingsClause struct {
	SettingsPos Pos
	ListEnd     Pos
	Items       []*SettingExprList
}

func (s *SettingsClause) Pos() Pos {
	return s.SettingsPos
}

func (s *SettingsClause) End() Pos {
	return s.ListEnd
}

func (p *PrintVisitor) VisitSettingsClause(s *SettingsClause) error {
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

type ParamExprList struct {
	LeftParenPos  Pos
	RightParenPos Pos
	Items         *ColumnExprList
	ColumnArgList *ColumnArgList
}

func (f *ParamExprList) Pos() Pos {
	return f.LeftParenPos
}

func (f *ParamExprList) End() Pos {
	return f.RightParenPos
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

type KeyValue struct {
	Key   StringLiteral
	Value Expr
}

type MapLiteral struct {
	LBracePos Pos
	RBracePos Pos
	KeyValues []KeyValue
}

func (m *MapLiteral) Pos() Pos {
	return m.LBracePos
}

func (m *MapLiteral) End() Pos {
	return m.RBracePos
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

type QueryParam struct {
	LBracePos Pos
	RBracePos Pos
	Name      *Ident
	Type      ColumnType
}

func (q *QueryParam) Pos() Pos {
	return q.LBracePos
}

func (q *QueryParam) End() Pos {
	return q.RBracePos
}

func (p *PrintVisitor) VisitQueryParam(q *QueryParam) error {
	builder := p.builder
	builder.WriteString("{")
	builder.WriteString(q.Name.String())
	builder.WriteString(": ")
	builder.WriteString(q.Type.String())
	builder.WriteString("}")
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

type ArrayParamList struct {
	LeftBracketPos  Pos
	RightBracketPos Pos
	Items           *ColumnExprList
}

func (a *ArrayParamList) Pos() Pos {
	return a.LeftBracketPos
}

func (a *ArrayParamList) End() Pos {
	return a.RightBracketPos
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

func (visitor MyDefaultVisitor) VisitArrayParamList(a *ArrayParamList) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Items.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type ObjectParams struct {
	Object Expr
	Params *ArrayParamList
}

func (o *ObjectParams) Pos() Pos {
	return o.Object.Pos()
}

func (o *ObjectParams) End() Pos {
	return o.Params.End()
}

func (p *PrintVisitor) VisitObjectParams(o *ObjectParams) error {
	builder := p.builder
	builder.WriteString(o.Object.String())
	builder.WriteString(o.Params.String())
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

type FunctionExpr struct {
	Name   *Ident
	Params *ParamExprList
}

func (f *FunctionExpr) Pos() Pos {
	return f.Name.NamePos
}

func (f *FunctionExpr) End() Pos {
	return f.Params.RightParenPos
}

func (p *PrintVisitor) VisitFunctionExpr(f *FunctionExpr) error {
	builder := p.builder
	builder.WriteString(f.Name.String())
	builder.WriteString(f.Params.String())
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

type WindowFunctionExpr struct {
	Function *FunctionExpr
	OverPos  Pos
	OverExpr Expr
}

func (w *WindowFunctionExpr) Pos() Pos {
	return w.Function.Pos()
}

func (w *WindowFunctionExpr) End() Pos {
	return w.OverExpr.End()
}

func (p *PrintVisitor) VisitWindowFunctionExpr(w *WindowFunctionExpr) error {
	builder := p.builder
	builder.WriteString(w.Function.String())
	builder.WriteString(" OVER ")
	builder.WriteString(w.OverExpr.String())
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

type ColumnExpr struct {
	Expr  Expr
	Alias *Ident
}

func (c *ColumnExpr) Pos() Pos {
	return c.Expr.Pos()
}

func (c *ColumnExpr) End() Pos {
	if c.Alias != nil {
		return c.Alias.NameEnd
	}
	return c.Expr.End()
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

type ColumnDef struct {
	NamePos   Pos
	ColumnEnd Pos
	Name      *NestedIdentifier
	Type      ColumnType
	NotNull   *NotNullLiteral
	Nullable  *NullLiteral

	DefaultExpr      Expr
	MaterializedExpr Expr
	AliasExpr        Expr

	Codec *CompressionCodec
	TTL   Expr

	Comment          *StringLiteral
	CompressionCodec *Ident
}

func (c *ColumnDef) Pos() Pos {
	return c.Name.Pos()
}

func (c *ColumnDef) End() Pos {
	return c.ColumnEnd
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

type ColumnType interface {
	Expr
	Type() string
}

type ScalarType struct {
	Name *Ident
}

func (s *ScalarType) Pos() Pos {
	return s.Name.NamePos
}

func (s *ScalarType) End() Pos {
	return s.Name.NameEnd
}

func (p *PrintVisitor) VisitScalarType(s *ScalarType) error {
	return s.Name.String()
}

func (visitor MyDefaultVisitor) VisitScalarType(s *ScalarType) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (s *ScalarType) Type() error {
	return s.Name.Name
}

type JSONPath struct {
	Idents []*Ident
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

type JSONOption struct {
	SkipPath  *JSONPath
	SkipRegex *StringLiteral
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

type JSONOptions struct {
	LParen Pos
	RParen Pos
	Items  []*JSONOption
}

func (j *JSONOptions) Pos() Pos {
	return j.LParen
}

func (j *JSONOptions) End() Pos {
	return j.RParen
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

type JSONType struct {
	Name    *Ident
	Options *JSONOptions
}

func (j *JSONType) Pos() Pos {
	return j.Name.NamePos
}

func (j *JSONType) End() Pos {
	if j.Options != nil {
		return j.Options.RParen
	}
	return j.Name.NameEnd
}

func (p *PrintVisitor) VisitJSONType(j *JSONType) error {
	builder := p.builder
	builder.WriteString(j.Name.String())
	if j.Options != nil {
		builder.WriteString(j.Options.String())
	}
	return nil
}

func (j *JSONType) Type() error {
	return j.Name.Name
}

func (visitor MyDefaultVisitor) VisitJSONType(j *JSONType) error {
	visitor.enter(j)
	defer visitor.leave(j)
	if err := j.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type PropertyType struct {
	Name *Ident
}

func (c *PropertyType) Pos() Pos {
	return c.Name.NamePos
}

func (c *PropertyType) End() Pos {
	return c.Name.NameEnd
}

func (p *PrintVisitor) VisitPropertyType(c *PropertyType) error {
	return c.Name.String()
}

func (visitor MyDefaultVisitor) VisitPropertyType(c *PropertyType) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}

func (c *PropertyType) Type() error {
	return c.Name.Name
}

type TypeWithParams struct {
	LeftParenPos  Pos
	RightParenPos Pos
	Name          *Ident
	Params        []Literal
}

func (s *TypeWithParams) Pos() Pos {
	return s.Name.NamePos
}

func (s *TypeWithParams) End() Pos {
	return s.RightParenPos
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

func (s *TypeWithParams) Type() error {
	return s.Name.Name
}

type ComplexType struct {
	LeftParenPos  Pos
	RightParenPos Pos
	Name          *Ident
	Params        []ColumnType
}

func (c *ComplexType) Pos() Pos {
	return c.Name.NamePos
}

func (c *ComplexType) End() Pos {
	return c.RightParenPos
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

func (c *ComplexType) Type() error {
	return c.Name.Name
}

type NestedType struct {
	LeftParenPos  Pos
	RightParenPos Pos
	Name          *Ident
	Columns       []Expr
}

func (n *NestedType) Pos() Pos {
	return n.Name.NamePos
}

func (n *NestedType) End() Pos {
	return n.RightParenPos
}

func (p *PrintVisitor) VisitNestedType(n *NestedType) error {
	builder := p.builder
	// on the same level as the column type
	builder.WriteString(n.Name.String())
	builder.WriteByte('(')
	for i, column := range n.Columns {
		builder.WriteString(column.String())
		if i != len(n.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	// right paren needs to be on the same level as the column
	builder.WriteByte(')')
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

func (n *NestedType) Type() error {
	return n.Name.Name
}

type CompressionCodec struct {
	CodecPos      Pos
	RightParenPos Pos
	Type          *Ident
	TypeLevel     *NumberLiteral
	Name          *Ident
	Level         *NumberLiteral // compression level
}

func (c *CompressionCodec) Pos() Pos {
	return c.CodecPos
}

func (c *CompressionCodec) End() Pos {
	return c.RightParenPos
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

type Literal interface {
	Expr
}

type NumberLiteral struct {
	NumPos  Pos
	NumEnd  Pos
	Literal string
	Base    int
}

func (n *NumberLiteral) Pos() Pos {
	return n.NumPos
}

func (n *NumberLiteral) End() Pos {
	return n.NumEnd
}

func (p *PrintVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	return n.Literal
}

func (visitor MyDefaultVisitor) VisitNumberLiteral(n *NumberLiteral) error {
	visitor.enter(n)
	defer visitor.leave(n)
	return nil
}

type StringLiteral struct {
	LiteralPos Pos
	LiteralEnd Pos
	Literal    string
}

func (s *StringLiteral) Pos() Pos {
	return s.LiteralPos
}

func (s *StringLiteral) End() Pos {
	return s.LiteralEnd
}

func (p *PrintVisitor) VisitStringLiteral(s *StringLiteral) error {
	return "'" + s.Literal + "'"
}

func (visitor MyDefaultVisitor) VisitStringLiteral(s *StringLiteral) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}

type PlaceHolder struct {
	PlaceholderPos Pos
	PlaceHolderEnd Pos
	Type           string
}

func (p *PlaceHolder) Pos() Pos {
	return p.PlaceholderPos
}

func (p *PlaceHolder) End() Pos {
	return p.PlaceHolderEnd
}

func (p *PrintVisitor) VisitPlaceHolder(p *PlaceHolder) error {
	return p.Type
}

func (visitor MyDefaultVisitor) VisitPlaceHolder(p *PlaceHolder) error {
	visitor.enter(p)
	defer visitor.leave(p)
	return nil
}

type RatioExpr struct {
	Numerator *NumberLiteral
	// numberLiteral (SLASH numberLiteral)?
	Denominator *NumberLiteral
}

func (r *RatioExpr) Pos() Pos {
	return r.Numerator.NumPos
}

func (r *RatioExpr) End() Pos {
	if r.Denominator != nil {
		return r.Denominator.NumEnd
	}
	return r.Numerator.NumEnd
}

func (p *PrintVisitor) VisitRatioExpr(r *RatioExpr) error {
	builder := p.builder
	builder.WriteString(r.Numerator.String())
	if r.Denominator != nil {
		builder.WriteString("/")
		builder.WriteString(r.Denominator.String())
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

type EnumValue struct {
	Name  *StringLiteral
	Value *NumberLiteral
}

func (e *EnumValue) Pos() Pos {
	return e.Name.Pos()
}

func (e *EnumValue) End() Pos {
	return e.Value.End()
}

func (p *PrintVisitor) VisitEnumValue(e *EnumValue) error {
	builder := p.builder
	builder.WriteString(e.Name.String())
	builder.WriteByte('=')
	builder.WriteString(e.Value.String())
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

type EnumType struct {
	Name    *Ident
	ListPos Pos
	ListEnd Pos
	Values  []EnumValue
}

func (e *EnumType) Pos() Pos {
	return e.ListPos
}

func (e *EnumType) End() Pos {
	return e.ListEnd
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

func (e *EnumType) Type() error {
	return e.Name.Name
}

type IntervalExpr struct {
	IntervalPos Pos
	Expr        Expr
	Unit        *Ident
}

func (i *IntervalExpr) Pos() Pos {
	return i.IntervalPos
}

func (i *IntervalExpr) End() Pos {
	return i.Unit.End()
}

func (p *PrintVisitor) VisitIntervalExpr(i *IntervalExpr) error {
	builder := p.builder
	builder.WriteString("INTERVAL ")
	builder.WriteString(i.Expr.String())
	builder.WriteByte(' ')
	builder.WriteString(i.Unit.String())
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

// TODO(@git-hulk): split into EngineClause and EngineExpr
type EngineExpr struct {
	EnginePos   Pos
	EngineEnd   Pos
	Name        string
	Params      *ParamExprList
	PrimaryKey  *PrimaryKeyClause
	PartitionBy *PartitionByClause
	SampleBy    *SampleByClause
	TTL         *TTLClause
	Settings    *SettingsClause
	OrderBy     *OrderByClause
}

func (e *EngineExpr) Pos() Pos {
	return e.EnginePos
}

func (e *EngineExpr) End() Pos {
	return e.EngineEnd
}

func (p *PrintVisitor) VisitEngineExpr(e *EngineExpr) error {
	// align with the engine level
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

type ColumnTypeExpr struct {
	Name *Ident
}

func (c *ColumnTypeExpr) Pos() Pos {
	return c.Name.NamePos
}

func (c *ColumnTypeExpr) End() Pos {
	return c.Name.NameEnd
}

func (p *PrintVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	return c.Name.String()
}

func (visitor MyDefaultVisitor) VisitColumnTypeExpr(c *ColumnTypeExpr) error {
	visitor.enter(c)
	defer visitor.leave(c)
	if err := c.Name.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type ColumnArgList struct {
	Distinct      bool
	LeftParenPos  Pos
	RightParenPos Pos
	Items         []Expr
}

func (c *ColumnArgList) Pos() Pos {
	return c.LeftParenPos
}

func (c *ColumnArgList) End() Pos {
	return c.RightParenPos
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

type ColumnExprList struct {
	ListPos     Pos
	ListEnd     Pos
	HasDistinct bool
	Items       []Expr
}

func (c *ColumnExprList) Pos() Pos {
	return c.ListPos
}

func (c *ColumnExprList) End() Pos {
	return c.ListEnd
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

type WhenClause struct {
	WhenPos Pos
	ThenPos Pos
	When    Expr
	Then    Expr
	ElsePos Pos
	Else    Expr
}

func (w *WhenClause) Pos() Pos {
	return w.WhenPos
}

func (w *WhenClause) End() Pos {
	if w.Else != nil {
		return w.Else.End()
	}
	return w.Then.End()
}

func (p *PrintVisitor) VisitWhenClause(w *WhenClause) error {
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

type CaseExpr struct {
	CasePos Pos
	EndPos  Pos
	Expr    Expr // optional
	Whens   []*WhenClause
	ElsePos Pos
	Else    Expr
}

func (c *CaseExpr) Pos() Pos {
	return c.CasePos
}

func (c *CaseExpr) End() Pos {
	return c.EndPos
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

type CastExpr struct {
	CastPos   Pos
	Expr      Expr
	Separator string
	AsPos     Pos
	AsType    Expr
}

func (c *CastExpr) Pos() Pos {
	return c.CastPos
}

func (c *CastExpr) End() Pos {
	return c.AsType.End()
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

type WithClause struct {
	WithPos Pos
	EndPos  Pos
	CTEs    []*CTEStmt
}

func (w *WithClause) Pos() Pos {
	return w.WithPos
}

func (w *WithClause) End() Pos {
	return w.EndPos
}

func (p *PrintVisitor) VisitWithClause(w *WithClause) error {
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

type TopClause struct {
	TopPos   Pos
	TopEnd   Pos
	Number   *NumberLiteral
	WithTies bool
}

func (t *TopClause) Pos() Pos {
	return t.TopPos
}

func (t *TopClause) End() Pos {
	return t.TopEnd
}

func (p *PrintVisitor) VisitTopClause(t *TopClause) error {
	builder := p.builder
	builder.WriteString("TOP ")
	builder.WriteString(t.Number.Literal)
	if t.WithTies {
		return "WITH TIES"
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

type CreateLiveView struct {
	CreatePos    Pos
	StatementEnd Pos
	Name         *TableIdentifier
	IfNotExists  bool
	UUID         *UUID
	OnCluster    *ClusterClause
	Destination  *DestinationClause
	TableSchema  *TableSchemaClause
	WithTimeout  *WithTimeoutClause
	SubQuery     *SubQuery
}

func (c *CreateLiveView) Type() error {
	return "LIVE_VIEW"
}

func (c *CreateLiveView) Pos() Pos {
	return c.CreatePos
}

func (c *CreateLiveView) End() Pos {
	return c.StatementEnd
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

type WithTimeoutClause struct {
	WithTimeoutPos Pos
	Expr           Expr
	Number         *NumberLiteral
}

func (w *WithTimeoutClause) Pos() Pos {
	return w.WithTimeoutPos
}

func (w *WithTimeoutClause) End() Pos {
	return w.Number.End()
}

func (p *PrintVisitor) VisitWithTimeoutClause(w *WithTimeoutClause) error {
	builder := p.builder
	builder.WriteString("WITH TIMEOUT ")
	builder.WriteString(w.Number.String())
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

type TableExpr struct {
	TablePos Pos
	TableEnd Pos
	Alias    *AliasExpr
	Expr     Expr
	HasFinal bool
}

func (t *TableExpr) Pos() Pos {
	return t.TablePos
}

func (t *TableExpr) End() Pos {
	return t.TableEnd
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

type OnClause struct {
	OnPos Pos
	On    *ColumnExprList
}

func (o *OnClause) Pos() Pos {
	return o.OnPos
}

func (o *OnClause) End() Pos {
	return o.On.End()
}

func (p *PrintVisitor) VisitOnClause(o *OnClause) error {
	builder := p.builder
	builder.WriteString("ON ")
	builder.WriteString(o.On.String())
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

type UsingClause struct {
	UsingPos Pos
	Using    *ColumnExprList
}

func (u *UsingClause) Pos() Pos {
	return u.UsingPos
}

func (u *UsingClause) End() Pos {
	return u.Using.End()
}

func (p *PrintVisitor) VisitUsingClause(u *UsingClause) error {
	builder := p.builder
	builder.WriteString("USING ")
	builder.WriteString(u.Using.String())
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

type JoinExpr struct {
	JoinPos     Pos
	Left        Expr
	Right       Expr
	Modifiers   []string
	Constraints Expr
}

func (j *JoinExpr) Pos() Pos {
	return j.JoinPos
}

func (j *JoinExpr) End() Pos {
	return j.Left.End()
}

func buildJoinString(builder *strings.Builder, expr Expr) {
	joinExpr, ok := expr.(*JoinExpr)
	if !ok {
		builder.WriteString(",")
		builder.WriteString(expr.String())
		return
	}

	if len(joinExpr.Modifiers) == 0 {
		builder.WriteString(",")
	} else {
		builder.WriteString(" ")
		builder.WriteString(strings.Join(joinExpr.Modifiers, " "))
		builder.WriteByte(' ')
	}
	builder.WriteString(joinExpr.Left.String())
	if joinExpr.Constraints != nil {
		builder.WriteByte(' ')
		builder.WriteString(joinExpr.Constraints.String())
	}
	if joinExpr.Right != nil {
		buildJoinString(builder, joinExpr.Right)
	}
}

func (p *PrintVisitor) VisitJoinExpr(j *JoinExpr) error {
	builder := p.builder
	builder.WriteString(j.Left.String())
	if j.Right != nil {
		buildJoinString(&builder, j.Right)
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

type JoinConstraintClause struct {
	ConstraintPos Pos
	On            *ColumnExprList
	Using         *ColumnExprList
}

func (j *JoinConstraintClause) Pos() Pos {
	return j.ConstraintPos
}

func (j *JoinConstraintClause) End() Pos {
	if j.On != nil {
		return j.On.End()
	}
	return j.Using.End()
}

func (p *PrintVisitor) VisitJoinConstraintClause(j *JoinConstraintClause) error {
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

type FromClause struct {
	FromPos Pos
	Expr    Expr
}

func (f *FromClause) Pos() Pos {
	return f.FromPos
}

func (f *FromClause) End() Pos {
	return f.Expr.End()
}

func (p *PrintVisitor) VisitFromClause(f *FromClause) error {
	builder := p.builder
	builder.WriteString("FROM ")
	builder.WriteString(f.Expr.String())
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

type IsNullExpr struct {
	IsPos Pos
	Expr  Expr
}

func (n *IsNullExpr) Pos() Pos {
	return n.IsPos
}

func (n *IsNullExpr) End() Pos {
	return n.Expr.End()
}

func (p *PrintVisitor) VisitIsNullExpr(n *IsNullExpr) error {
	builder := p.builder
	builder.WriteString(n.Expr.String())
	builder.WriteString(" IS NULL")
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

type IsNotNullExpr struct {
	IsPos Pos
	Expr  Expr
}

func (n *IsNotNullExpr) Pos() Pos {
	return n.Expr.Pos()
}

func (n *IsNotNullExpr) End() Pos {
	return n.Expr.End()
}

func (p *PrintVisitor) VisitIsNotNullExpr(n *IsNotNullExpr) error {
	builder := p.builder
	builder.WriteString(n.Expr.String())
	builder.WriteString(" IS NOT NULL")
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

type AliasExpr struct {
	Expr     Expr
	AliasPos Pos
	Alias    Expr
}

func (a *AliasExpr) Pos() Pos {
	return a.AliasPos
}

func (a *AliasExpr) End() Pos {
	return a.Alias.End()
}

func (p *PrintVisitor) VisitAliasExpr(a *AliasExpr) error {
	builder := p.builder
	if _, isSelect := a.Expr.(*SelectQuery); isSelect {
		builder.WriteByte('(')
		builder.WriteString(a.Expr.String())
		builder.WriteByte(')')
	} else {
		builder.WriteString(a.Expr.String())
	}
	builder.WriteString(" AS ")
	builder.WriteString(a.Alias.String())
	return nil
}

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

type WhereClause struct {
	WherePos Pos
	Expr     Expr
}

func (w *WhereClause) Pos() Pos {
	return w.WherePos
}

func (w *WhereClause) End() Pos {
	return w.Expr.End()
}

func (p *PrintVisitor) VisitWhereClause(w *WhereClause) error {
	builder := p.builder
	builder.WriteString("WHERE ")
	builder.WriteString(w.Expr.String())
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

type PrewhereClause struct {
	PrewherePos Pos
	Expr        Expr
}

func (w *PrewhereClause) Pos() Pos {
	return w.PrewherePos
}

func (w *PrewhereClause) End() Pos {
	return w.Expr.End()
}

func (p *PrintVisitor) VisitPrewhereClause(w *PrewhereClause) error {
	return "PREWHERE " + w.Expr.String()
}

func (visitor MyDefaultVisitor) VisitPrewhereClause(w *PrewhereClause) error {
	visitor.enter(w)
	defer visitor.leave(w)
	if err := w.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type GroupByClause struct {
	GroupByPos    Pos
	AggregateType string
	Expr          Expr
	WithCube      bool
	WithRollup    bool
	WithTotals    bool
}

func (g *GroupByClause) Pos() Pos {
	return g.GroupByPos
}

func (g *GroupByClause) End() Pos {
	return g.Expr.End()
}

func (p *PrintVisitor) VisitGroupByClause(g *GroupByClause) error {
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

func (visitor MyDefaultVisitor) VisitGroupByClause(g *GroupByClause) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type HavingClause struct {
	HavingPos Pos
	Expr      Expr
}

func (h *HavingClause) Pos() Pos {
	return h.HavingPos
}

func (h *HavingClause) End() Pos {
	return h.Expr.End()
}

func (p *PrintVisitor) VisitHavingClause(h *HavingClause) error {
	return "HAVING " + h.Expr.String()
}

func (visitor MyDefaultVisitor) VisitHavingClause(h *HavingClause) error {
	visitor.enter(h)
	defer visitor.leave(h)
	if err := h.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type LimitClause struct {
	LimitPos Pos
	Limit    Expr
	Offset   Expr
}

func (l *LimitClause) Pos() Pos {
	return l.LimitPos
}

func (l *LimitClause) End() Pos {
	if l.Offset != nil {
		return l.Offset.End()
	}
	return l.Limit.End()
}

func (p *PrintVisitor) VisitLimitClause(l *LimitClause) error {
	builder := p.builder
	builder.WriteString("LIMIT ")
	builder.WriteString(l.Limit.String())
	if l.Offset != nil {
		builder.WriteString(" OFFSET ")
		builder.WriteString(l.Offset.String())
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

type LimitByClause struct {
	Limit  *LimitClause
	ByExpr *ColumnExprList
}

func (l *LimitByClause) Pos() Pos {
	return l.Limit.Pos()
}

func (l *LimitByClause) End() Pos {
	if l.ByExpr != nil {
		return l.ByExpr.End()
	}
	if l.Limit != nil {
		return l.Limit.End()
	}
	return l.Limit.End()
}

func (p *PrintVisitor) VisitLimitByClause(l *LimitByClause) error {
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

type WindowExpr struct {
	LeftParenPos  Pos
	RightParenPos Pos
	PartitionBy   *PartitionByClause
	OrderBy       *OrderByClause
	Frame         *WindowFrameClause
}

func (w *WindowExpr) Pos() Pos {
	return w.LeftParenPos
}

func (w *WindowExpr) End() Pos {
	return w.RightParenPos
}

func (p *PrintVisitor) VisitWindowExpr(w *WindowExpr) error {
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

type WindowClause struct {
	*WindowExpr

	WindowPos Pos
	Name      *Ident
	AsPos     Pos
}

func (w *WindowClause) Pos() Pos {
	return w.WindowPos
}

func (w *WindowClause) End() Pos {
	return w.WindowExpr.End()
}

func (p *PrintVisitor) VisitWindowClause(w *WindowClause) error {
	builder := p.builder
	builder.WriteString("WINDOW ")
	builder.WriteString(w.Name.String())
	builder.WriteString(" ")
	builder.WriteString(w.WindowExpr.String())
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

type WindowFrameClause struct {
	FramePos Pos
	Type     string
	Extend   Expr
}

func (f *WindowFrameClause) Pos() Pos {
	return f.FramePos
}

func (f *WindowFrameClause) End() Pos {
	return f.Extend.End()
}

func (p *PrintVisitor) VisitWindowFrameClause(f *WindowFrameClause) error {
	builder := p.builder
	builder.WriteString(f.Type)
	builder.WriteString(" ")
	builder.WriteString(f.Extend.String())
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

type WindowFrameExtendExpr struct {
	Expr Expr
}

func (f *WindowFrameExtendExpr) Pos() Pos {
	return f.Expr.Pos()
}

func (f *WindowFrameExtendExpr) End() Pos {
	return f.Expr.End()
}

func (p *PrintVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	return f.Expr.String()
}

func (visitor MyDefaultVisitor) VisitWindowFrameExtendExpr(f *WindowFrameExtendExpr) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type BetweenClause struct {
	Expr    Expr
	Between Expr
	AndPos  Pos
	And     Expr
}

func (f *BetweenClause) Pos() Pos {
	return f.Expr.Pos()
}

func (f *BetweenClause) End() Pos {
	return f.And.End()
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

type WindowFrameCurrentRow struct {
	CurrentPos Pos
	RowEnd     Pos
}

func (f *WindowFrameCurrentRow) Pos() Pos {
	return f.CurrentPos
}

func (f *WindowFrameCurrentRow) End() Pos {
	return f.RowEnd
}

func (p *PrintVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	return "CURRENT ROW"
}

func (visitor MyDefaultVisitor) VisitWindowFrameCurrentRow(f *WindowFrameCurrentRow) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}

type WindowFrameUnbounded struct {
	UnboundedPos Pos
	UnboundedEnd Pos
	Direction    string
}

func (f *WindowFrameUnbounded) Pos() Pos {
	return f.UnboundedPos
}

func (f *WindowFrameUnbounded) End() Pos {
	return f.UnboundedEnd
}

func (p *PrintVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	return f.Direction + " UNBOUNDED"
}

func (visitor MyDefaultVisitor) VisitWindowFrameUnbounded(f *WindowFrameUnbounded) error {
	visitor.enter(f)
	defer visitor.leave(f)
	return nil
}

type WindowFrameNumber struct {
	Number       *NumberLiteral
	UnboundedEnd Pos
	Direction    string
}

func (f *WindowFrameNumber) Pos() Pos {
	return f.Number.Pos()
}

func (f *WindowFrameNumber) End() Pos {
	return f.UnboundedEnd
}

func (p *PrintVisitor) VisitWindowFrameNumber(f *WindowFrameNumber) error {
	builder := p.builder
	builder.WriteString(f.Number.String())
	builder.WriteByte(' ')
	builder.WriteString(f.Direction)
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

type ArrayJoinClause struct {
	ArrayPos Pos
	Type     string
	Expr     Expr
}

func (a *ArrayJoinClause) Pos() Pos {
	return a.ArrayPos
}

func (a *ArrayJoinClause) End() Pos {
	return a.Expr.End()
}

func (p *PrintVisitor) VisitArrayJoinClause(a *ArrayJoinClause) error {
	return a.Type + " ARRAY JOIN " + a.Expr.String()
}

func (visitor MyDefaultVisitor) VisitArrayJoinClause(a *ArrayJoinClause) error {
	visitor.enter(a)
	defer visitor.leave(a)
	if err := a.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type SelectQuery struct {
	SelectPos     Pos
	StatementEnd  Pos
	With          *WithClause
	Top           *TopClause
	SelectItems   []*SelectItem
	From          *FromClause
	ArrayJoin     *ArrayJoinClause
	Window        *WindowClause
	Prewhere      *PrewhereClause
	Where         *WhereClause
	GroupBy       *GroupByClause
	WithTotal     bool
	Having        *HavingClause
	OrderBy       *OrderByClause
	LimitBy       *LimitByClause
	Limit         *LimitClause
	Settings      *SettingsClause
	Format        *FormatClause
	UnionAll      *SelectQuery
	UnionDistinct *SelectQuery
	Except        *SelectQuery
}

func (s *SelectQuery) Pos() Pos {
	return s.SelectPos
}

func (s *SelectQuery) End() Pos {
	return s.StatementEnd
}

func (p *PrintVisitor) VisitSelectQuery(s *SelectQuery) error { // nolint: funlen
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

type SubQuery struct {
	HasParen bool
	Select   *SelectQuery
}

func (s *SubQuery) Pos() Pos {
	return s.Select.Pos()
}

func (s *SubQuery) End() Pos {
	return s.Select.End()
}

func (p *PrintVisitor) VisitSubQuery(s *SubQuery) error {
	if s.HasParen {
		builder := p.builder
		builder.WriteString("(")
		builder.WriteString(s.Select.String())
		builder.WriteString(")")
		return nil
	}
	return s.Select.String()
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

type NotExpr struct {
	NotPos Pos
	Expr   Expr
}

func (n *NotExpr) Pos() Pos {
	return n.NotPos
}

func (n *NotExpr) End() Pos {
	return n.Expr.End()
}

func (p *PrintVisitor) VisitNotExpr(n *NotExpr) error {
	return "NOT " + n.Expr.String()
}

func (visitor MyDefaultVisitor) VisitNotExpr(n *NotExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type NegateExpr struct {
	NegatePos Pos
	Expr      Expr
}

func (n *NegateExpr) Pos() Pos {
	return n.NegatePos
}

func (n *NegateExpr) End() Pos {
	return n.Expr.End()
}

func (p *PrintVisitor) VisitNegateExpr(n *NegateExpr) error {
	return "-" + n.Expr.String()
}

func (visitor MyDefaultVisitor) VisitNegateExpr(n *NegateExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type GlobalInOperation struct {
	GlobalPos Pos
	Expr      Expr
}

func (g *GlobalInOperation) Pos() Pos {
	return g.GlobalPos
}

func (g *GlobalInOperation) End() Pos {
	return g.Expr.End()
}

func (p *PrintVisitor) VisitGlobalInOperation(g *GlobalInOperation) error {
	return "GLOBAL " + g.Expr.String()
}

func (visitor MyDefaultVisitor) VisitGlobalInOperation(g *GlobalInOperation) error {
	visitor.enter(g)
	defer visitor.leave(g)
	if err := g.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type ExtractExpr struct {
	ExtractPos Pos
	Interval   *Ident
	FromPos    Pos
	FromExpr   Expr
}

func (e *ExtractExpr) Pos() Pos {
	return e.ExtractPos
}

func (e *ExtractExpr) End() Pos {
	return e.FromExpr.End()
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

func (visitor MyDefaultVisitor) VisitExtractExpr(e *ExtractExpr) error {
	visitor.enter(e)
	defer visitor.leave(e)
	if err := e.FromExpr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type DropDatabase struct {
	DropPos      Pos
	StatementEnd Pos
	Name         *Ident
	IfExists     bool
	OnCluster    *ClusterClause
}

func (d *DropDatabase) Pos() Pos {
	return d.DropPos
}

func (d *DropDatabase) End() Pos {
	return d.StatementEnd
}

func (d *DropDatabase) Type() error {
	return "DATABASE"
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

type DropStmt struct {
	DropPos      Pos
	StatementEnd Pos

	DropTarget  string
	Name        *TableIdentifier
	IfExists    bool
	OnCluster   *ClusterClause
	IsTemporary bool
	Modifier    string
}

func (d *DropStmt) Pos() Pos {
	return d.DropPos
}

func (d *DropStmt) End() Pos {
	return d.StatementEnd
}

func (d *DropStmt) Type() error {
	return "DROP " + d.DropTarget
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

type DropUserOrRole struct {
	DropPos      Pos
	Target       string
	StatementEnd Pos
	Names        []*RoleName
	IfExists     bool
	Modifier     string
	From         *Ident
}

func (d *DropUserOrRole) Pos() Pos {
	return d.DropPos
}

func (d *DropUserOrRole) End() Pos {
	return d.StatementEnd
}

func (d *DropUserOrRole) Type() error {
	return d.Target
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

type UseStmt struct {
	UsePos       Pos
	StatementEnd Pos
	Database     *Ident
}

func (u *UseStmt) Pos() Pos {
	return u.UsePos
}

func (u *UseStmt) End() Pos {
	return u.Database.End()
}

func (p *PrintVisitor) VisitUseStmt(u *UseStmt) error {
	return "USE " + u.Database.String()
}

func (visitor MyDefaultVisitor) VisitUseStmt(u *UseStmt) error {
	visitor.enter(u)
	defer visitor.leave(u)
	if err := u.Database.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type CTEStmt struct {
	CTEPos Pos
	Expr   Expr
	Alias  Expr
}

func (c *CTEStmt) Pos() Pos {
	return c.CTEPos
}

func (c *CTEStmt) End() Pos {
	return c.Expr.End()
}

func (p *PrintVisitor) VisitCTEStmt(c *CTEStmt) error {
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

type SetStmt struct {
	SetPos   Pos
	Settings *SettingsClause
}

func (s *SetStmt) Pos() Pos {
	return s.SetPos
}

func (s *SetStmt) End() Pos {
	return s.Settings.End()
}

func (p *PrintVisitor) VisitSetStmt(s *SetStmt) error {
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

func (visitor MyDefaultVisitor) VisitSetStmt(s *SetStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Settings.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type FormatClause struct {
	FormatPos Pos
	Format    *Ident
}

func (f *FormatClause) Pos() Pos {
	return f.FormatPos
}

func (f *FormatClause) End() Pos {
	return f.Format.End()
}

func (p *PrintVisitor) VisitFormatClause(f *FormatClause) error {
	return "FORMAT " + f.Format.String()
}

func (visitor MyDefaultVisitor) VisitFormatClause(f *FormatClause) error {
	visitor.enter(f)
	defer visitor.leave(f)
	if err := f.Format.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type OptimizeStmt struct {
	OptimizePos  Pos
	StatementEnd Pos
	Table        *TableIdentifier
	OnCluster    *ClusterClause
	Partition    *PartitionClause
	HasFinal     bool
	Deduplicate  *DeduplicateClause
}

func (o *OptimizeStmt) Pos() Pos {
	return o.OptimizePos
}

func (o *OptimizeStmt) End() Pos {
	return o.StatementEnd
}

func (p *PrintVisitor) VisitOptimizeStmt(o *OptimizeStmt) error {
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

type DeduplicateClause struct {
	DeduplicatePos Pos
	By             *ColumnExprList
	Except         *ColumnExprList
}

func (d *DeduplicateClause) Pos() Pos {
	return d.DeduplicatePos
}

func (d *DeduplicateClause) End() Pos {
	if d.By != nil {
		return d.By.End()
	} else if d.Except != nil {
		return d.Except.End()
	}
	return d.DeduplicatePos + Pos(len(KeywordDeduplicate))
}

func (p *PrintVisitor) VisitDeduplicateClause(d *DeduplicateClause) error {
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

type SystemStmt struct {
	SystemPos Pos
	Expr      Expr
}

func (s *SystemStmt) Pos() Pos {
	return s.SystemPos
}

func (s *SystemStmt) End() Pos {
	return s.Expr.End()
}

func (p *PrintVisitor) VisitSystemStmt(s *SystemStmt) error {
	return "SYSTEM " + s.Expr.String()
}

func (visitor MyDefaultVisitor) VisitSystemStmt(s *SystemStmt) error {
	visitor.enter(s)
	defer visitor.leave(s)
	if err := s.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type SystemFlushExpr struct {
	FlushPos     Pos
	StatementEnd Pos
	Logs         bool
	Distributed  *TableIdentifier
}

func (s *SystemFlushExpr) Pos() Pos {
	return s.FlushPos
}

func (s *SystemFlushExpr) End() Pos {
	return s.StatementEnd
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

type SystemReloadExpr struct {
	ReloadPos    Pos
	StatementEnd Pos
	Dictionary   *TableIdentifier
	Type         string
}

func (s *SystemReloadExpr) Pos() Pos {
	return s.ReloadPos
}

func (s *SystemReloadExpr) End() Pos {
	return s.StatementEnd
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

type SystemSyncExpr struct {
	SyncPos Pos
	Cluster *TableIdentifier
}

func (s *SystemSyncExpr) Pos() Pos {
	return s.SyncPos
}

func (s *SystemSyncExpr) End() Pos {
	return s.Cluster.End()
}

func (p *PrintVisitor) VisitSystemSyncExpr(s *SystemSyncExpr) error {
	builder := p.builder
	builder.WriteString("SYNC ")
	builder.WriteString(s.Cluster.String())
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

type SystemCtrlExpr struct {
	CtrlPos      Pos
	StatementEnd Pos
	Command      string // START, STOP
	Type         string // REPLICATED, DISTRIBUTED
	Cluster      *TableIdentifier
}

func (s *SystemCtrlExpr) Pos() Pos {
	return s.CtrlPos
}

func (s *SystemCtrlExpr) End() Pos {
	return s.StatementEnd
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

type SystemDropExpr struct {
	DropPos      Pos
	StatementEnd Pos
	Type         string
}

func (s *SystemDropExpr) Pos() Pos {
	return s.DropPos
}

func (s *SystemDropExpr) End() Pos {
	return s.StatementEnd
}

func (p *PrintVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	return "DROP " + s.Type
}

func (visitor MyDefaultVisitor) VisitSystemDropExpr(s *SystemDropExpr) error {
	visitor.enter(s)
	defer visitor.leave(s)
	return nil
}

type TruncateTable struct {
	TruncatePos  Pos
	StatementEnd Pos
	IsTemporary  bool
	IfExists     bool
	Name         *TableIdentifier
	OnCluster    *ClusterClause
}

func (t *TruncateTable) Pos() Pos {
	return t.TruncatePos
}

func (t *TruncateTable) End() Pos {
	return t.StatementEnd
}

func (t *TruncateTable) Type() error {
	return "TRUNCATE TABLE"
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

type SampleClause struct {
	SamplePos Pos
	Ratio     *RatioExpr
	Offset    *RatioExpr
}

func (s *SampleClause) Pos() Pos {
	return s.SamplePos
}

func (s *SampleClause) End() Pos {
	if s.Offset != nil {
		return s.Offset.End()
	}
	return s.Ratio.End()
}

func (p *PrintVisitor) VisitSampleClause(s *SampleClause) error {
	builder := p.builder
	builder.WriteString("SAMPLE ")
	builder.WriteString(s.Ratio.String())
	if s.Offset != nil {
		builder.WriteString(" OFFSET ")
		builder.WriteString(s.Offset.String())
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

type DeleteClause struct {
	DeletePos Pos
	Table     *TableIdentifier
	OnCluster *ClusterClause
	WhereExpr Expr
}

func (d *DeleteClause) Pos() Pos {
	return d.DeletePos
}

func (d *DeleteClause) End() Pos {
	return d.WhereExpr.End()
}

func (p *PrintVisitor) VisitDeleteClause(d *DeleteClause) error {
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

type ColumnNamesExpr struct {
	LeftParenPos  Pos
	RightParenPos Pos
	ColumnNames   []NestedIdentifier
}

func (c *ColumnNamesExpr) Pos() Pos {
	return c.LeftParenPos
}

func (c *ColumnNamesExpr) End() Pos {
	return c.RightParenPos
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

type AssignmentValues struct {
	LeftParenPos  Pos
	RightParenPos Pos
	Values        []Expr
}

func (v *AssignmentValues) Pos() Pos {
	return v.LeftParenPos
}

func (v *AssignmentValues) End() Pos {
	return v.RightParenPos
}

func (p *PrintVisitor) VisitAssignmentValues(v *AssignmentValues) error {
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

type InsertStmt struct {
	InsertPos   Pos
	Format      *FormatClause
	Table       Expr
	ColumnNames *ColumnNamesExpr
	Values      []*AssignmentValues
	SelectExpr  *SelectQuery
}

func (i *InsertStmt) Pos() Pos {
	return i.InsertPos
}

func (i *InsertStmt) End() Pos {
	if i.SelectExpr != nil {
		return i.SelectExpr.End()
	}
	return i.Values[len(i.Values)-1].End()
}

func (p *PrintVisitor) VisitInsertStmt(i *InsertStmt) error {
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

type CheckStmt struct {
	CheckPos  Pos
	Table     *TableIdentifier
	Partition *PartitionClause
}

func (c *CheckStmt) Pos() Pos {
	return c.CheckPos
}

func (c *CheckStmt) End() Pos {
	return c.Partition.End()
}

func (p *PrintVisitor) VisitCheckStmt(c *CheckStmt) error {
	builder := p.builder
	builder.WriteString("CHECK TABLE ")
	builder.WriteString(c.Table.String())
	if c.Partition != nil {
		builder.WriteString(" ")
		builder.WriteString(c.Partition.String())
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

type UnaryExpr struct {
	UnaryPos Pos
	Kind     TokenKind
	Expr     Expr
}

func (n *UnaryExpr) Pos() Pos {
	return n.UnaryPos
}

func (n *UnaryExpr) End() Pos {
	return n.Expr.End()
}

func (p *PrintVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	return "-" + n.Expr.String()
}

func (visitor MyDefaultVisitor) VisitUnaryExpr(n *UnaryExpr) error {
	visitor.enter(n)
	defer visitor.leave(n)
	if err := n.Expr.Accept(visitor); err != nil {
		return err
	}
	return nil
}

type RenameStmt struct {
	RenamePos    Pos
	StatementEnd Pos

	RenameTarget   string
	TargetPairList []*TargetPair
	OnCluster      *ClusterClause
}

func (r *RenameStmt) Pos() Pos {
	return r.RenamePos
}

func (r *RenameStmt) End() Pos {
	return r.StatementEnd
}

func (r *RenameStmt) Type() error {
	return "RENAME " + r.RenameTarget
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

type TargetPair struct {
	Old *TableIdentifier
	New *TableIdentifier
}

func (t *TargetPair) Pos() Pos {
	return t.Old.Pos()
}

func (t *TargetPair) End() Pos {
	return t.New.End()
}

func (p *PrintVisitor) VisitTargetPair(t *TargetPair) error {
	return t.Old.String() + " TO " + t.New.String()
}

type ExplainStmt struct {
	ExplainPos Pos
	Type       string
	Statement  Expr
}

func (e *ExplainStmt) Pos() Pos {
	return e.ExplainPos
}

func (e *ExplainStmt) End() Pos {
	return e.Statement.End()
}

func (p *PrintVisitor) VisitExplainStmt(e *ExplainStmt) error {
	builder := p.builder
	builder.WriteString("EXPLAIN ")
	builder.WriteString(e.Type)
	builder.WriteByte(' ')
	builder.WriteString(e.Statement.String())
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

type PrivilegeClause struct {
	PrivilegePos Pos
	PrivilegeEnd Pos
	Keywords     []string
	Params       *ParamExprList
}

func (p *PrivilegeClause) Pos() Pos {
	return p.PrivilegePos
}

func (p *PrivilegeClause) End() Pos {
	return p.PrivilegeEnd
}

func (p *PrintVisitor) VisitPrivilegeClause(p *PrivilegeClause) error {
	builder := p.builder
	for i, keyword := range p.Keywords {
		if i > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(keyword)
	}
	if p.Params != nil {
		builder.WriteString(p.Params.String())
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

type GrantPrivilegeStmt struct {
	GrantPos     Pos
	StatementEnd Pos
	OnCluster    *ClusterClause
	Privileges   []*PrivilegeClause
	On           *TableIdentifier
	To           []*Ident
	WithOptions  []string
}

func (g *GrantPrivilegeStmt) Pos() Pos {
	return g.GrantPos
}

func (g *GrantPrivilegeStmt) End() Pos {
	return g.StatementEnd
}

func (g *GrantPrivilegeStmt) Type() error {
	return "GRANT PRIVILEGE"
}

func (p *PrintVisitor) VisitGrantPrivilegeStmt(g *GrantPrivilegeStmt) error {
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
