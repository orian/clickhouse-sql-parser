package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_TimeSeriesTargets covers the TimeSeries engine target-clause tail
// (SAMPLES/DATA, TAGS, METRICS), both the external-table form and the inline
// INNER COLUMNS (...) / INNER ENGINE form. See issue #8.
func TestParser_TimeSeriesTargets(t *testing.T) {
	cases := []struct {
		name string
		sql  string
		want string // expected round-trip via String(); empty means same as sql
	}{
		{
			name: "external DATA alias",
			sql: "CREATE TABLE db.m (`id` UUID, `timestamp` DateTime64(3), `value` Float64) " +
				"ENGINE = TimeSeries DATA db.m_data TAGS db.m_tags METRICS db.m_metrics",
		},
		{
			name: "external SAMPLES canonical",
			sql:  "CREATE TABLE db.m ENGINE = TimeSeries SAMPLES db.m_samples TAGS db.m_tags METRICS db.m_metrics",
		},
		{
			name: "bare TimeSeries",
			sql:  "CREATE TABLE db.m ENGINE = TimeSeries",
		},
		{
			// SETTINGS spacing is normalised by the EngineExpr formatter.
			name: "settings only",
			sql:  "CREATE TABLE db.m ENGINE = TimeSeries SETTINGS id_generator = 'sipHash64(metric_name, all_tags)'",
			want: "CREATE TABLE db.m ENGINE = TimeSeries SETTINGS id_generator='sipHash64(metric_name, all_tags)'",
		},
		{
			name: "inner columns and inner engine",
			sql: "CREATE TABLE db.m ENGINE = TimeSeries " +
				"SAMPLES INNER COLUMNS (`id` UUID, `timestamp` DateTime64(3), `value` Float64) " +
				"SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp)",
		},
		{
			name: "SHOW CREATE engine-only targets",
			sql: "CREATE TABLE db.m ENGINE = TimeSeries " +
				"DATA ENGINE = MergeTree ORDER BY (id, timestamp) " +
				"TAGS ENGINE = AggregatingMergeTree PRIMARY KEY metric_name ORDER BY tuple(metric_name, id) " +
				"METRICS ENGINE = ReplacingMergeTree ORDER BY metric_family_name",
			want: "CREATE TABLE db.m ENGINE = TimeSeries " +
				"DATA ENGINE = MergeTree ORDER BY (id, timestamp) " +
				"TAGS ENGINE = AggregatingMergeTree ORDER BY tuple(metric_name, id) PRIMARY KEY metric_name " +
				"METRICS ENGINE = ReplacingMergeTree ORDER BY metric_family_name",
		},
		{
			name: "RECENT SAMPLES external",
			sql:  "CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES d.s TAGS d.t METRICS d.m RECENT SAMPLES d.r",
		},
		{
			name: "RECENT SAMPLES inner columns and engine",
			sql: "CREATE TABLE d.ts ENGINE = TimeSeries " +
				"RECENT SAMPLES INNER COLUMNS (`id` UUID, `timestamp` DateTime64(3), `value` Float64) " +
				"RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600)",
		},
		{
			name: "RECENT SAMPLES engine-only shorthand",
			sql:  "CREATE TABLE d.ts ENGINE = TimeSeries RECENT SAMPLES ENGINE = MergeTree ORDER BY id",
		},
		{
			name: "INNER ENGINE without INNER COLUMNS",
			sql:  "CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES INNER ENGINE = MergeTree ORDER BY id",
		},
		{
			name: "INNER UUID for every target",
			sql: "CREATE TABLE d.ts ENGINE = TimeSeries " +
				"SAMPLES INNER UUID '11111111-1111-1111-1111-111111111111' SAMPLES INNER ENGINE = MergeTree ORDER BY id " +
				"RECENT SAMPLES INNER UUID '22222222-2222-2222-2222-222222222222' " +
				"TAGS INNER UUID '33333333-3333-3333-3333-333333333333' " +
				"METRICS INNER UUID '44444444-4444-4444-4444-444444444444'",
		},
		{
			// ClickHouse accepts the parts of one target anywhere in the tail
			// and formats them grouped per target, as String() does.
			name: "parts of a target spread over the tail",
			sql: "CREATE TABLE d.ts ENGINE = TimeSeries DATA d.s TAGS INNER ENGINE = MergeTree ORDER BY id " +
				"DATA INNER ENGINE = Memory TAGS INNER COLUMNS (`id` UUID)",
			want: "CREATE TABLE d.ts ENGINE = TimeSeries DATA d.s DATA INNER ENGINE = Memory " +
				"TAGS INNER COLUMNS (`id` UUID) TAGS INNER ENGINE = MergeTree ORDER BY id",
		},
		{
			name: "lowercase RECENT SAMPLES",
			sql:  "CREATE TABLE d.ts ENGINE = TimeSeries recent samples d.r",
		},
		{
			name: "engine-only target after map setting",
			sql: `CREATE TABLE db.m ENGINE = TimeSeries SETTINGS tags_to_columns = {'foo\'bar':'foo_bar'} ` +
				"DATA ENGINE = MergeTree ORDER BY (id, timestamp)",
			want: `CREATE TABLE db.m ENGINE = TimeSeries SETTINGS tags_to_columns={'foo\'bar': 'foo_bar'} ` +
				"DATA ENGINE = MergeTree ORDER BY (id, timestamp)",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			stmts, err := NewParser(tc.sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)

			want := tc.want
			if want == "" {
				want = tc.sql
			}
			require.Equal(t, want, stmts[0].String())

			// Re-parsing the formatted output must succeed and reproduce it.
			gen2, err := NewParser(stmts[0].String()).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, want, gen2[0].String())
		})
	}
}

