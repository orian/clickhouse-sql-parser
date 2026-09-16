package parser

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_RecursiveCTEDocumentationExamples(t *testing.T) {
	files, err := filepath.Glob("testdata/recursive_cte/docs/*.sql")
	require.NoError(t, err)
	require.Len(t, files, 8)

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			sql, err := os.ReadFile(file)
			require.NoError(t, err)

			stmts, err := NewParser(string(sql)).ParseStmts()
			require.NoError(t, err)
			require.NotEmpty(t, stmts)

			recursiveStatements := 0
			for _, stmt := range stmts {
				if hasRecursiveWith(stmt) {
					recursiveStatements++
				}
				assertFormatterRoundTrip(t, stmt, func(expr Expr) string { return expr.String() })
				assertFormatterRoundTrip(t, stmt, compactRecursiveSQL)
				assertFormatterRoundTrip(t, stmt, beautifyRecursiveSQL)
			}
			require.Equal(t, 1, recursiveStatements)
		})
	}
}

func TestParser_RecursiveCTEASTAndFormatting(t *testing.T) {
	const input = "with recursive t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 3) SELECT * FROM t"
	const compact = "WITH RECURSIVE t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 3) SELECT * FROM t"
	const beautified = "WITH RECURSIVE\n  t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 3)\nSELECT *\nFROM t"

	stmts, err := NewParser(input).ParseStmts()
	require.NoError(t, err)
	require.Len(t, stmts, 1)

	query, ok := stmts[0].(*SelectQuery)
	require.True(t, ok)
	require.NotNil(t, query.With)
	require.True(t, query.With.HasRecursive)
	require.Equal(t, "WITH RECURSIVE t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 3)", query.With.String())
	require.Equal(t, compact, query.String())
	require.Equal(t, compact, compactRecursiveSQL(query))
	require.Equal(t, beautified, beautifyRecursiveSQL(query))

	encoded, err := json.Marshal(query.With)
	require.NoError(t, err)
	require.Contains(t, string(encoded), `"HasRecursive":true`)
	ordinary, err := NewParser("WITH t AS (SELECT 1) SELECT * FROM t").ParseStmts()
	require.NoError(t, err)
	ordinaryWith := ordinary[0].(*SelectQuery).With
	require.False(t, ordinaryWith.HasRecursive)
	encoded, err = json.Marshal(ordinaryWith)
	require.NoError(t, err)
	require.NotContains(t, string(encoded), "HasRecursive")

	for _, formatted := range []string{query.String(), compactRecursiveSQL(query), beautifyRecursiveSQL(query)} {
		reparsed, err := NewParser(formatted).ParseStmts()
		require.NoError(t, err)
		require.Len(t, reparsed, 1)
		require.True(t, reparsed[0].(*SelectQuery).With.HasRecursive)
	}
}

