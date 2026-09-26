package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_TableStream verifies the streaming-query modifier
// `FROM t [AS x] [FINAL] STREAM [BOUNDED] [UNORDERED]` (#60). STREAM used to
// be read as an implicit table alias, which silently turned a streaming
// query into a regular one (`FROM t AS STREAM`).
func TestParser_TableStream(t *testing.T) {
	for _, sql := range []string{
		"SELECT count() FROM t STREAM",
		"SELECT count() FROM t AS x STREAM",
		"SELECT count() FROM t FINAL STREAM",
		"SELECT count() FROM t STREAM BOUNDED UNORDERED",
		"SELECT * FROM t STREAM JOIN u ON t.a = u.a",
		"SELECT * FROM (SELECT 1) STREAM",
		"SELECT * FROM numbers(10) STREAM",
		"SELECT count() FROM t STREAM WHERE x > 1",
		"SELECT n FROM t STREAM LIMIT 1023",
		"SELECT count() FROM t_stream_gating STREAM SETTINGS enable_streaming_queries=1",
		// STREAM is not reserved: an explicit alias and a column keep working.
		"SELECT count() FROM t AS STREAM",
		"SELECT STREAM FROM t",
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
}

func TestParser_TableStreamRejected(t *testing.T) {
	for _, sql := range []string{
		"SELECT count() FROM t STREAM AS x",
		"SELECT count() FROM t STREAM FINAL",
		"SELECT count() FROM t STREAM SAMPLE 0.1",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
}
