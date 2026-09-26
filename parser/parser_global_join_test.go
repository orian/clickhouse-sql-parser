package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_GlobalJoin verifies that GLOBAL survives printing and combines
// with every join kind (#57). GLOBAL changes distributed execution: the right
// side is computed once and sent to every shard.
func TestParser_GlobalJoin(t *testing.T) {
	for _, sql := range []string{
		"SELECT 1 FROM a GLOBAL JOIN b ON a.x = b.x",
		"SELECT 1 FROM distributed_table1 AS t1 GLOBAL JOIN distributed_table2 AS t2 ON materialize(42) = t1.a",
		"SELECT 1 FROM a GLOBAL LEFT JOIN b ON a.x = b.x",
		"SELECT 1 FROM a GLOBAL ANY LEFT JOIN b ON a.x = b.x",
		"SELECT 1 FROM a GLOBAL ALL INNER JOIN b USING (x)",
		"SELECT 1 FROM a GLOBAL CROSS JOIN b",
		"SELECT 1 FROM a GLOBAL ASOF LEFT JOIN b USING (x, t)",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, sql, stmts[0].String())

			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())

			beautifyRoundTrip(t, stmts[0])
		})
	}
}

// TestParser_LocalIsTableAlias checks that LOCAL before JOIN aliases the
// left table, as in ClickHouse 26.8 (`FROM a AS LOCAL INNER JOIN b`), instead
// of being dropped as a join keyword.
func TestParser_LocalIsTableAlias(t *testing.T) {
	stmts, err := NewParser("SELECT 1 FROM a LOCAL JOIN b ON LOCAL.x = b.x").ParseStmts()
	require.NoError(t, err)
	require.Equal(t, "SELECT 1 FROM a AS LOCAL JOIN b ON LOCAL.x = b.x", stmts[0].String())
	beautifyRoundTrip(t, stmts[0])
}

// TestParser_JoinUsingParentheses verifies USING keeps its parentheses so a
// following comma join is not read as another USING column (#73).
func TestParser_JoinUsingParentheses(t *testing.T) {
	for _, sql := range []string{
		"SELECT 1 FROM a JOIN b USING (x), c",
		"SELECT 1 FROM a JOIN b USING (x, t)",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, sql, stmts[0].String())
			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())
			beautifyRoundTrip(t, stmts[0])
		})
	}
	// Unparenthesised input is accepted and printed with parentheses.
	stmts, err := NewParser("SELECT 1 FROM a JOIN b USING x, t").ParseStmts()
	require.NoError(t, err)
	require.Equal(t, "SELECT 1 FROM a JOIN b USING (x, t)", stmts[0].String())
}

func TestParser_GlobalJoinRejected(t *testing.T) {
	for _, sql := range []string{
		"SELECT 1 FROM a LEFT GLOBAL JOIN b ON a.x = b.x",
		"SELECT 1 FROM a GLOBAL ARRAY JOIN arr",
		"SELECT 1 FROM a GLOBAL LEFT ARRAY JOIN arr",
		"SELECT 1 FROM a GLOBAL b",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
