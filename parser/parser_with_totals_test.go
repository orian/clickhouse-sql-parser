package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_WithTotalsWithoutGroupBy verifies that WITH TOTALS without GROUP
// BY survives printing (#56). ClickHouse accepts it and returns a totals row
// over the whole query; the parser used to parse it and then drop it.
func TestParser_WithTotalsWithoutGroupBy(t *testing.T) {
	for _, sql := range []string{
		"SELECT count() FROM t WITH TOTALS",
		"SELECT count() AS c FROM t WHERE x = 1 WITH TOTALS SETTINGS totals_mode='before_having'",
		"SELECT count() FROM t WITH TOTALS HAVING count() > 1",
		"SELECT count() FROM t WITH TOTALS ORDER BY 1 LIMIT 1",
		"SELECT a, count() FROM t GROUP BY a WITH TOTALS",
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

	sql := "SELECT count() FROM t WITH TOTALS"
	stmts, err := NewParser(sql).ParseStmts()
	require.NoError(t, err)
	require.True(t, stmts[0].(*SelectQuery).WithTotal)
	require.Equal(t, Pos(len(sql)), stmts[0].End())
}
