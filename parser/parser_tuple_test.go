package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_SingleElementTuple verifies that the trailing comma of a
// parenthesised tuple survives printing (#54). `(x,)` is a one-element Tuple
// in ClickHouse, while `(x)` is just x, so dropping the comma changes the type.
func TestParser_SingleElementTuple(t *testing.T) {
	for _, sql := range []string{
		"SELECT (3,)",
		"SELECT ((3,),)",
		"SELECT (1, 2,)",
		"SELECT (*,)",
		"SELECT (*,).1",
		"SELECT CAST((3,) AS Nullable(Tuple(Int32)))",
		"SELECT 1 FROM t0 JOIN t0 ON (*,)",
		"SELECT 1 WHERE (a, b) IN ((1, 2),)",
		"SELECT count() FROM t GROUP BY (a,)",
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

// TestParser_TrailingCommaRejected lists places where ClickHouse 26.8 rejects
// a trailing comma with a syntax error: everything except a bare tuple.
func TestParser_TrailingCommaRejected(t *testing.T) {
	for _, sql := range []string{
		"SELECT tuple(3,)",
		"SELECT plus(1, 2,)",
		"SELECT [1, 2,]",
		"SELECT quantiles(0.5,)(x) FROM t",
		"SELECT quantiles(0.5)(x,) FROM t",
		"SELECT a FROM t GROUP BY CUBE(a,)",
		"SELECT a FROM t GROUP BY GROUPING SETS ((a),)",
		"CREATE TABLE t (x UInt8) ENGINE = MergeTree(x,)",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
