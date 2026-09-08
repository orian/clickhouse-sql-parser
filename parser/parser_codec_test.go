package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_GCDCodecs(t *testing.T) {
	for _, codec := range []string{"GCD", "GCD, Default", "GCD, ZSTD(1)"} {
		t.Run(codec, func(t *testing.T) {
			sql := "CREATE TABLE db.events (timestamp DateTime64(6) CODEC(" + codec + ")) ENGINE = MergeTree ORDER BY timestamp"
			stmts, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
			require.Len(t, stmts, 1)
			require.Equal(t, sql, stmts[0].String())
		})
	}
}
