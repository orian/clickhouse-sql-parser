package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatementEndOptionalClauses(t *testing.T) {
	for _, sql := range []string{
		"CHECK TABLE t",
		"CHECK TABLE t PARTITION 1",
		"INSERT INTO t",
		"INSERT INTO t (x)",
		"INSERT INTO t FORMAT CSV",
		"INSERT INTO t VALUES (1)",
		"INSERT INTO t SELECT 1",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.Equal(t, Pos(len(sql)), stmts[0].End())
		})
	}
}

func TestDictionaryEmptyFunctionArgument(t *testing.T) {
	sql := "CREATE DICTIONARY d (x UInt64) PRIMARY KEY x SOURCE(CLICKHOUSE(DB currentDatabase() TABLE 't')) LAYOUT(FLAT()) LIFETIME(0)"
	stmts, err := NewParser(sql).ParseStmts()
	require.NoError(t, err)
	require.Len(t, stmts, 1)
	node, found := Find(stmts[0], func(node Expr) bool {
		_, ok := node.(*FunctionExpr)
		return ok
	})
	require.True(t, found)
	require.Equal(t, "currentDatabase()", node.String())
	require.Equal(t, Pos(strings.Index(sql, "currentDatabase()")+len("currentDatabase()")), node.End())
	printer := NewPrintVisitor()
	require.NoError(t, stmts[0].Accept(printer))
	require.Contains(t, printer.String(), "currentDatabase()")
}