func TestParser_TimeSeriesEngineOnlyTargetAST(t *testing.T) {
	stmts, err := NewParser(
		"CREATE TABLE db.m ENGINE = TimeSeries DATA ENGINE = MergeTree ORDER BY (id, timestamp)",
	).ParseStmts()
	require.NoError(t, err)
	require.Len(t, stmts, 1)

	create, ok := stmts[0].(*CreateTable)
	require.True(t, ok)
	require.Len(t, create.TimeSeriesTargets, 1)

	target := create.TimeSeriesTargets[0]
	require.Nil(t, target.External)
	require.Nil(t, target.InnerColumns)
	require.NotNil(t, target.InnerEngine)
	require.Equal(t, "MergeTree", target.InnerEngine.Name)
}

func TestParser_TimeSeriesRecentSamplesAST(t *testing.T) {
	sql := "CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES d.s TAGS d.t " +
		"RECENT SAMPLES INNER UUID '22222222-2222-2222-2222-222222222222' " +
		"SAMPLES INNER ENGINE = Memory " +
		"RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY id"
	stmts, err := NewParser(sql).ParseStmts()
	require.NoError(t, err)
	create := stmts[0].(*CreateTable)
	require.Len(t, create.TimeSeriesTargets, 3)
	require.Equal(t, Pos(len(sql)), create.End())

	samples := create.TimeSeriesTargets[0]
	require.Equal(t, "samples", samples.Kind)
	require.Equal(t, "d.s", samples.External.String())
	require.Equal(t, "Memory", samples.InnerEngine.Name)
	require.False(t, samples.EngineShorthand)

	recent := create.TimeSeriesTargets[2]
	require.Equal(t, "recent_samples", recent.Kind)
	require.Equal(t, "RECENT SAMPLES", recent.Keyword)
	require.Equal(t, "'22222222-2222-2222-2222-222222222222'", recent.InnerUUID.Value.String())
	require.Equal(t, "MergeTree", recent.InnerEngine.Name)
	require.Equal(t, Pos(len(sql)), recent.End())

	var uuids int
	Walk(create, func(expr Expr) bool {
		if _, ok := expr.(*UUID); ok {
			uuids++
		}
		return true
	})
	require.Equal(t, 1, uuids)
}

// TestParser_TimeSeriesDuplicateTarget verifies that repeating a part of a
// target (including across the DATA/SAMPLES alias pair) is rejected, as
// ClickHouse does, and that malformed target keywords fail.
func TestParser_TimeSeriesDuplicateTarget(t *testing.T) {
	for _, sql := range []string{
		"CREATE TABLE db.m ENGINE = TimeSeries DATA db.a SAMPLES db.b",
		"CREATE TABLE db.m ENGINE = TimeSeries TAGS db.a TAGS db.b",
		"CREATE TABLE d.ts ENGINE = TimeSeries RECENT SAMPLES d.r RECENT SAMPLES d.q",
		"CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES INNER ENGINE = MergeTree ORDER BY id SAMPLES INNER ENGINE = Memory",
		"CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES ENGINE = MergeTree ORDER BY id SAMPLES INNER ENGINE = Memory",
		"CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES INNER COLUMNS (a UUID) SAMPLES INNER COLUMNS (b UUID)",
		"CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES INNER UUID '00000000-0000-0000-0000-000000000001' SAMPLES INNER UUID '00000000-0000-0000-0000-000000000002'",
		"CREATE TABLE d.ts ENGINE = TimeSeries RECENT d.r",
		"CREATE TABLE d.ts ENGINE = TimeSeries RECENT DATA d.r",
		"CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES COLUMNS (id UUID)",
		"CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES `RECENT` SAMPLES d.r",
	} {
		_, err := NewParser(sql).ParseStmts()
		require.Error(t, err, "expected duplicate-target error for: %s", sql)
	}
}