func TestParser_RecursiveCTESyntaxQuirks(t *testing.T) {
	tests := []struct {
		name          string
		sql           string
		recursiveWith int
	}{
		{
			name:          "multiple CTEs",
			sql:           "WITH RECURSIVE a AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM a WHERE n < 2), b AS (SELECT n FROM a) SELECT * FROM b",
			recursiveWith: 1,
		},
		{
			name:          "nested query",
			sql:           "SELECT * FROM (WITH RECURSIVE t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 2) SELECT * FROM t)",
			recursiveWith: 1,
		},
		{
			name:          "case insensitive modifier",
			sql:           "WiTh ReCuRsIvE t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 2) SELECT * FROM t",
			recursiveWith: 1,
		},
		{
			name:          "quoted keyword CTE name",
			sql:           "WITH RECURSIVE `select` AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM `select` WHERE n < 2) SELECT * FROM `select`",
			recursiveWith: 1,
		},
		{
			name:          "double quoted recursive CTE name",
			sql:           `WITH RECURSIVE "recursive" AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM "recursive" WHERE n < 2) SELECT * FROM "recursive"`,
			recursiveWith: 1,
		},
		{
			name:          "quoted recursive is ordinary CTE name",
			sql:           "WITH `recursive` AS (SELECT 1) SELECT * FROM `recursive`",
			recursiveWith: 0,
		},
		{
			name:          "ordinary WITH unchanged",
			sql:           "WITH t AS (SELECT 1) SELECT * FROM t",
			recursiveWith: 0,
		},
		{
			name:          "recursive remains a bare identifier outside WITH modifier position",
			sql:           "SELECT recursive FROM recursive",
			recursiveWith: 0,
		},
		{
			name:          "CTE column aliases",
			sql:           "WITH RECURSIVE t(n) AS (SELECT 1 UNION ALL SELECT n + 1 FROM t WHERE n < 2) SELECT * FROM t",
			recursiveWith: 1,
		},
		{
			name:          "scalar expression and recursive CTE",
			sql:           "WITH RECURSIVE 42 AS answer, t AS (SELECT answer AS n UNION ALL SELECT n + 1 FROM t WHERE n < 43) SELECT * FROM t",
			recursiveWith: 1,
		},
		{
			name:          "recursive modifier with scalar expression only",
			sql:           "WITH RECURSIVE 1 AS n SELECT n",
			recursiveWith: 1,
		},
		{
			name:          "recursive WITH trailing comma",
			sql:           "WITH RECURSIVE t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 2), SELECT * FROM t",
			recursiveWith: 1,
		},
		{
			name:          "ordinary WITH trailing comma",
			sql:           "WITH 1 AS n, SELECT n",
			recursiveWith: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts, err := NewParser(tt.sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)

			withClauses := collectWithClauses(stmts[0])
			recursiveWith := 0
			for _, with := range withClauses {
				if with.HasRecursive {
					recursiveWith++
				}
			}
			require.Equal(t, tt.recursiveWith, recursiveWith)

			compact := compactRecursiveSQL(stmts[0])
			reparsed, err := NewParser(compact).ParseStmts()
			require.NoError(t, err)
			require.Len(t, reparsed, 1)
			require.Equal(t, compact, compactRecursiveSQL(reparsed[0]))
		})
	}
}

func TestParser_RecursiveCTEInvalidModifierPlacement(t *testing.T) {
	for _, sql := range []string{
		"WITH recursive AS (SELECT 1) SELECT * FROM recursive",
		"WITH RECURSIVE RECURSIVE t AS (SELECT 1) SELECT * FROM t",
		"WITH t AS (SELECT 1), RECURSIVE r AS (SELECT 1) SELECT * FROM r",
		"WITH t AS (SELECT 1), select AS (SELECT 2) SELECT * FROM t",
	} {
		t.Run(sql, func(t *testing.T) {
			require.NotPanics(t, func() {
				_, err := NewParser(sql).ParseStmts()
				require.Error(t, err)
			})
		})
	}
}

func collectWithClauses(expr Expr) []*WithClause {
	var clauses []*WithClause
	Walk(expr, func(current Expr) bool {
		if with, ok := current.(*WithClause); ok {
			clauses = append(clauses, with)
		}
		return true
	})
	return clauses
}

func hasRecursiveWith(expr Expr) bool {
	for _, with := range collectWithClauses(expr) {
		if with.HasRecursive {
			return true
		}
	}
	return false
}

func compactRecursiveSQL(expr Expr) string {
	visitor := NewPrintVisitor()
	if err := expr.Accept(visitor); err != nil {
		panic(err)
	}
	return visitor.String()
}

func beautifyRecursiveSQL(expr Expr) string {
	visitor := NewBeautifyVisitor()
	if err := expr.Accept(visitor); err != nil {
		panic(err)
	}
	return visitor.String()
}

func assertFormatterRoundTrip(t *testing.T, stmt Expr, format func(Expr) string) {
	t.Helper()
	formatted := format(stmt)
	if hasRecursiveWith(stmt) {
		require.Contains(t, strings.ToUpper(formatted), "WITH RECURSIVE")
	}

	reparsed, err := NewParser(formatted).ParseStmts()
	require.NoError(t, err)
	require.Len(t, reparsed, 1)
	require.Equal(t, formatted, format(reparsed[0]))
}
