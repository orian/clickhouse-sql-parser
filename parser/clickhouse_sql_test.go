package parser

import (
	"flag"
	"testing"
)

var clickhouseSQL = flag.Bool("clickhouse-sql", false, "parse the ClickHouse source corpus and compare its baseline")
var clickhouseSQLUpdate = flag.Bool("clickhouse-sql-update", false, "parse the ClickHouse source corpus and replace its baseline")
var clickhouseSQLStrict = flag.Bool("clickhouse-sql-strict", false, "fail for each ClickHouse source example rejected by the parser")

func TestClickHouseSQL(t *testing.T) {
	if !*clickhouseSQL && !*clickhouseSQLUpdate && !*clickhouseSQLStrict {
		t.Skip("use -clickhouse-sql to parse the ClickHouse source corpus")
	}
	runSQLCorpus(t, "clickhouse-sql", *clickhouseSQLUpdate, *clickhouseSQLStrict)
}
