package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_ParametricAggregateDistinct verifies that DISTINCT in the
// argument list of a parametric aggregate survives printing (#55). Dropping
// it changes the result: groupArraySample(5, 1)(DISTINCT x) over
// [1,1,1,1,1,2] returns [1,2], without DISTINCT [2,1,1,1,1].
func TestParser_ParametricAggregateDistinct(t *testing.T) {
	for _, sql := range []string{
		"SELECT groupArraySample(5, 11111)(DISTINCT subdomain) FROM t",
		"SELECT quantiles(0.5)(DISTINCT x) FROM t",
		"SELECT quantilesTimingIf(0.1, 0.5)(DISTINCT x, y) FROM t",
		"SELECT quantile(0.5)(x) FROM t",
		"SELECT count(DISTINCT x) FROM t",
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
