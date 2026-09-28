package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_ColumnTransformers pins the column transformer grammar (#126):
// the parenthesis-free forms and STRICT parse, and what ClickHouse rejects
// is rejected.
func TestParser_ColumnTransformers(t *testing.T) {
	stmts, err := NewParser("SELECT * EXCEPT STRICT a REPLACE b + 1 AS b APPLY toString FROM t").ParseStmts()
	require.NoError(t, err)
	modifiers := stmts[0].(*SelectQuery).SelectItems[0].Modifiers
	require.Len(t, modifiers, 3)
	require.Equal(t, "EXCEPT", modifiers[0].Kind)
	require.True(t, modifiers[0].Strict)
	require.False(t, modifiers[0].HasParen)
	require.Equal(t, "REPLACE", modifiers[1].Kind)
	require.Equal(t, "APPLY", modifiers[2].Kind)

	for _, sql := range []string{
		"SELECT * EXCEPT FROM t",
		"SELECT * EXCEPT strict FROM t", // STRICT is the modifier, as in ClickHouse
		"SELECT * EXCEPT (t.a) FROM t",
		"SELECT * EXCEPT ('a', b) FROM t",
		"SELECT * REPLACE (a) FROM t",
		"SELECT * APPLY(f, g) FROM t",
		// A column matcher takes no alias.
		"SELECT * AS x FROM t",
		"SELECT t.* AS x FROM t",
		"SELECT * EXCEPT a AS x FROM t",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
