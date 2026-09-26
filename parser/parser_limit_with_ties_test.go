package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_LimitWithTies covers `LIMIT n [OFFSET m] WITH TIES` (#48) and
// `TOP n WITH TIES`.
// ClickHouse checks that ORDER BY is present only at analysis time
// (Code 476), so a LIMIT ... WITH TIES without ORDER BY is valid syntax.
func TestParser_LimitWithTies(t *testing.T) {
	for _, sql := range []string{
		"SELECT a FROM t ORDER BY a LIMIT 1 WITH TIES",
		"SELECT a FROM t ORDER BY a LIMIT 1 OFFSET 2 WITH TIES",
		"SELECT a FROM t ORDER BY a LIMIT 1 BY a LIMIT 3 WITH TIES",
		"SELECT a FROM t ORDER BY a LIMIT 1 WITH TIES SETTINGS max_threads=1",
		"SELECT a FROM t LIMIT 1 WITH TIES",
		// Valid syntax; ClickHouse rejects it at analysis time with
		// LIMIT_BY_WITH_TIES_IS_NOT_SUPPORTED (Code 498).
		"SELECT a FROM t ORDER BY a LIMIT 1 WITH TIES BY a",
		// TOP n WITH TIES used to print as `WITH TIES`, dropping TOP n.
		"SELECT TOP 1 WITH TIES * FROM t ORDER BY a",
		"SELECT DISTINCT TOP 1 WITH TIES * FROM t0 ORDER BY tuple()",
		"SELECT a FROM t ORDER BY a LIMIT 1 OFFSET 2",
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

	// `LIMIT m, n` is printed in its OFFSET form.
	stmts, err := NewParser("SELECT a FROM t ORDER BY a LIMIT 2, 1 WITH TIES").ParseStmts()
	require.NoError(t, err)
	require.Equal(t, "SELECT a FROM t ORDER BY a LIMIT 1 OFFSET 2 WITH TIES", stmts[0].String())

	sql := "SELECT a FROM t ORDER BY a LIMIT 1 WITH TIES"
	stmts, err = NewParser(sql).ParseStmts()
	require.NoError(t, err)
	limit := stmts[0].(*SelectQuery).Limit
	require.True(t, limit.WithTies)
	require.Equal(t, Pos(len(sql)), limit.End())
}

// TestParser_LimitWithTiesRejected lists forms ClickHouse 26.8 rejects with a
// syntax error.
func TestParser_LimitWithTiesRejected(t *testing.T) {
	for _, sql := range []string{
		"SELECT a FROM t ORDER BY a LIMIT 1 BY a WITH TIES",
		"SELECT a FROM t ORDER BY a LIMIT 1 WITH TIES WITH TIES",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
