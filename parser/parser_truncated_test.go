package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_TruncatedSQL(t *testing.T) {
	for _, sql := range []string{"ALTER ", "DROP ", "SHOW ", "CREATE OR REPLACE ", "CREATE--", "/*", "/* unfinished", "SELECT 1 /* unfinished", "CREATE USER A SETTINGS", "CREATE ROLE r SETTINGS", "CREATE USER u SETTINGS x = 1,", "SELECT 1 ORDER BY 1 AS"} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

func TestLexer_CommentAtEOF(t *testing.T) {
	for _, sql := range []string{"--", "-- comment", "-- comment\n", "/* comment */"} {
		t.Run(sql, func(t *testing.T) {
			lexer := NewLexer(sql)
			require.NoError(t, lexer.consumeToken())
			require.Equal(t, len(sql), lexer.current)
		})
	}
}
