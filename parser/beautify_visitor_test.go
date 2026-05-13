package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"
)

// TestParser_Beautify runs every input SQL fixture through BeautifyVisitor and
// compares it to upstream's beautify golden file. Goldens live under each
// fixture directory's `format/beautify/` subdir.
func TestParser_Beautify(t *testing.T) {
	for _, dir := range []string{"./testdata/dml", "./testdata/ddl", "./testdata/query", "./testdata/basic"} {
		outputDir := dir + "/format/beautify"
		entries, err := os.ReadDir(dir)
		require.NoError(t, err)
		for _, entry := range entries {
			if !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}
			// Skip if no upstream beautify golden exists for this input.
			if _, err := os.Stat(filepath.Join(outputDir, entry.Name())); err != nil {
				continue
			}
			t.Run(entry.Name(), func(t *testing.T) {
				fileBytes, err := os.ReadFile(filepath.Join(dir, entry.Name()))
				require.NoError(t, err)
				stmts, err := NewParser(string(fileBytes)).ParseStmts()
				require.NoError(t, err)
				var builder strings.Builder
				builder.WriteString("-- Origin SQL:\n")
				builder.Write(fileBytes)
				builder.WriteString("\n\n-- Beautify SQL:\n")
				for _, stmt := range stmts {
					v := NewBeautifyVisitor()
					require.NoError(t, stmt.Accept(v))
					builder.WriteString(v.String())
					builder.WriteByte(';')
					builder.WriteByte('\n')
				}
				g := goldie.New(t,
					goldie.WithNameSuffix(""),
					goldie.WithDiffEngine(goldie.ColoredDiff),
					goldie.WithFixtureDir(outputDir))
				g.Assert(t, entry.Name(), []byte(builder.String()))
			})
		}
	}
}

func beautify(t *testing.T, sql string) string {
	t.Helper()
	stmts, err := NewParser(sql).ParseStmts()
	require.NoError(t, err)
	require.Len(t, stmts, 1)
	v := NewBeautifyVisitor()
	require.NoError(t, stmts[0].Accept(v))
	return v.String()
}

// Quick shape checks for the major statement types. Detailed pinning lives in
// the golden fixtures under parser/testdata/**/format/beautify/ (see
// TestParser_Beautify above).

func TestBeautifyVisitor_Select(t *testing.T) {
	got := beautify(t, "SELECT a, b FROM t WHERE x > 1 AND y < 10")
	require.Contains(t, got, "SELECT\n")
	require.Contains(t, got, "FROM\n  t")
	require.Contains(t, got, "WHERE\n  x > 1\nAND\n  y < 10")
}

func TestBeautifyVisitor_WithCTE(t *testing.T) {
	got := beautify(t, "WITH x AS (SELECT 1) SELECT * FROM x")
	require.Contains(t, got, "WITH\n  x AS")
}

func TestBeautifyVisitor_CreateTable(t *testing.T) {
	got := beautify(t, "CREATE TABLE foo (id UInt64, name String) ENGINE = MergeTree ORDER BY id")
	require.Contains(t, got, "CREATE TABLE foo (\n  id UInt64,\n  name String\n)")
}

func TestBeautifyVisitor_Insert(t *testing.T) {
	got := beautify(t, "INSERT INTO t (a, b) VALUES (1, 2), (3, 4)")
	require.Contains(t, got, "INSERT INTO t\n  (a, b)\nVALUES\n  (1, 2),\n  (3, 4)")
}

// Falls through to nothing for unhandled statement types — documents scope:
// only major statement types have VisitX overrides; add more on demand.
func TestBeautifyVisitor_FallthroughForUnhandledStatements(t *testing.T) {
	got := beautify(t, "USE db1")
	require.Equal(t, "", got)
}

// SETTINGS clauses break one item per indented line, both at SELECT level
// and inside an engine spec.
func TestBeautifyVisitor_SelectSettings(t *testing.T) {
	got := beautify(t, "SELECT * FROM t SETTINGS max_threads=8, max_memory_usage=10000000")
	require.Contains(t, got, "SETTINGS\n  max_threads=8,\n  max_memory_usage=10000000")
}

func TestBeautifyVisitor_CreateTableEngineSettings(t *testing.T) {
	got := beautify(t, "CREATE TABLE foo (id UInt64) ENGINE = MergeTree ORDER BY id PARTITION BY toYYYYMM(d) SETTINGS index_granularity=8192, parts_to_throw_insert=300")
	require.Contains(t, got, "ENGINE = MergeTree\nORDER BY id\nPARTITION BY toYYYYMM(d)\nSETTINGS\n  index_granularity=8192,\n  parts_to_throw_insert=300")
}

// Subselects in FROM are beautified with an extra indent level.
func TestBeautifyVisitor_FromSubquery(t *testing.T) {
	got := beautify(t, "select one from (select main,sum(two) as one from tabl where x>1) where main like '%olsztyn%' order by one desc limit 10")
	want := "SELECT\n  one\nFROM\n  (\n    SELECT\n      main,\n      sum(two) AS one\n    FROM\n      tabl\n    WHERE\n      x > 1\n  )\nWHERE\n  main LIKE '%olsztyn%'\nORDER BY\n  one DESC\nLIMIT 10"
	require.Equal(t, want, got)
}

func TestBeautifyVisitor_FromSubqueryWithAlias(t *testing.T) {
	got := beautify(t, "SELECT * FROM (SELECT a FROM t) AS sub WHERE sub.a < 10")
	require.Contains(t, got, "FROM\n  (\n    SELECT\n      a\n    FROM\n      t\n  ) AS sub")
}

func TestBeautifyVisitor_FromSubqueryNested(t *testing.T) {
	got := beautify(t, "SELECT * FROM (SELECT * FROM (SELECT x FROM inner_t) AS a) AS b")
	// Outer subquery indents 1 level, inner indents 2 levels (4 spaces relative to outer SELECT).
	require.Contains(t, got, "FROM\n  (\n    SELECT\n      *\n    FROM\n      (\n        SELECT\n          x\n        FROM\n          inner_t\n      ) AS a\n  ) AS b")
}
