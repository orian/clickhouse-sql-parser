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
			// The empty () schema and SETTINGS spacing are normalised by the
			// existing CreateTable/EngineExpr formatters (pre-existing behaviour),
			// so spell out the expected round-trip output for these cases.
			name: "external SAMPLES canonical",
			sql:  "CREATE TABLE db.m () ENGINE = TimeSeries SAMPLES db.m_samples TAGS db.m_tags METRICS db.m_metrics",
			want: "CREATE TABLE db.m  ENGINE = TimeSeries SAMPLES db.m_samples TAGS db.m_tags METRICS db.m_metrics",
		},
		{
			name: "bare TimeSeries",
			sql:  "CREATE TABLE db.m () ENGINE = TimeSeries",
			want: "CREATE TABLE db.m  ENGINE = TimeSeries",
		},
		{
			name: "settings only",
			sql:  "CREATE TABLE db.m () ENGINE = TimeSeries SETTINGS id_generator = 'sipHash64(metric_name, all_tags)'",
			want: "CREATE TABLE db.m  ENGINE = TimeSeries SETTINGS id_generator='sipHash64(metric_name, all_tags)'",
		},
		{
			name: "inner columns and inner engine",
			sql: "CREATE TABLE db.m () ENGINE = TimeSeries " +
				"SAMPLES INNER COLUMNS (`id` UUID, `timestamp` DateTime64(3), `value` Float64) " +
				"SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp)",
			want: "CREATE TABLE db.m  ENGINE = TimeSeries " +
				"SAMPLES INNER COLUMNS (`id` UUID, `timestamp` DateTime64(3), `value` Float64) " +
				"SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp)",
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

			// Re-parsing the formatted output must succeed and converge to a
			// fixpoint. (The very first pass can collapse an empty `()` schema,
			// a pre-existing formatter quirk, so compare the 2nd and 3rd
			// generations rather than the 1st and 2nd.)
			gen2, err := NewParser(stmts[0].String()).ParseStmts()
			require.NoError(t, err)
			gen3, err := NewParser(gen2[0].String()).ParseStmts()
			require.NoError(t, err)
			require.Equal(t, gen2[0].String(), gen3[0].String())
		})
	}
}

// TestParser_TimeSeriesDuplicateTarget verifies that repeating a target slot
// (including the DATA/SAMPLES alias pair) is rejected.
func TestParser_TimeSeriesDuplicateTarget(t *testing.T) {
	for _, sql := range []string{
		"CREATE TABLE db.m () ENGINE = TimeSeries DATA db.a SAMPLES db.b",
		"CREATE TABLE db.m () ENGINE = TimeSeries TAGS db.a TAGS db.b",
	} {
		_, err := NewParser(sql).ParseStmts()
		require.Error(t, err, "expected duplicate-target error for: %s", sql)
	}
}
