package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_QueryOutputClauses covers ClickHouse's query output clauses
// `[FORMAT fmt] [SETTINGS ...]` (#41). FORMAT used to be parsed and silently
// dropped for every statement except SELECT and INSERT, and a SETTINGS clause
// after FORMAT was rejected. Every statement here is accepted by ClickHouse
// 26.8 and formats back to the same query.
func TestParser_QueryOutputClauses(t *testing.T) {
	for _, sql := range []string{
		"SHOW TABLES FORMAT JSON",
		"SHOW TABLES FORMAT JSON SETTINGS max_threads=1",
		"SHOW TABLES SETTINGS max_threads=1",
		"SHOW CREATE TABLE t FORMAT TSVRaw",
		"SHOW DATABASES FORMAT JSON",
		"SHOW DATABASES LIKE 'prod%' LIMIT 5 INTO OUTFILE '/tmp/prod_dbs.txt' FORMAT JSON",
		"DESCRIBE TABLE t FORMAT JSON SETTINGS max_threads=1",
		"CHECK TABLE t FORMAT JSON",
		"DROP TABLE IF EXISTS t ON CLUSTER c FORMAT Null",
		"DROP TABLE t SETTINGS max_threads=1",
		"DROP DATABASE db FORMAT Null",
		"DETACH TABLE t FORMAT Null",
		"ATTACH TABLE t FORMAT Null",
		"TRUNCATE TABLE t SETTINGS max_threads=1",
		"RENAME TABLE a TO b FORMAT Null",
		"OPTIMIZE TABLE t FORMAT Null",
		"ALTER TABLE t DELETE WHERE 1 SETTINGS mutations_sync=2",
		"CREATE DATABASE db FORMAT Null",
		"CREATE TABLE t (x UInt8) ENGINE = Memory FORMAT JSON",
		"CREATE TABLE t (x UInt8) ENGINE = Memory FORMAT JSON SETTINGS max_threads=1",
		"CREATE TABLE t (x UInt8) ENGINE = Memory SETTINGS max_threads=1 FORMAT JSON SETTINGS max_block_size=1",
		"CREATE DICTIONARY d (x UInt64) PRIMARY KEY x SOURCE(NULL()) LIFETIME(0) LAYOUT(FLAT()) FORMAT Null",
		"CREATE VIEW v AS SELECT 1 FORMAT Null",
		"EXPLAIN AST SELECT 1 FORMAT JSON SETTINGS max_threads=1",
		"SELECT 1 FORMAT JSON SETTINGS max_block_size=10",
		"SELECT 1 SETTINGS max_threads=1 FORMAT JSON SETTINGS max_block_size=10",
		"SELECT 1 UNION ALL SELECT 2 FORMAT JSON SETTINGS max_block_size=10",
		"INSERT INTO t FORMAT JSONEachRow",
	} {
		t.Run(sql, func(t *testing.T) {
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.Equal(t, sql, stmts[0].String())

			printer := NewPrintVisitor()
			require.NoError(t, stmts[0].Accept(printer))
			require.Equal(t, sql, printer.String())

			beautifyRoundTrip(t, stmts[0])
		})
	}
}

// TestParser_QueryOutputClausesRejected lists statements ClickHouse 26.8
// rejects with a syntax error. They used to parse with the FORMAT silently
// dropped.
func TestParser_QueryOutputClausesRejected(t *testing.T) {
	for _, sql := range []string{
		"SYSTEM FLUSH LOGS FORMAT Null",
		"SYSTEM FLUSH LOGS SETTINGS max_threads = 1",
		"USE db FORMAT Null",
		"SET max_threads = 1 FORMAT Null",
		"GRANT SELECT ON db.t TO u FORMAT Null",
		"GRANT SELECT ON db.t TO u SETTINGS max_threads = 1",
		"CREATE USER u FORMAT Null",
		"CREATE ROLE r FORMAT Null",
		"ALTER ROLE r FORMAT Null",
		"DROP USER u FORMAT Null",
		"DROP ROLE r FORMAT Null",
		"DELETE FROM t WHERE 1 FORMAT Null",
		"CREATE FUNCTION f AS (x) -> x FORMAT Null",
		"CREATE NAMED COLLECTION nc AS a = 1 FORMAT Null",
		"SELECT 1 FORMAT JSON SETTINGS a = 1 SETTINGS b = 2",
		"SHOW DATABASES FORMAT 'TabSeparated'",
		// CREATE TABLE's query-level SETTINGS may be written once, before or
		// after FORMAT.
		"CREATE TABLE t (x UInt8) ENGINE = MergeTree ORDER BY x SETTINGS a = 1 SETTINGS b = 2 FORMAT JSON SETTINGS c = 3",
		// After VALUES everything is data: ClickHouse fails to parse
		// "FORMAT Native" as a row (CANNOT_PARSE_INPUT_ASSERTION_FAILED).
		"INSERT INTO t (a, b) VALUES (1, 2) FORMAT Native",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

// TestParser_QueryOutputAST checks the AST shape: only statements with output
// clauses are wrapped, and SELECT keeps its two SETTINGS clauses apart
// (ClickHouse applies the SELECT's own value on a key clash).
func TestParser_QueryOutputAST(t *testing.T) {
	stmts, err := NewParser("SHOW TABLES").ParseStmts()
	require.NoError(t, err)
	require.IsType(t, &ShowStmt{}, stmts[0])

	sql := "DESCRIBE TABLE t FORMAT JSON SETTINGS max_threads=1"
	stmts, err = NewParser(sql).ParseStmts()
	require.NoError(t, err)
	wrapped := stmts[0].(*QueryWithOutput)
	require.IsType(t, &DescribeStmt{}, wrapped.Query)
	require.Equal(t, "FORMAT JSON", wrapped.Format.String())
	require.Equal(t, "SETTINGS max_threads=1", wrapped.Settings.String())
	require.Equal(t, Pos(len(sql)), wrapped.End())

	stmts, err = NewParser("SELECT 1 SETTINGS max_threads=1 FORMAT JSON SETTINGS max_threads=3").ParseStmts()
	require.NoError(t, err)
	sel := stmts[0].(*SelectQuery)
	require.Equal(t, "SETTINGS max_threads=1", sel.Settings.String())
	require.Equal(t, "SETTINGS max_threads=3", sel.OutputSettings.String())
}
