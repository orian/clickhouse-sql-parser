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
