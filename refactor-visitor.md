# refactor-visitor — what this branch did and how to merge upstream into it

This document is written primarily for an agent (or human) merging upstream `AfterShip/clickhouse-sql-parser` `master` into this fork's `refactor-visitor` branch. It explains the structural invariants of the branch so you can translate upstream code into the branch's pattern mechanically.

## TL;DR

- **Upstream pattern:** each AST node's `Accept(visitor)` walks its own children, then calls `visitor.VisitX(self)`.
- **Branch pattern:** each AST node's `Accept(visitor)` is a uniform 3-liner. Child-walking lives in `DefaultASTVisitor.VisitX` and uses `visitor.Self` so embedding gives virtual-like dispatch.
- **Pretty-printing:** moved from `Expr.String()` to `PrintVisitor` (`parser/print_visitor.go`). `String()` methods are still present in `ast.go` and still work, but the CLI's `-format` path uses `PrintVisitor`.

## The invariants

### 1. `Expr.Accept` is uniform

Every `Accept` in `parser/ast.go` has this exact shape:

```go
func (x *XNode) Accept(visitor ASTVisitor) error {
    visitor.enter(x)
    defer visitor.leave(x)
    return visitor.VisitXNode(x)
}
```

No child-walking. No conditionals. If upstream adds an `Accept` that walks children, **move that walking into `DefaultASTVisitor.VisitXNode`** when you port it.

### 2. `DefaultASTVisitor.VisitX` does the recursion

`parser/default_visitor.go` owns the default traversal. Each method looks like:

```go
func (visitor DefaultASTVisitor) VisitXNode(x *XNode) error {
    visitor.enter(x)
    defer visitor.leave(x)
    if x.Child != nil {
        if err := x.Child.Accept(visitor.Self); err != nil {
            return err
        }
    }
    for _, item := range x.Items {
        if err := item.Accept(visitor.Self); err != nil {
            return err
        }
    }
    return nil
}
```

Note the recursion is through `visitor.Self`, not `visitor`. **This is critical** — see invariant 3.

### 3. The `Self ASTVisitor` field

```go
type DefaultASTVisitor struct {
    Self ASTVisitor
}
```

Go embedding is not virtual dispatch. If `PrintVisitor` embeds `DefaultASTVisitor` and the embedded `VisitX` calls `child.Accept(visitor)`, dispatch goes back into `DefaultASTVisitor.VisitChild`, never into `PrintVisitor.VisitChild` — so overrides would be silently bypassed.

The `Self` field is a forwarding pointer set by the concrete visitor:

```go
v := &PrintVisitor{...}
v.Self = v
```

Default traversals call `child.Accept(visitor.Self)`. Any custom visitor created in tests or library code **must** set `Self = v` (see `parser/visitor_test.go`).

### 4. The `ASTVisitor` interface

`parser/ast_visitor.go` is the full interface — one `VisitX` per AST node plus `enter` / `leave`. When upstream adds a new node type, you must add a `VisitX` entry here.

### 5. `PrintVisitor` is the canonical formatter

`parser/print_visitor.go` defines `PrintVisitor` which embeds `DefaultASTVisitor` and writes SQL into an internal `strings.Builder`. `NewPrintVisitor()` sets `Self`.

`main.go`'s `-format` path uses `NewPrintVisitor()` instead of `stmt.String()`. Node `String()` methods in `ast.go` still exist (tests rely on them, and a few `PrintVisitor.VisitX` methods call `.String()` on children as shorthand), but new printing logic should go in `PrintVisitor`.

### 6. `BeautifyVisitor` for indented/line-broken output (separate visitor, not a node method)

Upstream `master` introduced a `Formatter` type and a `FormatSQL(*Formatter)` method on every AST node (commit `9275c63` and the beautify follow-ups). That puts formatting back onto the nodes, which contradicts invariant 5. We **do not** adopt that approach.

Instead, beautify output lives in `parser/beautify_visitor.go` (`BeautifyVisitor`), a sibling of `PrintVisitor` that embeds `DefaultASTVisitor`. The CLI exposes it via `-beautify`:

