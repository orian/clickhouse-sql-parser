package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// These fixtures are byte-for-byte copies of the documentation panic examples.
func TestParser_CreateUserDocumentationPanics(t *testing.T) {
	files, err := filepath.Glob("testdata/regressions/create_user/*.sql")
	require.NoError(t, err)
	require.Len(t, files, 4)
	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			sql, err := os.ReadFile(file)
			require.NoError(t, err)
			assertUserStatement(t, string(sql))
		})
	}
}

func assertUserStatement(t *testing.T, sql string) {
	t.Helper()
	var stmts []Expr
	var err error
	require.NotPanics(t, func() { stmts, err = NewParser(sql).ParseStmts() })
	require.NoError(t, err)
	require.Len(t, stmts, 1)
	require.Equal(t, Pos(len(strings.TrimRight(sql, " \t\r\n;"))), stmts[0].End())
	// Check that source positions and formatting are usable after successful parsing.
	Walk(stmts[0], func(expr Expr) bool {
		require.LessOrEqual(t, expr.Pos(), expr.End())
		require.LessOrEqual(t, expr.End(), Pos(len(sql)))
		return true
	})
	formatted := stmts[0].String()
	reparsed, err := NewParser(formatted).ParseStmts()
	require.NoError(t, err)
	require.Len(t, reparsed, 1)
	require.Equal(t, formatted, reparsed[0].String())
}

func TestParser_CreateUserClauseEnds(t *testing.T) {
	for _, clause := range []string{
		"NOT IDENTIFIED", "IDENTIFIED WITH kerberos", "IDENTIFIED WITH no_password",
		"HOST ANY", "HOST LOCAL", "HOST NONE", "DEFAULT ROLE NONE",
		"GRANTEES ANY", "GRANTEES NONE", "DEFAULT DATABASE NONE",
		"SETTINGS PROFILE p",
	} {
		for _, ending := range []string{"", ";", "\n"} {
			t.Run(clause+"/"+ending, func(t *testing.T) {
				assertUserStatement(t, "CREATE USER u "+clause+ending)
			})
		}
	}
}

func TestParser_CreateRoleProfileWithoutValue(t *testing.T) {
	assertUserStatement(t, "CREATE ROLE r SETTINGS PROFILE p")
}
