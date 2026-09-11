# ClickHouse source SQL corpus

Extracted from the local `../ClickHouse` checkout at revision
`387d28e024845a304c45183b06b792fb1352d10e`. This corpus is separate from both
`docs-sql/` and the permanent regression fixtures in `parser/testdata/`.

The initial extraction scanned **11,345 SQL files** and **639 performance XML
files**, collecting **18,214 occurrences / 17,780 byte-distinct examples**.

| Parser outcome | Examples |
| --- | ---: |
| Accepted | 11,283 |
| Rejected | 6,495 |
| Panic | 0 |
| Empty / no statements | 2 |

The 56 originally panicking inputs are now permanent regression fixtures in
`parser/testdata/regressions/clickhouse_panics/`. Of those scripts, 39 parse
successfully and 17 return ordinary errors at later unsupported constructs.

## Run

```sh
make clickhouse-sql-extract  # regenerate from ../ClickHouse
make clickhouse-sql-test     # compare outcomes with the checked-in baseline
make clickhouse-sql-strict   # fail individually for every non-accepted example
make clickhouse-sql-update   # explicitly accept current outcomes as the baseline
```

The Go corpus test is opt-in. It uses the same runner as the documentation corpus,
checks each SQL file against its SHA-256 ID, and writes `clickhouse-sql/report.json`.
The report includes all outcomes, full parser errors and panic stacks, source
paths and line numbers, XML attributes, substitution flags, and expected-error
annotations. Reports are generated and ignored by Git. The extracted queries,
manifest, baseline, and source license are persistent fixtures, so running the
corpus does not require the upstream checkout or a ClickHouse server.

Baseline mode fails for changed outcomes (including improvements), additions, and
removals. Its initial success does not mean all SQL is supported: known failures
and panics are recorded in the baseline. Strict mode fails on those examples.

```sh
# Inspect panic examples and stack traces.
jq '.results[] | select(.status == "panic") | {file, error, sources}' clickhouse-sql/report.json

# Rerun an individual example using its hash ID.
go test ./parser -run '^TestClickHouseSQL$/EXAMPLE_ID' -clickhouse-sql-strict -count=1

# Alternate source/output locations for extraction.
python3 scripts/extract_clickhouse_sql.py --source ../ClickHouse --output clickhouse-sql
```

`file` is relative to `clickhouse-sql/`; source paths are relative to the ClickHouse
checkout. Lines are one-based. A filtered test run writes only its selected results;
rerun the full corpus to restore the report. Baseline updates require a full run.

## Scope and interpretation

- All `.sql` files outside hidden directories are copied **byte-for-byte**, including
  functional tests, integration fixtures, benchmarks, and utility queries. No
  semicolons are added; comments, whitespace, invalid bytes, and INSERT payloads
  remain unchanged. Identical content shares one file and retains every source.
- XML files under `tests/performance/` and `tests/perf_drafts/` contribute `query`,
  `create_query`, `fill_query`, and `drop_query` elements. XML entities and CDATA
  are decoded to their text content. SQL text is otherwise unchanged, and query
  attributes and the presence of substitutions are recorded. Parameters are not
  substituted: 1,528 distinct examples come from files declaring substitutions.
- The manifest inventories **185 unrendered `.sql.j2` / `.sql.in` templates**.
  These require a generation step and are not sent to the SQL parser as if they
  were already SQL. SQL embedded in Python, shell, C++, Markdown, and sqllogic
  records is outside this extraction scope.
- `serverError` and `clientError` annotations remain in the SQL and are also
  exposed in source metadata. There are 2,509 examples with such annotations.
  Some errors are intentional syntax errors; others concern server execution.
  They are not automatically treated as parser success or filtered out.
- Each SQL file is parsed as a whole script via `ParseStmts`; XML query elements
  are separate examples. A failure reports the first parser error in that script,
  so later statements are not independently classified. “Accepted” means one or
  more AST statements were returned without error; it does not establish server
  validity or semantic correctness. These are example counts, not feature coverage
  percentages.
- Panics are recovered per example and their stacks are recorded so the rest of
  the corpus can run. The three-minute Go test timeout bounds the entire run;
  an unrecoverable process failure or hang can still abort a run.

SQL is from the ClickHouse project; see the copied [Apache 2.0 license](LICENSE).
`manifest.json` preserves source attribution and the checkout revision. The
revision identifies the source checkout's HEAD; local source modifications, if
any, are reflected in extracted content hashes.