```bash
$ clickhouse-sql-parser -beautify "SELECT a, b FROM t WHERE a > 1 GROUP BY a"
SELECT
  a,
  b
FROM t
WHERE a > 1
GROUP BY a
```

**Coverage:** top-level statement types (`SelectQuery`, `CreateTable`, `CreateView`, `CreateMaterializedView`, `AlterTable`, `InsertStmt`, `SubQuery`) have hand-written `VisitX` methods that break across lines and indent. Everything else falls through — child clauses are rendered via `child.String()` (compact). To beautify a specific subtree (e.g. add line breaks inside an `ALTER TABLE` clause, or inside a complex `WindowExpr`), add a corresponding `VisitX` method to `BeautifyVisitor` and recurse into nested statements via `child.Accept(b.Self)`.

Upstream's `parser/format.go` and `format/beautify/` golden fixtures are **deliberately skipped** during merges — see the conflict-hotspots note below.

## Checklist for merging upstream `master`

When upstream adds, renames, or modifies an AST node, do all the following:

1. **`parser/ast.go`**
   - Keep `Pos()`, `End()`, `String()` from upstream as-is (or merge per usual).
   - Replace the upstream `Accept` body with the uniform 3-liner. Save the upstream child-walking code — you'll need it in step 2.
2. **`parser/default_visitor.go`**
   - Add (or update) `DefaultASTVisitor.VisitX` containing the child-walking you saved.
   - Replace every `child.Accept(visitor)` with `child.Accept(visitor.Self)`.
   - Wrap the body with `visitor.enter(x)` / `defer visitor.leave(x)`.
3. **`parser/ast_visitor.go`**
   - Add the `VisitX(*XNode) error` method to the `ASTVisitor` interface.
4. **`parser/print_visitor.go`**
   - If upstream changed `Expr.String()`, mirror the change in the corresponding `PrintVisitor.VisitX`. If a new node was added, add a `PrintVisitor.VisitX` that emits its SQL (you can usually adapt the new `String()` body).
5. **`main.go`** — usually no changes needed; the CLI is small.
6. **Tests** — `parser/visitor_test.go` uses the `Self = v` pattern; any new test visitors must do the same.

### Conflict hotspots

- `parser/ast.go` — most likely place for upstream conflicts. Upstream will keep child-walking inside `Accept`; you must collapse to the 3-liner. Upstream may also try to remove `String()` methods (post `9275c63`) — keep them; we still rely on `String()` for `PrintVisitor` and node text in tests.
- `parser/default_visitor.go` — does not exist upstream. If upstream adds nodes, you must add `DefaultASTVisitor.VisitX` entries here that you wouldn't otherwise see in their diff.
- `parser/print_visitor.go` — does not exist upstream. Same situation.
- `parser/beautify_visitor.go` — does not exist upstream. New beautify behavior gets a `VisitX` method here, not a `FormatSQL` body on the node.
- `parser/ast_visitor.go` — upstream may add `VisitX` interface methods; merge those entries. Upstream's own `DefaultASTVisitor` block (callback-style, post `9275c63`) lives in this same file; drop everything past the interface closing brace and let our `parser/default_visitor.go` be the only implementation.
- `parser/format.go` — **delete on sight.** Upstream adds it in `9275c63` as the home for the `Formatter` API; we explicitly skip that refactor. The `parser/testdata/**/format/beautify/` golden directories should also be deleted when they appear.
- `main.go` — upstream's `-format` path uses `stmt.String()`. Keep the branch's `NewPrintVisitor()` version.

### Already cherry-picked from upstream

These fixes are already in the branch and should be skipped if you see them in a master-to-branch range:

- `601ddf7` ← `8bea76f` — #127 fix the expr might be empty in window frame.
- `a774291` ← `b083fa0` — #128 SELECT query as table function argument.
- `12b000d` ← `d03ad5b` — #130 panic with invalid SQL.

### Deliberately skipped upstream commits

These were intentionally **not** integrated because they conflict with the visitor architecture:

