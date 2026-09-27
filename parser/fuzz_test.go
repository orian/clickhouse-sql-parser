package parser

import (
	"crypto/sha256"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

var fuzzSQLCorpora = flag.String("fuzz-sql-corpora", "", "comma-separated directories of additional SQL fuzz seeds (up to 256 per directory)")

const maxFuzzSQLBytes = 16 * 1024

var sqlFixtureDirs = []string{"testdata/basic", "testdata/ddl", "testdata/dml", "testdata/query"}

type sqlSeed struct {
	name string
	sql  string
}

// A zero limit includes every fixture. External corpora use a stable hash order
// so their bounded sample is spread across filenames rather than one prefix.
func readSQLSeeds(t testing.TB, dir string, limit int) []sqlSeed {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if limit > 0 {
		sort.Slice(entries, func(i, j int) bool {
			a := sha256.Sum256([]byte(entries[i].Name()))
			b := sha256.Sum256([]byte(entries[j].Name()))
			return string(a[:]) < string(b[:])
		})
	}
	var seeds []sqlSeed
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		name := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}
		if len(data) > maxFuzzSQLBytes {
			continue
		}
		seeds = append(seeds, sqlSeed{name: name, sql: string(data)})
		if limit > 0 && len(seeds) == limit {
			break
		}
	}
	return seeds
}

// FuzzParseStmts exercises both valid SQL and malformed mutations. Parse errors
// are expected; panics and hangs are failures. Go saves and minimizes failures in
// testdata/fuzz/FuzzParseStmts so ordinary go test runs replay them automatically.
func FuzzParseStmts(f *testing.F) {
	addSQLFuzzSeeds(f)
	f.Fuzz(func(t *testing.T, sql string) {
		if len(sql) > maxFuzzSQLBytes {
			t.Skip()
		}
		_, _ = NewParser(sql).ParseStmts()
	})
}

// FuzzAST checks consumers of successfully parsed trees. Ordinary parse and
// visitor errors are allowed; this target isolates the no-panic invariant from
// formatting correctness and round-trip equivalence.
func FuzzAST(f *testing.F) {
	addSQLFuzzSeeds(f)
	f.Fuzz(func(t *testing.T, sql string) {
		if len(sql) > maxFuzzSQLBytes {
			t.Skip()
		}
		stmts, err := NewParser(sql).ParseStmts()
		if err != nil {
			return
		}
		for _, stmt := range stmts {
			walked := map[Expr]bool{}
			Walk(stmt, func(node Expr) bool {
				walked[node] = true
				_ = node.Pos()
				_ = node.End()
				// PrintVisitor applied to any subtree prints the same SQL as
				// the node's String() (#98).
				if got, err := printNode(node); err == nil && got != node.String() {
					t.Fatalf("PrintVisitor(%T) = %q, String() = %q", node, got, node.String())
				} else if err != nil && strings.HasPrefix(err.Error(), "panic") {
					t.Fatalf("PrintVisitor(%T) %v", node, err)
				}
				return true
			})
			// Walk reaches every node the visitor recursion reaches (#105).
			entered := &enteredNodes{seen: map[Expr]bool{}}
			entered.Self = entered
			if err := stmt.Accept(entered); err == nil {
				for _, node := range entered.list {
					if !walked[node] {
						t.Fatalf("Walk misses %T %q in %q", node, node.String(), stmt.String())
					}
				}
			}
			_ = stmt.String()
			printer := NewPrintVisitor()
			_ = stmt.Accept(printer)
			_ = printer.String()
			beautifier := NewBeautifyVisitor()
			_ = stmt.Accept(beautifier)
			_ = beautifier.String()
		}
	})
}

