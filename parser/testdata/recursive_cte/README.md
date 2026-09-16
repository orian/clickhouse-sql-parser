# Recursive CTE fixtures

The eight files under `docs/` are the SQL blocks containing `WITH RECURSIVE`
from ClickHouse's official [`WITH` documentation][with-docs], extracted verbatim
from ClickHouse revision `9b4f1726945218e95ab9b47d9b5fc164e3b68570` on
2026-09-16.

They were validated against a real ClickHouse server in Docker:

- image: `clickhouse/clickhouse-server:26.6`
- server version: `26.6.1.1193`
- `sum_integers.sql` and `infinite_with_limit.sql`: result `5050`
- all three tree traversal examples: the documented four rows and ordering
- graph traversal without a cycle: the documented five rows
- `graph_with_cycle.sql`: the documented `TOO_DEEP_RECURSION` error (validated
  with `max_recursive_cte_evaluation_depth = 10` to keep the check bounded)
- graph cycle detection: the documented three cycle rows

The Go regression test parses every file and checks AST string, compact, and
beautified formatting round trips. The graph fixtures do not need their setup
DDL during parser tests; the live validation used the table definitions and
seed rows from the same documentation page.

[with-docs]: https://github.com/ClickHouse/ClickHouse/blob/9b4f1726945218e95ab9b47d9b5fc164e3b68570/docs/reference/statements/select/with.mdx