- `9275c63` — *Introduce `FormatSQL` API to allow customizing different formats.* Moves all `String()` methods to a new `parser/format.go` and adds `FormatSQL(*Formatter)` to every node. That re-puts formatting on the nodes (violating invariant 5). Replacement on this branch: `BeautifyVisitor` (invariant 6).
- `925428d`, `95a195f`, `2588297`, `50c0046`, `6d0175c`, `1cf1e3d`, `033e873`, `4b92d70`, `3b8db28`, `b2d2add`, `17e829d`, `bf5be12`, `a812844`, `ece3a11`, `089230e`, `74e0ba3`, `af5073c`, `461102b`, `5d7a71f`, `d468e13` — beautify-format improvements that depend on the `Formatter` API from `9275c63`. Port piecemeal into `BeautifyVisitor` if a specific feature is needed.

## Commit-by-commit walkthrough

In chronological order (`git log master..HEAD --reverse`):

1. `edd6e43` — move `DefaultASTVisitor` out of `parser/ast_visitor.go` into its own file.
2. `b150d60` — first cut at separating `String` and `Accept` into distinct visitors.
3. `0baceb2` — split print and default visitors into two files (`default_visitor.go`, `print_visitor.go`).
4. `046472f`, `d41158d` — bug-fix passes on the freshly extracted visitors.
5. `6bdeca3` — remove the child-walking bodies from every `Expr.Accept` in `ast.go` (-1232 / +139 lines).
6. `4037103` — add the `self` forwarding pointer ("had to implement .self hack").
7. `3cb039d` — switch `main.go -format` from `stmt.String()` to `NewPrintVisitor()`.
8. `d2ab646` — parser tweak: accept select-item alias without an `AS` keyword.
9. `6d15e58` — delete the leftover deprecated visitor file.
10. `9825d50` — rename `self` → `Self` (exported, so external visitors can set it).
11. `601ddf7`, `a774291`, `12b000d` — cherry-picked upstream fixes (#127, #128, #130).
12. `26b1c0c` — fix: `DefaultASTVisitor.VisitAlterTable` / `VisitCreateTable` now visit the table identifier (they had been missed during the refactor).

## Concrete examples

### `Accept` shape — `parser/ast.go`

```go
func (s *SelectItem) Accept(visitor ASTVisitor) error {
    visitor.enter(s)
    defer visitor.leave(s)
    return visitor.VisitSelectItem(s)
}
```

### Default traversal — `parser/default_visitor.go`

```go
func (visitor DefaultASTVisitor) VisitAlterTable(a *AlterTable) error {
    visitor.enter(a)
    defer visitor.leave(a)

    if a.TableIdentifier != nil {
        if err := a.TableIdentifier.Accept(visitor.Self); err != nil {
            return err
        }
    }
    if a.OnCluster != nil {
        if err := a.OnCluster.Accept(visitor.Self); err != nil {
            return err
        }
    }
    for _, expr := range a.AlterExprs {
        if err := expr.Accept(visitor.Self); err != nil {
            return err
        }
    }
    return nil
}
```

### Custom visitor — `parser/visitor_test.go`

```go
type simpleRewriteVisitor struct {
    DefaultASTVisitor
}

func (v *simpleRewriteVisitor) VisitTableIdentifier(expr *TableIdentifier) error {
    if err := v.DefaultASTVisitor.VisitTableIdentifier(expr); err != nil {
        return err
    }
    if expr.Table.String() == "group_by_all" {
        expr.Table = &Ident{Name: "hack"}
    }
    return nil
}

v := &simpleRewriteVisitor{}
v.Self = v                    // required
_ = stmt.Accept(v)
```

## Verification after merging upstream

1. `go build ./...` — must compile.
2. `go test ./parser/...` — all golden-file tests in `parser/testdata/{ddl,dml,query,basic}` must pass. `TestVisitor_Identical` exercises every fixture through `DefaultASTVisitor`, so a broken traversal will surface here.
3. `TestVisitor_SimpleRewrite` and `TestVisitor_NestRewrite` (`parser/visitor_test.go`) confirm the `Self` dispatch still works for custom visitors.
4. Spot-check formatting: build the CLI and run `./clickhouse-sql-parser -format "<a non-trivial query>"`. Output should match the corresponding `.golden.json` fixture's `Format SQL` block.
