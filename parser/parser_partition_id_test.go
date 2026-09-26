package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_PartitionID verifies that PARTITION ID 'x' keeps ID (#58).
// PARTITION ID 'x' selects a partition by ID, PARTITION 'x' by the value of
// the partition expression, so dropping ID can target another partition.
func TestParser_PartitionID(t *testing.T) {
	for _, sql := range []string{
		"OPTIMIZE TABLE t PARTITION ID 'all' FINAL",
		"OPTIMIZE TABLE t PARTITION 'all'",
		"ALTER TABLE t DROP PARTITION ID '202401'",
		"ALTER TABLE t DROP PARTITION 202401",
		"ALTER TABLE t DROP DETACHED PARTITION ID '202401'",
		"ALTER TABLE t DETACH PARTITION ID '202401'",
		"ALTER TABLE t ATTACH PARTITION ID '202401' FROM t2",
		"ALTER TABLE t REPLACE PARTITION ID '202401' FROM t2",
		"ALTER TABLE t UPDATE x = 1 IN PARTITION ID '202401' WHERE 1",
		"ALTER TABLE t DELETE IN PARTITION ID '202401' WHERE 1",
		"ALTER TABLE t DELETE WHERE x = 1",
		"ALTER TABLE t CLEAR COLUMN c IN PARTITION ID '202401'",
		"ALTER TABLE t MATERIALIZE INDEX i IN PARTITION ID '202401'",
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
