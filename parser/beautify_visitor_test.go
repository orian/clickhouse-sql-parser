package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func beautify(t *testing.T, sql string) string {
	t.Helper()
	stmts, err := NewParser(sql).ParseStmts()
	require.NoError(t, err)
	require.Len(t, stmts, 1)
	v := NewBeautifyVisitor()
	require.NoError(t, stmts[0].Accept(v))
	return v.String()
}

func TestBeautifyVisitor_Select(t *testing.T) {
	got := beautify(t, "SELECT a, b, c FROM t WHERE x > 1 GROUP BY a ORDER BY b LIMIT 10")
	want := "SELECT\n  a,\n  b,\n  c\nFROM t\nWHERE x > 1\nGROUP BY a\nORDER BY b\nLIMIT 10"
	require.Equal(t, want, got)
}

func TestBeautifyVisitor_SelectDistinctOn(t *testing.T) {
	got := beautify(t, "SELECT DISTINCT ON (a) a, b FROM t")
	want := "SELECT DISTINCT ON (a)\n  a,\n  b\nFROM t"
	require.Equal(t, want, got)
}

func TestBeautifyVisitor_WithCTE(t *testing.T) {
	got := beautify(t, "WITH x AS (SELECT 1), y AS (SELECT 2) SELECT * FROM x")
	want := "WITH\n  x AS (SELECT 1),\n  y AS (SELECT 2)\nSELECT\n  *\nFROM x"
	require.Equal(t, want, got)
}

func TestBeautifyVisitor_CreateTable(t *testing.T) {
	got := beautify(t, "CREATE TABLE foo (id UInt64, name String) ENGINE = MergeTree ORDER BY id")
	want := "CREATE TABLE foo (\n  id UInt64,\n  name String\n)\nENGINE = MergeTree ORDER BY id"
	require.Equal(t, want, got)
}

func TestBeautifyVisitor_Insert(t *testing.T) {
	got := beautify(t, "INSERT INTO t (a, b) VALUES (1, 2), (3, 4)")
	want := "INSERT INTO t\n(a, b)\nVALUES\n  (1, 2),\n  (3, 4)"
	require.Equal(t, want, got)
}

// Falls through to compact for unhandled statement types.
func TestBeautifyVisitor_FallthroughForUnhandledStatements(t *testing.T) {
	// USE statement isn't overridden — default descent does nothing, so the
	// builder stays empty. This documents current scope: only the major
	// statement types beautify; everything else needs a VisitX added.
	got := beautify(t, "USE db1")
	require.Equal(t, "", got)
}
