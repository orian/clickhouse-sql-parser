package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_UnionGroups covers parenthesised set-operation operands followed
// by more UNION/EXCEPT operands (#78) and statements starting with a
// parenthesised query. Every statement is accepted by ClickHouse 26.8 and
// formats back to the same query.
func TestParser_UnionGroups(t *testing.T) {
	for _, sql := range []string{
		"(SELECT 1) UNION ALL SELECT 2",
		"((SELECT 1)) UNION ALL SELECT 2",
		"(SELECT 1 UNION ALL SELECT 2) UNION ALL SELECT 3",
		"(SELECT 1 UNION DISTINCT SELECT 1) UNION ALL SELECT 1",
		"SELECT 1 UNION ALL (SELECT 2 UNION ALL SELECT 3) UNION ALL SELECT 4",
		"SELECT 1 UNION DISTINCT (SELECT 1 UNION ALL SELECT 1) UNION ALL SELECT 1",
		"((SELECT 1 UNION ALL SELECT 2) UNION ALL SELECT 3) UNION ALL SELECT 4",
		"SELECT * FROM ((SELECT 1) UNION ALL SELECT 2)",
		"SELECT * FROM ((SELECT 1 UNION ALL SELECT 2) EXCEPT SELECT 2)",
		"(SELECT 1) EXCEPT (SELECT 2)",
		"(SELECT 1)",
		"(SELECT 1) FORMAT JSON",
		"(SELECT 1) FORMAT JSON SETTINGS max_block_size=1",
		"(SELECT 1) SETTINGS max_threads=1",
		"(SELECT 1 UNION ALL SELECT 2) SETTINGS max_threads=1",
		"SELECT 1 UNION ALL (SELECT 2) FORMAT JSON",
		"(SELECT 1 AS x) UNION ALL SELECT 2 SETTINGS max_threads=1",
		"WITH x AS ((SELECT 1) UNION ALL SELECT 2) SELECT * FROM x",
		"EXPLAIN AST (SELECT 1) UNION ALL SELECT 2",
		"CREATE VIEW v AS (SELECT 1) UNION ALL SELECT 2",
		"CREATE VIEW v AS (SELECT 1 UNION ALL SELECT 2) EXCEPT SELECT 2",
		"CREATE MATERIALIZED VIEW mv TO t AS (SELECT 1) UNION ALL SELECT 2",
		"INSERT INTO t (SELECT 1) UNION ALL SELECT 2",
		"INSERT INTO t WITH x AS (SELECT 1) SELECT * FROM x",
		"SELECT * FROM t WHERE a IN ((SELECT 1) UNION ALL SELECT 2)",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.Equal(t, sql, stmts[0].String())
			require.NoError(t, checkStatementSpan(sql, stmts[0]))

			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())

			beautifyRoundTrip(t, stmts[0])
		})
	}
}

// TestParser_UnionGroupsRejected lists forms ClickHouse 26.8 rejects with a
// syntax error: a group may only be followed by a set operation or the
// query-level SETTINGS/FORMAT clauses.
func TestParser_UnionGroupsRejected(t *testing.T) {
	for _, sql := range []string{
		"(SELECT 1) UNION ALL (SELECT 2) ORDER BY 1",
		"(SELECT 1) LIMIT 1",
		"(SELECT 1) SETTINGS max_threads = 1 FORMAT JSON SETTINGS max_block_size = 1",
		"WITH a AS (SELECT 1) (SELECT * FROM a) UNION ALL SELECT 2",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

// TestParser_UnionGroupAST checks the group-node shape: the parenthesised
// query sits in Group, and what follows the closing parenthesis hangs off the
// group node.
func TestParser_UnionGroupAST(t *testing.T) {
	stmts, err := NewParser("(SELECT 1 UNION ALL SELECT 2) UNION ALL SELECT 3").ParseStmts()
	require.NoError(t, err)
	group := stmts[0].(*SelectQuery)
	require.True(t, group.HasParen)
	require.Empty(t, group.SelectItems)
	require.Equal(t, "SELECT 1 UNION ALL SELECT 2", group.Group.String())
	require.Equal(t, "SELECT 3", group.UnionAll.String())

	stmts, err = NewParser("SELECT 1 UNION ALL (SELECT 2 UNION ALL SELECT 3)").ParseStmts()
	require.NoError(t, err)
	first := stmts[0].(*SelectQuery)
	require.False(t, first.HasParen)
	require.NotNil(t, first.UnionAll.Group)
	require.Equal(t, "SELECT 2 UNION ALL SELECT 3", first.UnionAll.Group.String())
}
