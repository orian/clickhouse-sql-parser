# ClickHouse documentation SQL corpus

Extracted from the local `clickhouse-docs/` copy of the official ClickHouse
 documentation. This corpus is separate from `parser/testdata`; no parser behavior
was changed to accommodate it.

Current results: **12,390 unique examples**, representing **89,716 occurrences**
in 25,035 Markdown, MDX, and SQL source files (all translations included).

| Parser outcome | Unique examples |
| --- | ---: |
| Accepted | 6,909 |
| Rejected | 5,480 |
| Panic | 0 |
| Empty / no statements | 1 |

## Reproduce

From the repository root, using Python 3 and Go:

```sh
make docs-sql-extract   # regenerate SQL files and manifest from local docs
make docs-sql-test      # parse every example; compare outcomes with baseline
make docs-sql-strict    # fail individually for every non-accepted example
make docs-sql-update    # explicitly accept current outcomes as the new baseline
python3 -m unittest discover -s scripts -p 'test_*.py'
```

The ordinary test suite skips this corpus. Each corpus command writes
`docs-sql/report.json`, including accepted examples, full parser errors, statement
counts for accepted blocks, and every source location. The generated report is
ignored by Git; `baseline.json`, `manifest.json`, and `queries/` are persistent
fixtures. No documentation checkout is needed to run the corpus after extraction.
The baseline test fails for new/removed examples and changed statuses, including
improvements and newly introduced panics. Known failures and panics remain visible
in the report while allowing the baseline test to pass. Strict mode deliberately
fails on the current corpus. A Go test timeout bounds the complete run.

To inspect rejected examples (substitute `panic` or `accepted` as needed):

```sh
jq '.results[] | select(.status == "rejected") | {file, error, sources}' docs-sql/report.json
```

Read `file` relative to `docs-sql/`; read source paths relative to
`clickhouse-docs/`. `line` and `end_line` are one-based documentation lines.
To rerun one example, use its SHA-256 ID:

```sh
go test ./parser -run '^TestDocumentationSQL$/EXAMPLE_ID' -docs-sql-strict -count=1
```

A filtered run generates a report of only the selected examples; rerun the full
corpus to restore the complete report. Use baseline comparison/update only for
full runs.

## Extraction scope and interpretation

- Fenced blocks labeled `sql` (case-insensitive) are the main automatic marker.
  Fence titles, `runnable`, highlights, indentation, backtick/tilde fences, and
  longer enclosing non-SQL fences are handled. Related SQL dialect labels are
  included; the exact language allowlist is stored in `manifest.json`.
- Standalone `.sql` files are included. The current snapshot's SQL `CodeBlock`
  components resolve to static JavaScript template maps; their variants are
  extracted without executing JavaScript. Unresolved SQL components or unclosed
  SQL fences cause extraction to fail rather than silently disappear.
- Every block is retained verbatim, including whitespace, comments, placeholders,
  syntax fragments, SQL-labeled responses, and foreign database dialects. Only
  byte-identical examples are deduplicated. SHA-256 filenames stay stable across
  documentation moves, and the manifest retains all duplicate source locations.
- Each block/file is passed to `ParseStmts` as a whole script. Semicolons inside
  strings, function bodies, comments, and INSERT payloads are not split by a
  heuristic. For a rejected multi-statement block, the error describes the first
  failure; subsequent statements are not independently classified.
- “Accepted” means this parser returned one or more statements without an error.
  It does not assert execution success, semantic correctness, or agreement with
  a live ClickHouse server. Rejections include intentionally incomplete grammar
  templates and expression-only examples, so these counts are not a percentage
  of supported ClickHouse features.
- Unmarked/inline SQL, SQL embedded in shell commands or other programming
  languages, and runtime-generated SQL are outside this explicit-marker corpus.
  The manifest records a broken source symlink (`_includes/content/changelog.md`)
  whose target is missing from the supplied copy.

The documentation SQL retains the source documentation's license; see
[LICENSE](LICENSE), copied from `clickhouse-docs/LICENSE`. Source attribution is
preserved in the manifest. The extractor and test are repository code.
