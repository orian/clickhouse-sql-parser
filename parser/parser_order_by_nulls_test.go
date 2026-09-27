package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_OrderByNullsAndCollate covers the ORDER BY element modifiers
// `[ASC|DESC] [NULLS FIRST|LAST] [COLLATE 'locale'] [WITH FILL ...]` (#47).
// NULLS and COLLATE were rejected in every position. Every statement here is
// accepted by ClickHouse 26.8.
func TestParser_OrderByNullsAndCollate(t *testing.T) {
	for _, sql := range []string{
		"SELECT 1 FROM t ORDER BY a NULLS FIRST",
		"SELECT 1 FROM t ORDER BY a NULLS LAST",
		"SELECT 1 FROM t ORDER BY a DESC NULLS FIRST",
		"SELECT 1 FROM t ORDER BY a ASC NULLS LAST",
		"SELECT 1 FROM t ORDER BY a NULLS FIRST, b DESC",
		"SELECT 1 FROM t ORDER BY a COLLATE 'en'",
		"SELECT 1 FROM t ORDER BY a NULLS FIRST COLLATE 'en'",
		"SELECT 1 FROM t ORDER BY a DESC NULLS LAST COLLATE 'en'",
		"SELECT 1 FROM t ORDER BY a NULLS FIRST WITH FILL",
		"SELECT 1 FROM t ORDER BY a DESC NULLS LAST COLLATE 'en' WITH FILL FROM 1 TO 10 STEP 1",
		"SELECT a, count() OVER (ORDER BY a NULLS FIRST) FROM t",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, sql, stmts[0].String())
			require.NoError(t, checkStatementSpan(sql, stmts[0]))

			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())

			beautifyRoundTrip(t, stmts[0])
		})
	}

	// DESCENDING is printed as DESC.
	stmts, err := NewParser("SELECT 1 FROM t ORDER BY a DESCENDING NULLS FIRST").ParseStmts()
	require.NoError(t, err)
	require.Equal(t, "SELECT 1 FROM t ORDER BY a DESC NULLS FIRST", stmts[0].String())

	order := stmts[0].(*SelectQuery).OrderBy.Items[0].(*OrderExpr)
	require.Equal(t, "FIRST", order.Nulls)

	// PrintVisitor used to drop WITH FILL when printing an ORDER BY element.
	stmts, err = NewParser("SELECT 1 FROM t ORDER BY a WITH FILL").ParseStmts()
	require.NoError(t, err)
	printer := NewPrintVisitor()
	require.NoError(t, stmts[0].(*SelectQuery).OrderBy.Items[0].Accept(printer))
	require.Equal(t, "a WITH FILL", printer.String())
}

// TestParser_OrderByNullsRejected lists forms ClickHouse 26.8 rejects with a
// syntax error.
func TestParser_OrderByNullsRejected(t *testing.T) {
	for _, sql := range []string{
		"SELECT 1 FROM t ORDER BY a NULLS LAST DESC",
		"SELECT 1 FROM t ORDER BY a NULLS",
		"SELECT 1 FROM t ORDER BY a NULLS FIRST NULLS LAST",
		"SELECT 1 FROM t ORDER BY a COLLATE 'en' NULLS FIRST",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

// TestPrintVisitor_OrderExprStreams checks, for every ORDER BY element and
// WITH FILL clause in the fixtures, that PrintVisitor (which streams into one
// builder and recurses via Accept) prints the same SQL as String().
func TestPrintVisitor_OrderExprStreams(t *testing.T) {
	for _, dir := range sqlFixtureDirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
		require.NoError(t, err)
		for _, file := range files {
			src, err := os.ReadFile(file)
			require.NoError(t, err)
			stmts, err := NewParser(string(src)).ParseStmts()
			if err != nil {
				continue
			}
			for _, stmt := range stmts {
				Walk(stmt, func(node Expr) bool {
					switch node.(type) {
					case *OrderExpr, *Fill:
						printer := NewPrintVisitor()
						require.NoError(t, node.Accept(printer))
						require.Equal(t, node.String(), printer.String(), "%s: %T", file, node)
					}
					return true
				})
			}
		}
	}
}
