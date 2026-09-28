# ClickHouse SQL Parser (fork)

This is a fork of [AfterShip/clickhouse-sql-parser](https://github.com/AfterShip/clickhouse-sql-parser) with a different AST-traversal architecture and a built-in beautifier. The headline differences:

- **All AST traversal goes through visitors.** Each `Expr.Accept` is a uniform three-liner; children are walked by `DefaultASTVisitor`, never by the nodes themselves.
- **`PrintVisitor`** produces compact, one-line SQL output (replaces `Expr.String()` for serialization).
- **`BeautifyVisitor`** produces indented, line-broken SQL with per-clause inline thresholds — also a visitor, not a method on the nodes.

The full rationale and the merge-from-upstream checklist live in [`refactor-visitor.md`](./refactor-visitor.md). The parser surface (lexer, AST node types, semantics) tracks upstream `master`.

## What this parser does

- Parse ClickHouse SQL into a typed AST
- Render an AST back to compact SQL (`PrintVisitor`)
- Render an AST as indented, line-broken SQL (`BeautifyVisitor`)
- Traverse / mutate / search an AST with either a visitor or the `Walk` helper

Inspired by [memefish](https://github.com/cloudspannerecosystem/memefish).

## Use as a Go library

Parse SQL into an AST:

```go
import clickhouse "github.com/orian/clickhouse-sql-parser/parser"

stmts, err := clickhouse.NewParser("SELECT * FROM t WHERE a = 100").ParseStmts()
if err != nil {
    return err
}
```

Render an AST back to one-line SQL via `PrintVisitor`:

```go
for _, stmt := range stmts {
    printer := clickhouse.NewPrintVisitor()
    if err := stmt.Accept(printer); err != nil {
        return err
    }
    fmt.Println(printer.String())
}
```

Render an AST as beautified SQL via `BeautifyVisitor`:

```go
for _, stmt := range stmts {
    b := clickhouse.NewBeautifyVisitor()
    if err := stmt.Accept(b); err != nil {
        return err
    }
    fmt.Println(b.String())
}
```

Customise per-clause line-wrapping thresholds:

```go
b := clickhouse.NewBeautifyVisitorWith(clickhouse.BeautifyOptions{
    MaxWidth:            100, // expression-level wrap budget
    SelectItemsMaxLen:   80,
    SelectItemsMaxCount: 7,
    FromMaxLen:          80,
    WhereMaxLen:         80,
    WhereMaxConjuncts:   1,
    OrderByMaxLen:       80,
    OrderByMaxCount:     5,
})
```

Defaults (`clickhouse.DefaultBeautifyOptions`) match what the CLI uses.

## Use as a CLI

Build from this checkout:

```bash
$ go build -o clickhouse-sql-parser .
```

(`go install github.com/orian/clickhouse-sql-parser@latest` installs this fork's binary. The upstream project lives at `github.com/AfterShip/clickhouse-sql-parser`.)

Commands:

```bash
## Parse query into AST (JSON dump)
$ clickhouse-sql-parser "SELECT * FROM clickhouse WHERE a=100"

## Compact one-line format
$ clickhouse-sql-parser -format "SELECT * FROM clickhouse WHERE a=100"

## Beautify (indented, line-broken)
$ clickhouse-sql-parser -beautify "SELECT a, b FROM t WHERE a > 1"

## Parse query from file
$ clickhouse-sql-parser -f ./test.sql
```

Beautify example:

```sql
-- input
SELECT coalesce(very_long_function_name_a(a, b, c), very_long_function_name_b(d, e, f), some_default_value) FROM t

-- output
SELECT
  coalesce(
    very_long_function_name_a(a, b, c),
    very_long_function_name_b(d, e, f),
    some_default_value
  )
FROM t
```

The beautifier keeps short clauses on one line (`SELECT a, b, c FROM t`) and only line-breaks when an item count or rendered length exceeds the per-clause threshold; long function calls split arguments recursively (a nested call whose arguments also overflow splits next).

## AST Traversal

Two complementary mechanisms ship on this fork:

### Custom visitor (recommended for new code)

Embed `DefaultASTVisitor`, override the `VisitX` methods you care about, set `Self = v`:

```go
type myVisitor struct {
    clickhouse.DefaultASTVisitor
    tables []string
}

func (v *myVisitor) VisitTableIdentifier(t *clickhouse.TableIdentifier) error {
    v.tables = append(v.tables, t.String())
    return v.DefaultASTVisitor.VisitTableIdentifier(t) // continue default descent
}

v := &myVisitor{}
v.Self = v
_ = stmt.Accept(v)
fmt.Println(v.tables)
```

The `Self = v` line is required so the default child-walking dispatches back to your overrides; see [`refactor-visitor.md`](./refactor-visitor.md) for the rationale.

### Walk helpers

For one-off traversals where a full visitor is overkill:

```go
clickhouse.Walk(stmts[0], func(node clickhouse.Expr) bool {
    fmt.Printf("Node type: %T\n", node)
    return true // return false to stop descending this subtree
})
```

Available functions:

- `Walk(node Expr, fn WalkFunc)` — depth-first traversal
- `WalkWithBreak(node Expr, fn WalkFunc)` — same, with early termination
- `Find(root Expr, predicate func(Expr) bool) (Expr, bool)`
- `FindAll(root Expr, predicate func(Expr) bool) []Expr`
- `Transform(root Expr, transformer func(Expr) Expr) Expr`

## Known differences from ClickHouse

The parser follows the grammar of ClickHouse 26.8 and rejects what ClickHouse
rejects, with a few deliberate exceptions. It also accepts:

- **`?` placeholders**, e.g. `SELECT * FROM t WHERE id = ?`. Client-side
  drivers use them; ClickHouse itself only has `{name:Type}` parameters.
- **`ATTACH DICTIONARY d (…) PRIMARY KEY … SOURCE(…) …` with a full
  definition.** ClickHouse's SQL parser accepts only `ATTACH DICTIONARY d`,
  but its metadata files (`metadata/<db>/<dictionary>.sql`) store dictionaries
  in this form, and tools that read them need to parse it.
- **`INSERT INTO t [(columns)]` with no `VALUES`, `FORMAT` or `SELECT`.**
  Clients send such inserts with the data out of band, and query logs record
  them this way.
- **`CREATE LIVE VIEW`**, which ClickHouse 26.8 removed but older servers
  support.

The parser's output can also differ from ClickHouse's own formatter without
changing the meaning (for example `LAYOUT(FLAT)` is printed as
`LAYOUT(FLAT())`); see [#66](https://github.com/orian/clickhouse-sql-parser/issues/66).

## Testing

```bash
$ go test ./...                 # full suite
$ go test -v ./parser/...       # with test names
$ make test                     # full suite with -race (what CI runs)
$ make cover                    # same, with a coverage profile
```

The beautifier has two test layers:

1. **Hand-curated fixtures** in `parser/testdata/beautify/` — pairs of `<name>.in.sql` / `<name>.expected.sql`. Each case is one tightly-scoped behaviour (e.g. `from_inline_when_short`, `wrap_nested_function`, `select_items_inline_seven`). Drop two files in to add a new case; no Go code edit needed.
2. **Corpus sweep** — every input under `parser/testdata/{basic,ddl,dml,query}/*.sql` is beautified and diffed against a golden in the matching `format/beautify/` subdirectory.

To re-pin goldens after an intentional change:

```bash
$ go test ./parser -run TestBeautifyVisitor_Fixtures -update-beautify  # fixtures
$ go test -v ./parser -run TestParser_Beautify -update                 # corpus
```

To regenerate the AST and compact-format goldens:

```bash
$ make update_test
```

## Benchmarks

```bash
$ make bench                       # parse, String() and PrintVisitor benchmarks
$ make bench-compare BASE=v1.0.4   # compare BASE with HEAD on identical inputs
```

`make bench` covers parsing every `testdata/query` fixture and the large
PostHog queries, and printing them with `String()` and `PrintVisitor` (the
formatter behind `-format`). `make bench-compare` (`scripts/bench_compare.sh`)
checks both refs out into temporary worktrees, runs the current benchmarks in
both on the fixtures they both parse, alternating rounds, and prints a
`benchstat` table.

v1.1.0 against v1.0.4 (AMD Ryzen AI 9 HX 370, 90 fixtures, 10 rounds):
`PrintVisitor` is about twice as fast with ~60% less memory; parsing
allocates ~5% less, and the large PostHog queries parse 6–11% faster.

## Contact

Open an issue or discussion. For the architectural divergence specifically, [`refactor-visitor.md`](./refactor-visitor.md) is the source of truth — including the rules for merging future upstream changes.

### Documentation SQL coverage

The separate [documentation corpus](docs-sql/README.md) contains extracted SQL
examples from the local ClickHouse documentation, with source locations and a
parser outcome baseline. Run `make docs-sql-test` to regenerate the detailed
acceptance/failure report, or `make docs-sql-strict` to fail on each rejected
example. Regenerate fixtures with `make docs-sql-extract`.

The [ClickHouse source corpus](clickhouse-sql/README.md) adds upstream functional,
integration, benchmark, and performance SQL from `../ClickHouse`. Run
`make clickhouse-sql-test` for its separate baseline and detailed report, or
`make clickhouse-sql-extract` to refresh the extracted examples.

### Fuzz testing

Run `make fuzz` for a bounded Go fuzzing session. See
[parser fuzzing](parser/FUZZING.md) for seed selection, failure reproduction,
and the nightly/manual workflow.
