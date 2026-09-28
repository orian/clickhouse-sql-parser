package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_RejectsWhatClickHouseRejects pins SQL that ClickHouse 26.8
// rejects with a syntax error and that the parser used to accept (#50).
func TestParser_RejectsWhatClickHouseRejects(t *testing.T) {
	for _, sql := range []string{
		// Column transformers follow only *, t.* or COLUMNS(...).
		"SELECT c0 REPLACE(c0 AS c1) FROM t0",
		"SELECT c0 APPLY(toString) FROM t0",
		"SELECT c0 EXCEPT (c1) FROM t0",
		// CUBE(...)/ROLLUP(...) already imply WITH CUBE/ROLLUP.
		"SELECT a FROM t GROUP BY CUBE(a) WITH CUBE",
		"SELECT a FROM t GROUP BY ROLLUP(a) WITH ROLLUP",
		"SELECT a FROM t GROUP BY CUBE(a) WITH ROLLUP",
		// A JOIN needs ON or USING unless it is CROSS, NATURAL, PASTE or ARRAY.
		"SELECT * FROM a JOIN b",
		"SELECT * FROM a LEFT JOIN b",
		"SELECT * FROM a JOIN b JOIN c ON a.x = c.x",
		// `{` in an expression starts a query parameter, never a map literal.
		"SELECT {'a': 1}",
		"SELECT * FROM t SETTINGS x = {'a': {'b': 1}}",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
	// The valid neighbours of the cases above still parse.
	for _, sql := range []string{
		"SELECT * REPLACE(i + 1 AS i) EXCEPT (j) APPLY(sum) FROM t",
		"SELECT t.* APPLY(toString) FROM t",
		"SELECT COLUMNS('c') REPLACE(c0 AS c1) FROM t",
		"SELECT a FROM t GROUP BY CUBE(a) WITH TOTALS",
		"SELECT a FROM t GROUP BY GROUPING SETS ((a)) WITH CUBE",
		"SELECT * FROM a CROSS JOIN b JOIN c ON true",
		"SELECT * FROM a, b JOIN c USING (x)",
		"SELECT * FROM a NATURAL JOIN b",
		"SELECT * FROM t SETTINGS additional_table_filters = {'t': 'x = 1'}",
		"SELECT {x:UInt8}",
		// A LIKE/ILIKE matcher takes transformers too.
		"SELECT * ILIKE 'foo%' EXCEPT (foo_extra) FROM t",
		"SELECT t.* LIKE 'a%' EXCEPT (ab) FROM t",
		// VALUES data is read by the Values format, which accepts maps.
		"INSERT INTO t VALUES (1, {'a': 1}, [{'l': 0.1}, {'l': 0.2}])",
		// ARRAY JOIN (any case) takes no ON/USING.
		"SELECT a FROM t array join arr AS a",
		"SELECT a FROM t left array join arr AS a, b",
		"SELECT 1 FROM l JOIN r ON l.x = r.y array join r.a",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
		})
	}
}

// TestParser_LowercaseArrayJoin checks that `array join` in lower case is an
// ARRAY JOIN like `ARRAY JOIN`, not an ordinary join with a table named by
// its first expression: the modifiers keep their source case.
func TestParser_LowercaseArrayJoin(t *testing.T) {
	for _, sql := range []string{"SELECT a FROM t ARRAY JOIN arr", "SELECT a FROM t array join arr"} {
		stmts, err := NewParser(sql).ParseStmts()
		require.NoError(t, err)
		join := stmts[0].(*SelectQuery).From.Expr.(*JoinExpr).Right.(*JoinExpr)
		_, isColumnList := join.Left.(*ColumnExprList)
		require.True(t, isColumnList, "%s: ARRAY JOIN operand is %T", sql, join.Left)
	}
}
