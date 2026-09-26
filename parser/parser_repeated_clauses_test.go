package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_CreateTableQuerySettings(t *testing.T) {
	for _, tc := range []struct {
		sql     string
		storage string
		query   string
	}{
		{
			sql:     "CREATE TABLE d.t (id UInt64) ENGINE = Kafka SETTINGS kafka_broker_list = 'b', kafka_topic_list = 't'",
			storage: "SETTINGS kafka_broker_list='b', kafka_topic_list='t'",
		},
		{
			sql:     "CREATE TABLE d.t (id UInt64) ENGINE = Kafka SETTINGS kafka_broker_list = 'b', kafka_topic_list = 't' SETTINGS flatten_nested = 0",
			storage: "SETTINGS kafka_broker_list='b', kafka_topic_list='t'",
			query:   "SETTINGS flatten_nested=0",
		},
		{
			sql:     "CREATE TABLE d.t (id UInt64) ENGINE = MergeTree ORDER BY id SETTINGS index_granularity = 8192 SETTINGS flatten_nested = 0",
			storage: "SETTINGS index_granularity=8192",
			query:   "SETTINGS flatten_nested=0",
		},
		{
			sql:   "CREATE TABLE d.t (id UInt64) ENGINE = Memory COMMENT 'c' SETTINGS max_threads = 4",
			query: "SETTINGS max_threads=4",
		},
	} {
		t.Run(tc.sql, func(t *testing.T) {
			stmts, err := NewParser(tc.sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			ct := stmts[0].(*CreateTable)

			if tc.storage == "" {
				require.Nil(t, ct.Engine.Settings)
			} else {
				require.NotNil(t, ct.Engine.Settings)
				require.Equal(t, tc.storage, ct.Engine.Settings.String())
			}
			if tc.query == "" {
				require.Nil(t, ct.Settings)
			} else {
				require.NotNil(t, ct.Settings)
				require.Equal(t, tc.query, ct.Settings.String())
				require.Equal(t, ct.Settings.End(), ct.End())
			}

			// Both clauses survive a round trip in source order.
			reparsed, err := NewParser(ct.String()).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, ct.String(), reparsed[0].String())

			beautify := NewBeautifyVisitor()
			require.NoError(t, ct.Accept(beautify))
			reparsed, err = NewParser(beautify.String()).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, ct.String(), reparsed[0].String())

			if tc.query != "" {
				found := false
				Walk(ct, func(expr Expr) bool {
					if expr == ct.Settings {
						found = true
					}
					return true
				})
				require.True(t, found, "Walk should visit CreateTable.Settings")
			}
		})
	}
}

func TestParser_CreateTableRepeatedClauses(t *testing.T) {
	for _, sql := range []string{
		"CREATE TABLE d.t (id UInt64) ENGINE = MergeTree ORDER BY id SETTINGS index_granularity = 8192 SETTINGS flatten_nested = 0 SETTINGS max_threads = 4",
		"CREATE TABLE d.t (id UInt64) ENGINE = MergeTree ORDER BY id ORDER BY (id, id)",
		"CREATE TABLE d.t (id UInt64) ENGINE = MergeTree PARTITION BY id PARTITION BY id ORDER BY id",
		"CREATE TABLE d.t (id UInt64) ENGINE = MergeTree ORDER BY id PRIMARY KEY id PRIMARY KEY id",
		"CREATE TABLE d.t (id UInt64) ENGINE = MergeTree ORDER BY id SAMPLE BY id SAMPLE BY id",
		"CREATE TABLE d.t (d DateTime) ENGINE = MergeTree ORDER BY d TTL d + INTERVAL 1 DAY TTL d + INTERVAL 2 DAY",
		// Engine clauses cannot follow the query-level SETTINGS clause.
		"CREATE TABLE d.t (id UInt64) ENGINE = MergeTree SETTINGS index_granularity = 8192 SETTINGS flatten_nested = 0 ORDER BY id",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

func TestParser_CreateUserRepeatedClauses(t *testing.T) {
	stmts, err := NewParser("CREATE USER u HOST IP '10.0.0.1' HOST LOCAL SETTINGS max_threads = 1 SETTINGS readonly = 1").ParseStmts()
	require.NoError(t, err)
	user := stmts[0].(*CreateUser)
	// HOST and SETTINGS clauses merge, as in ClickHouse.
	require.Len(t, user.Hosts, 2)
	require.Len(t, user.Settings, 2)
	reparsed, err := NewParser(user.String()).ParseStmts()
	require.NoError(t, err)
	require.Equal(t, user.String(), reparsed[0].String())

	for _, sql := range []string{
		"CREATE USER u IDENTIFIED BY 'a' IDENTIFIED BY 'b'",
		"CREATE USER u NOT IDENTIFIED IDENTIFIED BY 'b'",
		"CREATE USER u VALID UNTIL '2025-01-01' VALID UNTIL '2030-01-01'",
		"CREATE USER u DEFAULT ROLE r1 DEFAULT ROLE r2",
		"CREATE USER u DEFAULT DATABASE db DEFAULT DATABASE NONE",
		"CREATE USER u DEFAULT DATABASE NONE DEFAULT DATABASE db",
		"CREATE USER u GRANTEES a GRANTEES b",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}

func TestParser_GroupByModifiers(t *testing.T) {
	for _, sql := range []string{
		"SELECT a, count() FROM t GROUP BY a WITH ROLLUP",
		"SELECT a, count() FROM t GROUP BY a WITH CUBE WITH TOTALS",
		"SELECT a, count() FROM t GROUP BY a WITH TOTALS",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
		})
	}
	for _, sql := range []string{
		"SELECT a, count() FROM t GROUP BY a WITH ROLLUP WITH CUBE",
		"SELECT a, count() FROM t GROUP BY a WITH ROLLUP WITH ROLLUP",
		"SELECT a, count() FROM t GROUP BY a WITH TOTALS WITH ROLLUP",
		"SELECT a, count() FROM t GROUP BY a WITH TOTALS WITH TOTALS",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
