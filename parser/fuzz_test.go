package parser

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var fuzzSQLCorpora = flag.String("fuzz-sql-corpora", "", "comma-separated directories of additional SQL fuzz seeds (up to 256 per directory)")

const maxFuzzSQLBytes = 16 * 1024

// FuzzParseStmts exercises both valid SQL and malformed mutations. Parse errors
// are expected; panics and hangs are failures. Go saves and minimizes failures in
// testdata/fuzz/FuzzParseStmts so ordinary go test runs replay them automatically.
func FuzzParseStmts(f *testing.F) {
	for _, sql := range []string{
		"", "SELECT 1", "SELECT (", "SELECT '\\'", "-- comment\nSELECT 1",
		"/* comment */ SELECT 1; SELECT 2", "SELECT * FROM numbers(1_000)",
		"SELECT * FROM numbers(n=)", "SET x = 0.1", "SET x =",
		"CREATE TABLE t (x UInt64) ENGINE = MergeTree ORDER BY x",
		"INSERT INTO t VALUES (1), (2)", "ALTER TABLE t ATTACH PARTITION ALL",
		"CREATE USER u NOT IDENTIFIED", "CREATE USER u SETTINGS PROFILE p",
		"SELECT CAST(a = b, 'Bool') FROM t", "SELECT {value:UInt64}",
		"WITH x AS (SELECT 1) SELECT * FROM x", "SYSTEM STOP MERGES t",
	} {
		f.Add(sql)
	}
	dirs := []string{"testdata/regressions/create_user", "testdata/regressions/clickhouse_panics"}
	if *fuzzSQLCorpora != "" {
		dirs = append(dirs, strings.Split(*fuzzSQLCorpora, ",")...)
	}
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir) // Sorted, so optional seed selection is reproducible.
		if err != nil {
			f.Fatal(err)
		}
		added := 0
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}
			data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
			if err != nil {
				f.Fatal(err)
			}
			if len(data) > maxFuzzSQLBytes {
				continue
			}
			f.Add(string(data))
			added++
			if added == 256 {
				break
			}
		}
	}
	f.Fuzz(func(t *testing.T, sql string) {
		// Keep mutation cost bounded; the existing corpora cover larger scripts.
		if len(sql) > maxFuzzSQLBytes {
			t.Skip()
		}
		_, _ = NewParser(sql).ParseStmts()
	})
}
