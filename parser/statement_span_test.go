package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

// skipSpaceAndComments returns the offset after any whitespace and SQL
// comments starting at i.
func skipSpaceAndComments(src string, i int) int {
	for i < len(src) {
		r, size := utf8.DecodeRuneInString(src[i:])
		switch {
		case unicode.IsSpace(r):
			i += size
		case strings.HasPrefix(src[i:], "--"):
			if j := strings.IndexByte(src[i:], '\n'); j >= 0 {
				i += j + 1
			} else {
				i = len(src)
			}
		case strings.HasPrefix(src[i:], "/*"):
			if j := strings.Index(src[i+2:], "*/"); j >= 0 {
				i += j + 4
			} else {
				i = len(src)
			}
		default:
			return i
		}
	}
	return i
}

// checkStatementSpan verifies that stmt's [Pos, End) covers exactly the
// statement in src: it starts at a non-space character, its last character is
// neither whitespace nor ';', and only whitespace and comments separate End
// from the terminating ';' or the end of input.
func checkStatementSpan(src string, stmt Expr) error {
	pos, end := int(stmt.Pos()), int(stmt.End())
	if pos < 0 || end > len(src) || pos >= end {
		return fmt.Errorf("span [%d, %d) out of range (input length %d)", pos, end, len(src))
	}
	if r, _ := utf8.DecodeRuneInString(src[pos:]); unicode.IsSpace(r) || r == ';' {
		return fmt.Errorf("Pos() is at %q", r)
	}
	if r, _ := utf8.DecodeLastRuneInString(src[:end]); unicode.IsSpace(r) || r == ';' {
		return fmt.Errorf("End() includes trailing %q", r)
	}
	if next := skipSpaceAndComments(src, end); next < len(src) && src[next] != ';' {
		tail := src[end:]
		if len(tail) > 40 {
			tail = tail[:40]
		}
		return fmt.Errorf("End() stops before %q", tail)
	}
	return nil
}

// TestStatementSpans checks Pos()/End() of every statement in the SQL
// fixtures (#65). Tools slice the source with these positions, so a statement
// that ends early (e.g. before a closing quote, `)`, DESC or FINAL) or late
// (including the `;`) silently truncates or corrupts the extracted SQL.
func TestStatementSpans(t *testing.T) {
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
					require.NoError(t, checkStatementSpan(string(src), stmt), "statement: %s", stmt.String())
				}
			})
		}
	}
}

// TestStatementSpansQuotedTokens pins spans that end in quoted tokens and
// other closing delimiters, which used to stop one character early.
func TestStatementSpansQuotedTokens(t *testing.T) {
	for _, sql := range []string{
		"SELECT 'abc'",
		"SELECT `quoted`",
		`SELECT "quoted"`,
		"SELECT * FROM 'test_table'",
		"SELECT CAST(x AS DateTime('UTC'))",
		"SELECT count() FROM t GROUP BY a ORDER BY a DESC",
		"SELECT a FROM t FINAL",
		"SELECT a FROM t WHERE a IS NOT NULL",
		"SELECT CASE WHEN 1 THEN 2 END",
		"SELECT quantile(0.5)(x) FROM t",
		"SELECT 1 UNION ALL SELECT 2",
		"SELECT a FROM t JOIN u USING (a)",
		"SELECT * FROM t0 WHERE id = ?",
		"SHOW DATABASES",
		"OPTIMIZE TABLE t FINAL",
		"OPTIMIZE TABLE t DEDUPLICATE BY * EXCEPT c",
		"DROP TABLE t SYNC",
		"CREATE TABLE t AS other",
		"CREATE TABLE t (x Enum('a' = 1)) ENGINE = Memory",
		"CREATE TABLE t (x UInt64 CODEC(LZ4HC(1))) ENGINE = Memory",
		"CREATE DATABASE db COMMENT 'c'",
		"CREATE NAMED COLLECTION nc AS a = '1' OVERRIDABLE",
		"CREATE VIEW v AS (SELECT 1) COMMENT 'c'",
		"GRANT SELECT ON db.t TO u WITH GRANT OPTION",
		"ALTER TABLE t MODIFY COLUMN c REMOVE COMMENT",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql + ";").ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.NoError(t, checkStatementSpan(sql+";", stmts[0]))
			require.Equal(t, Pos(len(sql)), stmts[0].End())
		})
	}
}
