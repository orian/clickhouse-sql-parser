package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_TTLAssignments(t *testing.T) {
	for _, sql := range []string{
		"CREATE TABLE hits (timestamp DateTime, id String, hits Int64) ENGINE = MergeTree ORDER BY (id, timestamp) TTL timestamp + INTERVAL 1 DAY GROUP BY id SET hits = sum(hits), timestamp + INTERVAL 1 MONTH DELETE;",
		"ALTER TABLE hits MODIFY TTL timestamp + INTERVAL 1 DAY GROUP BY id SET hits = sum(hits), timestamp + INTERVAL 1 MONTH GROUP BY id SET hits = max(hits);",
		"ALTER TABLE hits MODIFY TTL timestamp + INTERVAL 1 DAY GROUP BY id SET `hits` = sumIf(hits, id IN ('a', 'b'));",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			validFormatSQL(t, stmts[0].String()+";\n")
			beautify := NewBeautifyVisitor()
			require.NoError(t, stmts[0].Accept(beautify))
			_, err = NewParser(beautify.String()).ParseStmts()
			require.NoError(t, err)

			var assignments []*UpdateAssignment
			Walk(stmts[0], func(expr Expr) bool {
				if policy, ok := expr.(*TTLPolicy); ok && len(policy.Assignments) > 0 {
					last := policy.Assignments[len(policy.Assignments)-1]
					require.Equal(t, last.End(), policy.End())
					require.Greater(t, policy.End(), policy.GroupBy.End())
				}
				if assignment, ok := expr.(*UpdateAssignment); ok {
					assignments = append(assignments, assignment)
				}
				return true
			})
			require.NotEmpty(t, assignments)
			visitor := &ttlAssignmentVisitor{err: errors.New("stop at assignment")}
			visitor.Self = visitor
			require.ErrorIs(t, stmts[0].Accept(visitor), visitor.err)
		})
	}
}

type ttlAssignmentVisitor struct {
	DefaultASTVisitor
	err error
}

func (v *ttlAssignmentVisitor) VisitUpdateAssignment(*UpdateAssignment) error {
	return v.err
}

func TestParser_InvalidTTLAssignments(t *testing.T) {
	for _, clause := range []string{
		"SET hits = sum(hits)",
		"GROUP BY id SET",
		"GROUP BY id SET hits",
		"GROUP BY id SET hits =",
		"GROUP BY id SET hits = sum(hits),",
		"GROUP BY id SET hits = sum(hits), other =",
	} {
		t.Run(clause, func(t *testing.T) {
			_, err := NewParser("ALTER TABLE hits MODIFY TTL timestamp + INTERVAL 1 DAY " + clause).ParseStmts()
			require.Error(t, err)
		})
	}
}
