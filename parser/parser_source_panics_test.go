package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_ClickHouseSourcePanics(t *testing.T) {
	files, err := filepath.Glob("testdata/regressions/clickhouse_panics/*.sql")
	require.NoError(t, err)
	require.Len(t, files, 56)
	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			sql, err := os.ReadFile(file)
			require.NoError(t, err)
			// Upstream scripts can contain deliberately invalid or unsupported SQL.
			// Either an AST or an ordinary error is valid here; a panic never is.
			require.NotPanics(t, func() { _, _ = NewParser(string(sql)).ParseStmts() })
		})
	}
}

func TestParser_NumericSeparators(t *testing.T) {
	for _, sql := range []string{
		"SELECT sum(prime) FROM primes(20_000_000)",
		"SET max_ast_depth = 10_000_000",
		"SELECT 1 SETTINGS max_rows_to_read = 50_000_000",
		"SELECT 1_000, 1_000.25_50, 1e1_0, 0xCA_FE",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.Contains(t, stmts[0].String(), "_")
			_, err = NewParser(stmts[0].String()).ParseStmts()
			require.NoError(t, err)
		})
	}
}

func TestParser_InvalidTableArgumentsAndSettings(t *testing.T) {
	for _, sql := range []string{
		"SELECT * FROM numbers(1e)", "SELECT * FROM numbers('unterminated)",
		"SELECT * FROM numbers(n=)", "SELECT * FROM numbers(n=1e)",
		"SET x =", "SET x = 1e", "SELECT 1 SETTINGS x = 1e",
		"SELECT timestamp, url FROM table WHERE visitor_id = 1001`",
	} {
		t.Run(sql, func(t *testing.T) {
			var err error
			require.NotPanics(t, func() { _, err = NewParser(sql).ParseStmts() })
			require.Error(t, err)
		})
	}
}

func TestParser_PartitionAllEnd(t *testing.T) {
	for _, sql := range []string{
		"ALTER TABLE t ATTACH PARTITION ALL",
		"ALTER TABLE t ATTACH PARTITION /* gap */ ALL",
		"ALTER TABLE t DETACH PARTITION ALL",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.Equal(t, Pos(len(sql)), stmts[0].End())
			found := false
			Walk(stmts[0], func(e Expr) bool {
				if partition, ok := e.(*PartitionClause); ok {
					found = true
					require.True(t, partition.All)
					require.Equal(t, Pos(len(sql)), partition.End())
				}
				return true
			})
			require.True(t, found)
			_, err = NewParser(stmts[0].String()).ParseStmts()
			require.NoError(t, err)
		})
	}
}

func TestParser_PropagatesLexerErrors(t *testing.T) {
	for _, sql := range []string{"SET x = 1e", "SELECT * FROM numbers(1e)", "SELECT * FROM numbers(n=1e)"} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.ErrorContains(t, err, "exponent part should contain at least one digit")
		})
	}
}
