package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_ParenthesisedUnionOperand verifies that a parenthesised operand
// of UNION/EXCEPT keeps its parentheses (#59). They group the operand's own
// UNION chain: `a UNION DISTINCT (b UNION ALL c)` is not the same query as
// `a UNION DISTINCT b UNION ALL c`.
func TestParser_ParenthesisedUnionOperand(t *testing.T) {
	for _, sql := range []string{
		"SELECT 1 UNION ALL (SELECT 2 UNION ALL SELECT 3)",
		"SELECT 1 UNION DISTINCT (SELECT 1 UNION ALL SELECT 1)",
		"SELECT 1 UNION ALL (SELECT 2)",
		"SELECT * FROM (SELECT 1 UNION ALL (SELECT 2 UNION ALL SELECT 3))",
		"WITH x AS (SELECT 1) SELECT * FROM x",
		"WITH x AS (SELECT 1 UNION ALL (SELECT 2)) SELECT * FROM x",
		"SELECT * FROM (WITH RECURSIVE q AS (SELECT 1 FROM t0 UNION ALL (WITH RECURSIVE x AS (SELECT 1 FROM t0) SELECT 1 FROM x)) SELECT 1 FROM q)",
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
