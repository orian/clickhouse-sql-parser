package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_RejectsInvalidDDL pins DDL that ClickHouse 26.8 rejects and the
// parser used to accept (#50), next to the valid forms that still parse.
func TestParser_RejectsInvalidDDL(t *testing.T) {
	for _, sql := range []string{
		// OR REPLACE and IF NOT EXISTS are exclusive (except for DICTIONARY).
		"CREATE OR REPLACE TABLE IF NOT EXISTS t (a Int32) ENGINE = Memory",
		"CREATE OR REPLACE VIEW IF NOT EXISTS v AS SELECT 1",
		// RENAME DATABASE takes one pair.
		"RENAME DATABASE a TO b, c TO d",
		// A partition is a literal, a tuple or a query parameter.
		"ALTER TABLE t DROP PARTITION p",
		"ALTER TABLE t DROP PARTITION (p)",
		"ALTER TABLE t DROP PARTITION toDate('2020-01-01')",
		"ALTER TABLE t DROP PARTITION 1 + 1",
		"ALTER TABLE t DROP PARTITION [p]",
		"ALTER TABLE t DROP PARTITION CAST(p, 'UInt32')",
		"ALTER TABLE t DROP PARTITION p::UInt32",
		"ALTER TABLE t DROP PARTITION toUInt32(1)",
		"ALTER TABLE t CLEAR COLUMN c IN PARTITION partition_name",
		"OPTIMIZE TABLE t PARTITION p",
		// A UUID is 32 hex digits, optionally 8-4-4-4-12.
		"CREATE TABLE t UUID '1234' (a Int32) ENGINE = Memory",
		"CREATE TABLE t UUID '{12345678-1234-1234-1234-123456789012}' (a Int32) ENGINE = Memory",
		"CREATE DATABASE d UUID ''",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
	for _, sql := range []string{
		"CREATE OR REPLACE DICTIONARY IF NOT EXISTS d (k UInt64) PRIMARY KEY k SOURCE(NULL()) LAYOUT(FLAT()) LIFETIME(0)",
		"CREATE OR REPLACE TABLE t (a Int32) ENGINE = Memory",
		"RENAME DICTIONARY a TO b, c TO d",
		"ALTER TABLE t DROP PARTITION 201901",
		"ALTER TABLE t DROP PARTITION '2019-01-01'",
		"ALTER TABLE t DROP PARTITION -1",
		"ALTER TABLE t DROP PARTITION (1)",
		"ALTER TABLE t DROP PARTITION (1, p)",
		"ALTER TABLE t DROP PARTITION (toDate('2020-01-01'), 1)",
		"ALTER TABLE t DROP PARTITION tuple(1, 'a')",
		"ALTER TABLE t DROP PARTITION [1, 2]",
		"ALTER TABLE t DROP PARTITION NULL",
		"ALTER TABLE t DROP PARTITION true",
		"ALTER TABLE t DROP PARTITION CAST(20260624, 'UInt32')",
		"ALTER TABLE t DROP PARTITION _CAST(20260624, 'UInt32')",
		"ALTER TABLE t DROP PARTITION CAST((1, 2) AS Tuple(UInt8, UInt8))",
		"ALTER TABLE t DROP PARTITION 20260624::UInt32",
		"ALTER TABLE t UPDATE flag = false IN PARTITION _CAST(20260624, 'UInt32') WHERE 1",
		"ALTER TABLE t DROP PARTITION [true, NULL]",
		"ALTER TABLE t DROP PARTITION {p:String}",
		"ALTER TABLE t DROP PARTITION ID 'x'",
		"ALTER TABLE t DROP PARTITION ALL",
		"CREATE TABLE t UUID '12345678-1234-1234-1234-123456789012' (a Int32) ENGINE = Memory",
		"CREATE TABLE t UUID '12345678123412341234123456789ABC' (a Int32) ENGINE = Memory",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
		})
	}
}
