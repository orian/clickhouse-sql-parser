package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// beautifyRoundTrip beautifies stmt and checks that the output parses back to
// a statement with the same String() form, i.e. that no clause was lost.
func beautifyRoundTrip(t *testing.T, stmt Expr) string {
	t.Helper()
	beautify := NewBeautifyVisitor()
	require.NoError(t, stmt.Accept(beautify))
	out := beautify.String()
	// No trailing ';': a statement must parse at end of input (#46).
	reparsed, err := NewParser(out).ParseStmts()
	require.NoError(t, err, "beautified output does not parse:\n%s", out)
	require.Len(t, reparsed, 1)
	require.Equal(t, stmt.String(), reparsed[0].String(), "beautified output:\n%s", out)
	return out
}

// TestBeautify_AllFixtureStatements beautifies every statement in the SQL
// fixtures. Statement types without a BeautifyVisitor override used to fall
// back to DefaultASTVisitor and emit nothing (#44); this catches any statement
// type that is added without beautifier support.
func TestBeautify_AllFixtureStatements(t *testing.T) {
	for _, dir := range sqlFixtureDirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
		require.NoError(t, err)
		for _, file := range files {
			t.Run(file, func(t *testing.T) {
				src, err := os.ReadFile(file)
				require.NoError(t, err)
				stmts, err := NewParser(string(src)).ParseStmts()
				if err != nil {
					t.Skip("fixture does not parse")
				}
				for _, stmt := range stmts {
					beautifyRoundTrip(t, stmt)
				}
			})
		}
	}
}

func TestBeautify_StatementsWithoutFixtures(t *testing.T) {
	for _, tc := range []struct {
		sql  string
		want string
	}{
		{
			sql:  "EXPLAIN AST SELECT a FROM t",
			want: "EXPLAIN AST\nSELECT a\nFROM t",
		},
		{
			sql:  "CREATE LIVE VIEW v AS SELECT 1",
			want: "CREATE LIVE VIEW v\nAS\nSELECT 1",
		},
		{
			sql:  "CREATE LIVE VIEW IF NOT EXISTS v UUID '3493e374-e2bb-481b-b493-e374e2bb981b' WITH TIMEOUT 10 TO dst (id UInt64) AS SELECT id FROM t",
			want: "CREATE LIVE VIEW IF NOT EXISTS v UUID '3493e374-e2bb-481b-b493-e374e2bb981b'\nWITH TIMEOUT 10\nTO dst (\n  id UInt64\n)\nAS\nSELECT id\nFROM t",
		},
		{
			sql:  "ATTACH DICTIONARY d (x UInt64) PRIMARY KEY x SOURCE(NULL()) LAYOUT(FLAT()) LIFETIME(0)",
			want: "ATTACH DICTIONARY d (x UInt64) PRIMARY KEY x SOURCE(NULL()) LIFETIME(0) LAYOUT(FLAT())",
		},
		{
			sql:  "CREATE VIEW v AS (SELECT 1)",
			want: "CREATE VIEW v\nAS\n(\n  SELECT 1\n)",
		},
	} {
		t.Run(tc.sql, func(t *testing.T) {
			stmts, err := NewParser(tc.sql).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, tc.want, beautifyRoundTrip(t, stmts[0]))
		})
	}
}
