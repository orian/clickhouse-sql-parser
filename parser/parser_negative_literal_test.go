package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_SignedNumberLiteral verifies that a sign is printed directly
// against a numeric literal (#61). ClickHouse turns `<literal>::T` into
// CAST('<literal source text>', 'T'), so `-1::Int32` must not become
// `- 1::Int32`, which ClickHouse evaluates as CAST('- 1', 'Int32') and
// rejects with CANNOT_PARSE_NUMBER.
func TestParser_SignedNumberLiteral(t *testing.T) {
	for _, sql := range []string{
		"SELECT -1::Int32",
		"SELECT -1.5::Float64",
		"SELECT -1e-5::Float64",
		"SELECT -128::Int8",
		"SELECT -0x10::Int32",
		"SELECT -1::Int32 + 2",
		"SELECT +1",
		"SELECT -1",
		"SELECT 1 - 1::Int32",
		"SELECT 2 - -1",
		"SELECT (-1)::Int32",
		// Non-literal operands keep the space.
		"SELECT - x::Int32",
		"SELECT - (1)::Int32",
		"SELECT NOT 1",
		// `- -1` must not become `--1`, which starts a comment.
		"SELECT - -1",
		"SELECT - +1",
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

	// Input written with a space is printed without it.
	stmts, err := NewParser("SELECT - 1::Int32").ParseStmts()
	require.NoError(t, err)
	require.Equal(t, "SELECT -1::Int32", stmts[0].String())
}
