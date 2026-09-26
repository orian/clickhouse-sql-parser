package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_InsertIntoTableFunction verifies that `INSERT INTO [TABLE]
// FUNCTION` keeps FUNCTION (#53). Without it ClickHouse reads `file(x)` as an
// insert into a table named file with column list (x), or rejects the query.
func TestParser_InsertIntoTableFunction(t *testing.T) {
	for _, sql := range []string{
		"INSERT INTO FUNCTION file('out.parquet') SELECT * FROM numbers(10)",
		"INSERT INTO TABLE FUNCTION file('/dev/null', 'Parquet', 'number UInt64') SELECT * FROM numbers(10)",
		"INSERT INTO FUNCTION null() SELECT 1",
		"INSERT INTO FUNCTION remote('127.1', currentDatabase(), t) SELECT toUInt64(number) FROM system.numbers LIMIT 1",
		"INSERT INTO FUNCTION s3(s3_conn, url = 'http://localhost:11111/test/x.parquet', format = Parquet) SELECT number AS id FROM numbers(1)",
		"INSERT INTO FUNCTION mysql('localhost:3306', 'test', 'test', 'bayonet', '123', 1) (int_id, float) VALUES (1, 3)",
		"INSERT INTO file (x) VALUES (1)",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, sql, stmts[0].String())

			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())

			beautifyRoundTrip(t, stmts[0])
		})
	}

	stmts, err := NewParser("INSERT INTO FUNCTION mysql('h', 'db', 't', 'u', 'p') (a, b) VALUES (1, 2)").ParseStmts()
	require.NoError(t, err)
	insert := stmts[0].(*InsertStmt)
	require.True(t, insert.IsTableFunction())
	require.Nil(t, insert.Table.(*FunctionExpr).Params.ColumnArgList)
	require.Equal(t, "(a, b)", insert.ColumnNames.String())
}
