# Parser fuzzing

Go's built-in coverage-guided fuzzer mutates SQL and exercises `ParseStmts`.
Malformed or unsupported input may return an ordinary error; panics and hangs
are bugs. The target does not execute SQL or require a ClickHouse server.

From the repository root:

```sh
make fuzz                         # 30 seconds, four workers, five-minute overall timeout
make fuzz FUZZ_TIME=10m FUZZ_TIMEOUT=15m FUZZ_PARALLEL=2
```

Go minimizes a failure and saves it under `parser/testdata/fuzz/FuzzParseStmts/`.
Its output includes a command for reproducing that exact case. From the repository
root, add `./parser`, for example:

```sh
go test ./parser -run '^FuzzParseStmts$/002ef2e526c1c171$'
```

Keep minimized failure files when fixing a bug: ordinary `go test ./...` and
`make test` replay them automatically, without running mutation-based fuzzing.
New coverage-increasing inputs also accumulate in Go's local build-cache fuzz
corpus. A clean cache may explore different paths. Passing a bounded run does not
prove the absence of parser bugs.

## Seeds and scope

The default seeds include small examples of SQL statement types, malformed input,
and existing documentation/source panic regression fixtures. Inputs over 16 KiB
are skipped to keep mutation cost bounded; the full corpus tests cover larger
scripts separately. Error formatting is exercised naturally by rejected inputs.
AST formatting, AST traversal, and comparison against ClickHouse are separate
checks, outside this target's parse-only invariant.

Additional seeds can come from either extracted corpus. Paths are relative to the
`parser/` package directory; each directory contributes at most 256 eligible SQL
files in sorted filename order:

```sh
make fuzz FUZZ_TIME=2m FUZZ_SQL_CORPORA='../docs-sql/queries,../clickhouse-sql/queries'
```

Seed SQL is read unchanged. The imported fixtures retain the upstream licenses
and attribution documented in their corpus directories. A failing fuzz input is
stored in Go's `go test fuzz v1` encoding, not as a rewritten original fixture.

## Automation

The `Parser fuzzing` GitHub Actions workflow runs for two minutes nightly and can
also be dispatched manually. On failure it uploads the minimized regression
corpus. Ordinary CI continues to replay saved seeds through `make test`; it does
not depend on a random mutation campaign passing on every pull request.
