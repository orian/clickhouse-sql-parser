# Compatibility notes

This file lists every change to the public API of the `parser` package, so
that code consuming the AST can be updated when a release is cut. "Public
API" covers exported types, fields, methods and the `ASTVisitor` interface;
the AST shape returned for an input; `Pos()`/`End()` and position fields;
the SQL printed by `String()`, `PrintVisitor` and `BeautifyVisitor`; and
which input parses.

Every pull request that changes any of these adds an entry under
**Unreleased** and states the change as a WARNING in its description. When a
release is tagged, the Unreleased section is renamed to the release version.
Groups: **Breaking** (code may stop compiling or behave differently),
**Changed behavior** (same types, different AST, positions or output),
**Newly rejected input** (ClickHouse rejects all of it with a syntax error
unless noted) and **Additive**.

## Unreleased

## v1.1.0 (2026-09-28)

### Breaking

- **Go 1.27** is required (`go.mod` said `go 1.21.0`).
- **`ASTVisitor` has new methods**: `VisitStreamClause` (#102), `VisitCreateIndex`
  and `VisitDropIndex` (#135). Visitors embedding `DefaultASTVisitor` or
  `PrintVisitor` inherit them; direct implementations must add them.
  `StreamClause.Accept` now dispatches to `VisitStreamClause`.
- **Changed types**: `ShowStmt.Format` is `*FormatClause`, from the embedded
  `OutputClauses` (name in `Format.Format.Name`; was `*StringLiteral`), and
  prints unquoted, `FORMAT JSON` rather than `FORMAT 'JSON'` (#87, #49).
- **Fields that can now be nil, or hold other node types**:
  - `SubQuery.Select` is nil for `(EXPLAIN …)`; see `SubQuery.Explain` (#133);
  - `TableIndex.Granularity` is nil when `GRANULARITY` is omitted (#134);
  - `ExplainStmt.Statement` can be any statement, not only `*SelectQuery`, and
    is nil for `EXPLAIN CURRENT TRANSACTION`; `ExplainStmt.Type` is empty when
    the kind is omitted (#86);
  - `SystemCtrlExpr.Cluster` (the table) is nil when none is given (#108);
  - `TableArgListExpr.Args` can hold any expression (#129).
- **Positions** (#92):
  - `Pos()`/`End()` of a quoted `StringLiteral`/`Ident` include the quotes; the
    `LiteralPos`/`LiteralEnd`/`NamePos`/`NameEnd` fields still exclude them;
  - `End()` of nodes closed by `)`, `]` or `}` points after the delimiter
    (`ParamExprList`, `FunctionExpr`, `ArrayParamList`, `MapLiteral`,
    `TableSchemaClause`, `TableArgListExpr`, `ComplexType`, `QueryParam`, …);
  - a statement's `[Pos(), End())` covers exactly its text (it used to stop
    early or include the `;`);
  - corrected fields: `CompressionCodec.RightParenPos` (the `)`), column
    `COMMENT` `LiteralPos` (the string), `TypedPlaceholder.RightBracePos` (the
    `}`), `Fill.FillPos` (`FILL`), `IsNullExpr`/`IsNotNullExpr.IsPos` (`IS`;
    `Pos()` is the operand), `ShowStmt`/`GrantPrivilegeStmt.StatementEnd`;
  - `SystemSyncExpr.End()`, `SystemFlushExpr.End()`, `ExplainStmt.End()` and
    `TableIndex.End()` cover the whole node (#62, #114, #86, #134).
- **AST shapes**:
  - `ATTACH`/`DETACH` keep their verb (`IsAttach`/`IsDetach`; `DropStmt.Type()`
    returns `DETACH TABLE`); `DETACH TABLE t` used to print as `DROP TABLE t`
    (#43, #52, #68);
  - a parenthesised set-operation operand is a group node: a `SelectQuery` with
    `HasParen` and the inner query in `Group` (#79, #93);
  - trailing `FORMAT`/query-level `SETTINGS` of a statement live in the embedded
    `OutputClauses`; after `FORMAT` on a SELECT in `SelectQuery.OutputSettings`
    (#87);
  - all parts of one TimeSeries target form a single `TimeSeriesTargetClause`
    (#42);
  - `FROM t STREAM` is `TableExpr.Stream`, not the alias `STREAM` (#81);
  - `JoinExpr.Modifiers` keeps `GLOBAL`, and can hold `NATURAL`/`PASTE`;
    `a LOCAL JOIN b` aliases `a` as `LOCAL` (#74, #89);
  - lower-case `array join` is an ARRAY JOIN (`JoinExpr.Left` is a
    `*ColumnExprList`), not a join on a table named by its first expression
    (#50);
  - `SystemCtrlExpr.Type` is the whole target, e.g. `"REPLICATION QUEUES"`
    (#108).

### Changed behavior

- **Traversal**: `Walk` (and `Find`, `FindAll`, `WalkWithBreak`, `Transform`)
  and `DefaultASTVisitor` now reach every child node and agree with each other
  (#105, #106). They used to skip, among others, `EXCEPT` operands, `INSERT …
  VALUES` rows, TTL actions, named-collection parameters, `DISTINCT ON`
  columns, projection `ORDER BY` columns and `RENAME` pairs. Map keys are
  walked in place, so rewriting one changes the AST. `IntervalFrom` and
  `WindowFrameParam` no longer visit their child twice (#98).
- **`PrintVisitor`** prints every node like its `String()` (#98) and streams
  into one builder (#99). `BeautifyVisitor` emits every statement type (#51).
- **Printed SQL fixed** (the old output was invalid or meant something else):
  - `WINDOW` after `HAVING` (#88); `IF [NOT] EXISTS` before the name in `DROP
    INDEX|PROJECTION` and `ADD INDEX` (#64);
  - `PRECEDING`/`FOLLOWING` after an `INTERVAL` frame bound (#109);
  - `SYSTEM SYNC REPLICA` (#62), `SYSTEM FLUSH DISTRIBUTED` (#114), `SYSTEM
    STOP|START DISTRIBUTED SENDS` without `… SENDS SENDS` (#108);
  - `SELECT 1, limit(1)` keeps the column instead of becoming `LIMIT (1)`
    (#138); `PASTE`/`NATURAL JOIN` are joins, not aliases (#89);
  - `(3,)` keeps its comma (#70), `f(p)(DISTINCT x)` its `DISTINCT` (#71),
    `INSERT INTO FUNCTION` its `FUNCTION` (#72), `PARTITION ID` its `ID` (#76),
    `WITH TOTALS` without `GROUP BY` (#77), `WITH FILL` (#47);
  - `USING` is always parenthesised and comma joins print `, ` (#74); `-1::T`
    keeps the sign on the literal (#83); `TOP n WITH TIES` (#90); no double
    spaces in `CREATE TABLE t AS …`/`CREATE DATABASE … ENGINE` (#45, #68);
  - `CREATE TABLE` keeps its query-level `SETTINGS`, and repeated `CREATE USER`
    `HOST`/`SETTINGS` clauses are merged (#37).
- **Parsing**: `INTERSECT` is a keyword, not an implicit alias (#107);
  `SELECT 1 EXCEPT (SELECT 2)` is a set operation, not a column transformer
  (#107); in an expression `{` always starts a query parameter, and a
  `{'k': v}` map literal is valid only in `INSERT … VALUES` rows and as a
  `SETTINGS` value (#50).

### Newly rejected input

- **Select list and joins**: a missing select list (`SELECT`, `SELECT FROM t`;
  a lone `SELECT` used to parse to zero statements) (#82); a comma before a
  clause other than `FROM` (`SELECT a, ORDER BY a`) (#138); a `JOIN` without
  `ON`/`USING` except `CROSS`, `NATURAL`, `PASTE`, `ARRAY` (#50); `PASTE`/
  `NATURAL` as an implicit alias, or with `ON`/`USING`, a strictness, `CROSS`
  or `ARRAY` (#89); `GLOBAL ARRAY JOIN` and `LEFT GLOBAL JOIN` (#74).
- **Expressions**: column transformers (`EXCEPT`, `APPLY`, `REPLACE`) after
  anything but `*`, `t.*`, `COLUMNS(…)` or `* LIKE 'p'` (#50); `WITH CUBE/ROLLUP`
  after `GROUP BY CUBE(…)/ROLLUP(…)` (#50); a map literal in an expression
  (#50); a trailing comma except in a one-element tuple (`f(a,)`, `[1,]`)
  (#70).
- **Repeated or misplaced clauses**: repeated engine clauses in `CREATE TABLE`,
  repeated `CREATE USER` clauses, `GROUP BY` modifiers out of order (#37);
  repeated TimeSeries target parts and `RECENT` without `SAMPLES` (#42); an
  alias, `FINAL` or `SAMPLE` after `STREAM` (#81); repeated `WITH TIES` and
  `LIMIT … BY … WITH TIES` (#90); `FORMAT`/`SETTINGS` after statements that
  take none (GRANT, SYSTEM, USE, SET, DELETE, USER/ROLE, CREATE FUNCTION/NAMED
  COLLECTION; used to be dropped silently), a quoted `SHOW … FORMAT 'x'`, and
  `INSERT … VALUES (…) FORMAT x` (#87).
- **DDL**: an empty column list `()` (#45); `ATTACH OR REPLACE`,
  `ATTACH`/`DETACH` of FUNCTION/ROLE/USER/NAMED COLLECTION, `DROP …
  PERMANENTLY`, `CREATE DICTIONARY d` without a definition (#68); `OR REPLACE`
  with `IF NOT EXISTS` in `CREATE TABLE`/`VIEW` (#50); several pairs in `RENAME
  DATABASE` (#50); a `PARTITION` value other than a literal, query parameter,
  tuple or CAST of one (`DROP PARTITION p`, `PARTITION toDate(…)`) (#50); a
  malformed `UUID '…'` (ClickHouse: `CANNOT_PARSE_UUID`) (#50).
- **SYSTEM**: `START|STOP DISTRIBUTED MERGES|FETCHES|TTL MERGES` (used to print
  as `… MERGES MERGES`) (#108).
- **Access control** (#50): several roles in `ALTER ROLE … RENAME TO`; `ON
  CLUSTER` before the last role name; `WITH ADMIN OPTION` on a privilege grant;
  a column list on a wildcard target; the target `*.table`; `GRANT ADMIN
  OPTION ON …`.

### Additive

- **New types**: `OutputClauses` (#87), `StreamClause` (#81), `CreateIndex`,
  `DropIndex` (#135).
- **New fields**: `CreateTable.Settings`, `WithTimeoutClause.WithTimeoutEnd`
  (#37); `TimeSeriesTargetClause.InnerUUID`, `EngineShorthand` (#42);
  `IsAttach`/`IsDetach`/`Permanently`, `DropDatabase.Modifier` (#68);
  `ColumnExprList.HasTrailingComma` (#70); `AlterTableDelete.InPartition`
  (#76); `SelectQuery.HasParen`, `Group`, `OutputSettings`, `Union`,
  `Intersect`, `IntersectModifier`, `ExceptModifier` (#79, #87, #93, #107,
  #125); `TableExpr.Stream` (#81); `LimitClause.WithTies`, `WithTiesEnd` (#90);
  `OrderExpr.Nulls`, `Collate` (#47); `SystemSyncExpr.SyncEnd`, `Target`,
  `OnCluster`, `Database`, `IfExists`, `Mode`, `From`, `CacheName` (#62);
  `SystemCtrlExpr.OnCluster` (#108); `SystemFlushExpr.AsyncInsertQueue`,
  `OnCluster`, `Tables`, `Settings` (#114); `ExplainStmt.Settings`,
  `ExplainEnd` (#86); `Ident.Param` for `{name:Identifier}` (#131);
  `SubQuery.Explain` (#133); span fields `CastExpr.RightParenPos`,
  `SubQuery.RightParenPos`, `UsingClause.UsingEnd`, `OrderExpr.OrderEnd`,
  `NamedCollectionParam.ParamEnd`, `IsNullExpr`/`IsNotNullExpr.NullEnd` (#92).
- **New methods**: `VisitStreamClause`, `VisitCreateIndex`, `VisitDropIndex` on
  `DefaultASTVisitor` and `PrintVisitor`; `InsertStmt.IsTableFunction()` (#72).
- **Newly accepted input**:
  - SELECT: `INTERSECT [DISTINCT|ALL]`, `EXCEPT DISTINCT|ALL`, bare `UNION`
    (#107, #125); `(SELECT …) UNION …` and a statement starting with `(`
    (#93); `EXPLAIN` without a kind, with settings, of any statement, and as a
    subquery `FROM (EXPLAIN …)` (#86, #133); any expression as a
    table-function argument (#129); `{name:Identifier}` as a table, database or
    other name (#131); `PASTE JOIN`, `NATURAL … JOIN`, `GLOBAL <kind> JOIN`
    (#74, #89); `FROM t STREAM [BOUNDED] [UNORDERED]` (#81); `LIMIT n WITH
    TIES` (#90); `NULLS FIRST|LAST` and `COLLATE` in ORDER BY (#47); a
    keyword-named function after a comma (`SELECT 1, format(…)`) (#138); map
    literals in `INSERT … VALUES` (#50).
  - DDL: `CREATE [UNIQUE] INDEX` and `DROP INDEX` (#135); an index without
    `GRANULARITY` (#134); `CREATE TABLE … SETTINGS … SETTINGS …` (#37);
    TimeSeries `RECENT SAMPLES`, `INNER UUID`, `INNER ENGINE` (#42); `DETACH …
    PERMANENTLY`, `ATTACH DICTIONARY d`, `DROP|DETACH DATABASE … SYNC` (#68);
    `DROP [DETACHED] PARTITION ID`, `DELETE IN PARTITION` (#76); `FORMAT …
    SETTINGS` on non-SELECT statements (#87); a final single-token clause at
    end of input, e.g. `ALTER TABLE t FREEZE` (#46).
  - SYSTEM: every `SYNC` form (#62), `START|STOP` targets with `ON CLUSTER`
    and table (#108), `FLUSH LOGS [log, …]`, `FLUSH DISTRIBUTED … SETTINGS`,
    `FLUSH ASYNC INSERT QUEUE` (#114).