func addSQLFuzzSeeds(f *testing.F) {
	f.Helper()
	for _, sql := range []string{
		"", "SELECT 1", "SELECT (", "SELECT '\\'", "-- comment\nSELECT 1",
		"/* comment */ SELECT 1; SELECT 2", "SELECT * FROM numbers(1_000)",
		"SELECT * FROM numbers(n=)", "SET x = 0.1", "SET x =",
		"CREATE TABLE t (x UInt64) ENGINE = MergeTree ORDER BY x",
		"CREATE TABLE t (x UInt64) ENGINE = Kafka SETTINGS kafka_topic_list = 't' SETTINGS flatten_nested = 0",
		"CREATE TABLE t (x UInt64) ENGINE = MergeTree ORDER BY x SETTINGS a = 1 SETTINGS b = 2 SETTINGS c = 3",
		"CREATE TABLE t (x UInt64) ENGINE = MergeTree ORDER BY x ORDER BY (x, x)",
		"CREATE TABLE t ENGINE = TimeSeries SAMPLES d.s RECENT SAMPLES INNER UUID '00000000-0000-0000-0000-000000000001' SAMPLES INNER ENGINE = Memory",
		"CREATE TABLE t ENGINE = TimeSeries RECENT SAMPLES d.r RECENT SAMPLES d.q",
		"CREATE TABLE t ENGINE = TimeSeries RECENT",
		"CREATE TABLE t () ENGINE = Memory", "CREATE TABLE t AS other ENGINE = Memory",
		"DETACH TABLE IF EXISTS t ON CLUSTER c PERMANENTLY SYNC", "ATTACH DICTIONARY d", "ATTACH TABLE t",
		"DETACH DATABASE db PERMANENTLY", "ATTACH OR REPLACE TABLE t", "DROP TABLE t PERMANENTLY",
		"SELECT (3,), ((3,),), (*,).1, (1, 2,)", "SELECT plus(1, 2,)", "SELECT [1,]",
		"CREATE TABLE t (x UInt64) ENGINE = Memory COMMENT 'c' SETTINGS max_threads = 1",
		"INSERT INTO t VALUES (1), (2)", "ALTER TABLE t ATTACH PARTITION ALL",
		"INSERT INTO FUNCTION mysql('h', 'db', 't', 'u', 'p') (a, b) VALUES (1, 2)", "INSERT INTO TABLE FUNCTION null() SELECT 1",
		"ALTER TABLE t DROP DETACHED PARTITION ID '1', DELETE IN PARTITION ID '2' WHERE 1", "OPTIMIZE TABLE t PARTITION ID 'all' FINAL",
		"CREATE USER u NOT IDENTIFIED", "CREATE USER u SETTINGS PROFILE p",
		"CREATE USER u HOST LOCAL HOST ANY SETTINGS a = 1 SETTINGS b = 2",
		"CREATE USER u IDENTIFIED BY 'a' IDENTIFIED BY 'b'",
		"CREATE USER u DEFAULT DATABASE db DEFAULT DATABASE NONE",
		"SELECT a FROM t GROUP BY a WITH ROLLUP WITH TOTALS",
		"SELECT a FROM t GROUP BY a WITH TOTALS WITH CUBE",
		"CREATE LIVE VIEW v WITH TIMEOUT AS SELECT 1",
		"SELECT CAST(a = b, 'Bool') FROM t", "SELECT {value:UInt64}",
		"SELECT -1::Int32, - -1, 2 - -1, -x::Int8, +1e-5::Float64", "SELECT - -",
		"SELECT groupArraySample(5, 1)(DISTINCT x), quantilesTimingIf(0.1)(DISTINCT x, y) FROM t",
		"WITH x AS (SELECT 1) SELECT * FROM x", "SYSTEM STOP MERGES t",
		"GRANT SELECT(x, y) ON db.t TO u", "SELECT sum(x) OVER (ROWS BETWEEN UNBOUNDED PRECEDING AND {n:UInt32} FOLLOWING) FROM t GROUP BY ALL",
		"SELECT a FROM t ORDER BY a DESC NULLS LAST COLLATE 'en' WITH FILL, b NULLS FIRST", "SELECT 1 ORDER BY a NULLS",
		"((SELECT 1 UNION ALL SELECT 2) EXCEPT (SELECT 3)) UNION DISTINCT SELECT 4 FORMAT JSON", "(SELECT 1", "(SELECT 1))",
		"SELECT CAST(x AS Enum('a' = 1)), `q`, 'x' IS NOT NULL FROM 't' FINAL GROUP BY a ORDER BY a DESC", "SELECT CASE WHEN 1 THEN 2 END",
		"SELECT a FROM t ORDER BY a LIMIT 1 BY a LIMIT 2, 3 WITH TIES", "SELECT 1 OFFSET 5", "SELECT a FROM t LIMIT 1 WITH",
		"SELECT 1 SETTINGS a = 1 FORMAT JSON SETTINGS b = 2", "SHOW TABLES FORMAT JSON SETTINGS a = 1", "DROP TABLE t FORMAT",
		"SELECT 1 FROM a GLOBAL ANY LEFT JOIN b ON a.x = b.x GLOBAL CROSS JOIN c", "SELECT 1 FROM a LOCAL JOIN b USING (x)",
		"SELECT count() FROM t WITH TOTALS HAVING count() > 1 ORDER BY 1", "SELECT count() FROM t WITH",
		"SELECT 1 UNION DISTINCT (SELECT 1 UNION ALL (SELECT 2 EXCEPT SELECT 3))",
		"SELECT count() FROM t AS x FINAL STREAM BOUNDED UNORDERED JOIN u STREAM ON t.a = u.a", "SELECT 1 FROM t STREAM",
		"WITH RECURSIVE t AS (SELECT 1 AS n UNION ALL SELECT n + 1 FROM t WHERE n < 10) SELECT sum(n) FROM t",
		"WITH RECURSIVE t AS (SELECT 1 UNION ALL SELECT * FROM t",
		"CREATE DICTIONARY d (x UInt64) PRIMARY KEY x SOURCE(CLICKHOUSE(DB currentDatabase() TABLE 't')) LAYOUT(FLAT()) LIFETIME(0)",
	} {
		f.Add(sql)
	}
	dirs := append([]string{"testdata/regressions/create_user", "testdata/regressions/clickhouse_panics"}, sqlFixtureDirs...)
	for _, dir := range dirs {
		for _, seed := range readSQLSeeds(f, dir, 0) {
			f.Add(seed.sql)
		}
	}
	if *fuzzSQLCorpora != "" {
		for _, dir := range strings.Split(*fuzzSQLCorpora, ",") {
			for _, seed := range readSQLSeeds(f, dir, 256) {
				f.Add(seed.sql)
			}
		}
	}
}
