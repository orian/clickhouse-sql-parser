package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_EmptyColumnList verifies that an empty column list `()` is
// rejected, as ClickHouse does, wherever a table schema is accepted (#38).
func TestParser_EmptyColumnList(t *testing.T) {
	for _, sql := range []string{
		"CREATE TABLE t () ENGINE = Memory",
		"CREATE TABLE t () ENGINE = Memory AS SELECT 1",
		"CREATE TEMPORARY TABLE t () ENGINE = Memory",
		"ATTACH TABLE t () ENGINE = Memory",
		"CREATE VIEW v () AS SELECT 1",
		"CREATE MATERIALIZED VIEW v TO t () AS SELECT 1",
		"CREATE TABLE ts ENGINE = TimeSeries SAMPLES INNER COLUMNS ()",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

// TestParser_TableSchemaAsSpacing verifies that the `AS <table>` and
// `AS <table function>` schema forms print with a single space (#38).
func TestParser_TableSchemaAsSpacing(t *testing.T) {
	for _, sql := range []string{
		"CREATE TABLE t AS other ENGINE = Memory",
		"CREATE TABLE t AS db.other",
		"CREATE TABLE t AS remote('h', db, t)",
		"CREATE TABLE t ON CLUSTER c AS db.other ENGINE = Distributed(c, db, other, rand())",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, sql, stmts[0].String())

			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())

			beautify := NewBeautifyVisitor()
			require.NoError(t, stmts[0].Accept(beautify))
			require.NotContains(t, beautify.String(), "  AS")
		})
	}
}
