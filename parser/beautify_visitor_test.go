package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"
)

// TestParser_Beautify runs every input SQL fixture under
// parser/testdata/{basic,ddl,dml,query}/*.sql through BeautifyVisitor and
// compares it to the golden under that directory's `format/beautify/`
// subdirectory. These goldens were initially seeded from origin/master and
// re-pinned against our output; see refactor-visitor.md for the rationale.
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

// TestBeautifyVisitor_ProjectionSelect covers issue #19: BeautifyVisitor must
// render a projection's inner SELECT (ProjectionSelectStmt) and its ORDER BY
// (ProjectionOrderByClause) directly. Before the VisitProjectionSelect /
// VisitProjectionOrderBy overrides existed, Accept fell through to
// DefaultASTVisitor and emitted nothing meaningful.
func TestBeautifyVisitor_ProjectionSelect(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		want string
	}{
		{
			name: "select_star_order_by",
			sql:  "CREATE TABLE t (id UInt64, PROJECTION p (SELECT * ORDER BY id)) ENGINE = MergeTree ORDER BY id",
			want: "(\n  SELECT *\n  ORDER BY id\n)",
		},
		{
			name: "group_by",
			sql:  "CREATE TABLE t (user_id UInt64, user_agent String, pages UInt64, PROJECTION p (SELECT user_agent, sum(pages) GROUP BY user_id, user_agent)) ENGINE = MergeTree ORDER BY user_agent",
			want: "(\n  SELECT user_agent, sum(pages)\n  GROUP BY\n    user_id, user_agent\n)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts, err := NewParser(tt.sql).ParseStmts()
			require.NoError(t, err)
			ct := stmts[0].(*CreateTable)
			var proj *TableProjection
			for _, c := range ct.TableSchema.Columns {
				if p, ok := c.(*TableProjection); ok {
					proj = p
				}
			}
			require.NotNil(t, proj)
			v := NewBeautifyVisitor()
			require.NoError(t, proj.Select.Accept(v))
			require.Equal(t, tt.want, v.String())
		})
	}
}

// TestBeautifyVisitor_Fixtures is a fixture-driven runner for the
// hand-curated beautify cases in parser/testdata/beautify/. Each test is a
// pair of files:
//
//	<name>.in.sql        — input SQL fed to the parser
//	<name>.expected.sql  — exact beautified output expected from BeautifyVisitor
//
// Adding a new test case is just dropping two files into that directory.
// To re-pin the expected output after an intentional beautify change, run:
//
//	go test ./parser -run TestBeautifyVisitor_Fixtures -update-beautify
func TestBeautifyVisitor_Fixtures(t *testing.T) {
	const fixtureDir = "./testdata/beautify"
	entries, err := os.ReadDir(fixtureDir)
	require.NoError(t, err)
	cases := 0
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".in.sql") {
			continue
		}
		caseName := strings.TrimSuffix(name, ".in.sql")
		cases++
		t.Run(caseName, func(t *testing.T) {
			inPath := filepath.Join(fixtureDir, name)
			expectedPath := filepath.Join(fixtureDir, caseName+".expected.sql")

			inputBytes, err := os.ReadFile(inPath)
			require.NoError(t, err)
			input := strings.TrimRight(string(inputBytes), "\n")

			stmts, err := NewParser(input).ParseStmts()
			require.NoError(t, err)

			var got strings.Builder
			for i, stmt := range stmts {
				if i > 0 {
					got.WriteString(";\n")
				}
				v := NewBeautifyVisitor()
				require.NoError(t, stmt.Accept(v))
				got.WriteString(v.String())
			}
			gotStr := got.String()
			// Pin one trailing newline to keep the fixture files
			// well-formed under standard editors / pre-commit hooks.
			fileText := gotStr
			if !strings.HasSuffix(fileText, "\n") {
				fileText += "\n"
			}

			if updateBeautifyFixtures {
				require.NoError(t, os.WriteFile(expectedPath, []byte(fileText), 0o644))
				return
			}
			expectedBytes, err := os.ReadFile(expectedPath)
			require.NoError(t, err, "missing fixture %s", expectedPath)
			// Compare with trailing-newline normalization — the fixture
			// may or may not have one; what matters is the SQL content.
			require.Equal(t,
				strings.TrimRight(string(expectedBytes), "\n"),
				strings.TrimRight(gotStr, "\n"),
				"beautify output for %s did not match %s — rerun with -update-beautify to re-pin",
				inPath, expectedPath)
		})
	}
	require.Greater(t, cases, 0, "no .in.sql fixtures found in %s", fixtureDir)
}
