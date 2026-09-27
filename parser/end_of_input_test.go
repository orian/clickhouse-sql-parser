package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_StringReparsesAtEndOfInput re-parses every fixture statement's
// String() on its own, with no trailing ';', and checks it prints the same.
// A final single-token clause such as `ALTER TABLE t FREEZE` used to be
// skipped at end of input (#46).
func TestParser_StringReparsesAtEndOfInput(t *testing.T) {
	for _, dir := range sqlFixtureDirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
		require.NoError(t, err)
		for _, file := range files {
			src, err := os.ReadFile(file)
			require.NoError(t, err)
			stmts, err := NewParser(string(src)).ParseStmts()
			if err != nil {
				continue
			}
			for _, stmt := range stmts {
				out := stmt.String()
				reparsed, err := NewParser(out).ParseStmts()
				require.NoError(t, err, "%s: %s", file, out)
				require.Len(t, reparsed, 1, "%s: %s", file, out)
				require.Equal(t, out, reparsed[0].String(), file)
			}
		}
	}
}

// TestParser_EndOfInput pins statements whose last token used to be dropped
// at end of input (#46), and inputs that used to parse to zero statements or
// an empty select list instead of an error (#82).
func TestParser_EndOfInput(t *testing.T) {
	for _, sql := range []string{
		"ALTER TABLE t FREEZE",
		"ALTER TABLE t FREEZE, FREEZE",
		"SELECT limit",
		"SELECT from FROM t",
		"SELECT a, FROM t",
		"SELECT format('{} {}', a, b) FROM t",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
		})
	}
	for _, sql := range []string{
		"SELECT",
		"SELECT --1",
		"SELECT FROM t",
		"SELECT DISTINCT",
		"SELECT DISTINCT FROM t",
		"WITH 1 AS x SELECT",
		"EXPLAIN",
		"SELECT 1; SELECT",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
