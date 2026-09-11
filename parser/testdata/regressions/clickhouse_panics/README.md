# ClickHouse source panic regressions

These 56 SQL files are byte-for-byte copies of the initially panicking examples
from `clickhouse-sql/queries/`. Filenames are their SHA-256 content hashes.
`sources.json` records their upstream paths and line numbers. See
`clickhouse-sql/LICENSE` for the source license and `clickhouse-sql/manifest.json`
for the checkout revision.

The default Go test suite parses every file and requires that none panic.
Some full upstream scripts intentionally contain invalid SQL or reach later
unsupported constructs, so an ordinary parse error is permitted. Focused tests
in `parser_source_panics_test.go` separately require successful parsing for
numeric separators and `PARTITION ALL`, check source positions and round trips,
and require useful errors for malformed input. The complete opt-in source corpus
baseline records acceptance/rejection changes for every example.
