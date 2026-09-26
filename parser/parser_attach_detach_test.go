package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_AttachDetachRoundTrip verifies that ATTACH and DETACH keep their
// verb instead of printing as CREATE and DROP (#43, #52). Every statement here
// is accepted by ClickHouse 26.8 and formats back to itself.
func TestParser_AttachDetachRoundTrip(t *testing.T) {
	for _, sql := range []string{
		"ATTACH TABLE t (x UInt8) ENGINE = Memory",
		"ATTACH TABLE t",
		"ATTACH TABLE IF NOT EXISTS t",
		"ATTACH TEMPORARY TABLE t (x UInt8) ENGINE = Memory",
		"ATTACH VIEW v AS SELECT 1",
		"ATTACH MATERIALIZED VIEW v TO t AS SELECT 1",
		"ATTACH DICTIONARY d",
		"ATTACH DICTIONARY IF NOT EXISTS db.d ON CLUSTER c",
		"ATTACH DATABASE db",
		"ATTACH DATABASE db ENGINE = Atomic",
		"DETACH TABLE t",
		"DETACH TABLE IF EXISTS t ON CLUSTER c PERMANENTLY SYNC",
		"DETACH VIEW v",
		"DETACH DICTIONARY d",
		"DETACH DATABASE db",
		"DETACH DATABASE IF EXISTS db PERMANENTLY",
		"DETACH DATABASE db PERMANENTLY SYNC",
		"DROP DATABASE IF EXISTS db ON CLUSTER c SYNC",
		"DETACH TEMPORARY TABLE t",
		"DETACH TABLE t NO DELAY",
		"DROP TABLE IF EXISTS t ON CLUSTER c SYNC",
		"CREATE DATABASE db ENGINE = Atomic",
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

// TestParser_AttachDetachRejected lists forms ClickHouse 26.8 rejects.
func TestParser_AttachDetachRejected(t *testing.T) {
	for _, sql := range []string{
		"ATTACH OR REPLACE TABLE t (x UInt8) ENGINE = Memory",
		"ATTACH FUNCTION f AS (x) -> x",
		"ATTACH ROLE r",
		"ATTACH USER u",
		"ATTACH NAMED COLLECTION nc AS a = 1",
		"DETACH USER u",
		"DETACH ROLE r",
		"DETACH FUNCTION f",
		"DROP TABLE t PERMANENTLY",
		"DROP DATABASE db PERMANENTLY",
		"CREATE DICTIONARY d",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

// TestParser_AttachDetachFlags checks the AST flags and positions.
func TestParser_AttachDetachFlags(t *testing.T) {
	sql := "DETACH TABLE IF EXISTS t ON CLUSTER c PERMANENTLY SYNC"
	stmts, err := NewParser(sql).ParseStmts()
	require.NoError(t, err)
	drop := stmts[0].(*DropStmt)
	require.True(t, drop.IsDetach)
	require.True(t, drop.Permanently)
	require.Equal(t, "DETACH TABLE", drop.Type())

	sql = "DETACH DATABASE db PERMANENTLY"
	stmts, err = NewParser(sql).ParseStmts()
	require.NoError(t, err)
	dropDB := stmts[0].(*DropDatabase)
	require.True(t, dropDB.IsDetach)
	require.True(t, dropDB.Permanently)
	require.Equal(t, Pos(len(sql)), dropDB.End())

	sql = "ATTACH DICTIONARY db.d"
	stmts, err = NewParser(sql).ParseStmts()
	require.NoError(t, err)
	dict := stmts[0].(*CreateDictionary)
	require.True(t, dict.IsAttach)
	require.Nil(t, dict.Schema)
	require.Equal(t, Pos(len(sql)), dict.End())

	// `permanently` is not reserved and still works as an identifier.
	_, err = NewParser("SELECT permanently FROM t").ParseStmts()
	require.NoError(t, err)
}
