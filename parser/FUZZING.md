# Parser fuzzing

Go's built-in coverage-guided fuzzer mutates SQL and exercises `ParseStmts`.
Malformed or unsupported input may return an ordinary error; panics and hangs
are bugs. The target does not execute SQL or require a ClickHouse server.

From the repository root:

```sh
make fuzz                         # 30 seconds, four workers, five-minute overall timeout
make fuzz FUZZ_TIME=10m FUZZ_TIMEOUT=15m FUZZ_PARALLEL=2
make fuzz FUZZ_TARGET=FuzzAST      # exercise consumers of successfully parsed ASTs
go test ./parser -run '^TestParser_FixturePrefixes$'
```

Go minimizes a failure and saves it under `parser/testdata/fuzz/<target>/`.
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

Both targets use small examples of SQL statement types, malformed input,
existing documentation/source panic regression fixtures, and all eligible SQL
fixtures directly under `testdata/{basic,ddl,dml,query}`. Inputs over 16 KiB
are skipped to keep mutation cost bounded; the full corpus tests cover larger
scripts separately. Error formatting is exercised naturally by rejected inputs.
`FuzzAST` walks successfully parsed trees, calls each visited node's `Pos` and
`End`, and exercises statement `String`, `PrintVisitor`, and `BeautifyVisitor`.
Parse and visitor errors are allowed; panics fail. Formatting equivalence and
comparison against ClickHouse remain outside these targets' invariants.

`TestParser_FixturePrefixes` runs during ordinary tests. It parses every byte
prefix of fixtures up to 1 KiB, including cuts inside quotes and comments. For
larger fixtures up to 16 KiB it checks token starts and ends, plus the empty and
complete input, to bound runtime. A prefix may parse successfully or return an
error, but must never panic. Failures identify the fixture, byte offset, and SQL.

Additional seeds can come from either extracted corpus. Paths are relative to the
`parser/` package directory; each directory contributes at most 256 eligible SQL
files in a deterministic hash order of filenames, avoiding the previous
alphabetical-prefix bias:

```sh
make fuzz FUZZ_TIME=2m FUZZ_SQL_CORPORA='../docs-sql/queries,../clickhouse-sql/queries'
```

Seed SQL is read unchanged. The imported fixtures retain the upstream licenses
and attribution documented in their corpus directories. A failing fuzz input is
stored in Go's `go test fuzz v1` encoding, not as a rewritten original fixture.

## Automation

The `Parser fuzzing` GitHub Actions workflow is dispatched manually and runs for
five minutes per target in independent matrix jobs with seeds from both
extracted SQL corpora.
On failure it uploads the minimized regression
corpus. Ordinary CI continues to replay saved seeds through `make test`; it does
not depend on a random mutation campaign passing on every pull request.
