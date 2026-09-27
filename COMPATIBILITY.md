# Compatibility notes

This file lists every change to the public API of the `parser` package, so
that code consuming the AST can be updated when a release is cut. "Public
API" covers more than Go signatures. It also includes:

- exported types, fields, methods, functions and the `ASTVisitor` interface;
- the shape of the AST the parser returns for a given input;
- the meaning of positions (`Pos()`, `End()` and exported position fields);
- the SQL printed by `String()`, `PrintVisitor` and `BeautifyVisitor`;
- input that used to parse and is now rejected, or the reverse.

Every pull request that changes any of these adds an entry under
**Unreleased** and states the change as a WARNING in its description. When a
release is tagged, the Unreleased section is renamed to the release version.

Entries are grouped as:

- **Breaking**: code that compiled or behaved one way may no longer compile or
  may behave differently.
- **Changed behavior**: same types, different AST shape, positions or output.
- **Newly rejected input**: SQL that parsed before and is now an error.
  ClickHouse rejects all of it with a syntax error, unless noted otherwise.
- **Additive**: new fields, types or methods. Existing code keeps working.

## Unreleased (since v1.0.4)

### Breaking

- `ASTVisitor` has a new method, `VisitStreamClause(*StreamClause) error`
  (#102, #106). `StreamClause.Accept` now dispatches to it; it used to call
  only `Enter`/`Leave`. A visitor that embeds `DefaultASTVisitor` or
  `PrintVisitor` gets it for free; a type implementing `ASTVisitor` directly
  must add it.
- `ShowStmt.Format` changed type from `*StringLiteral` to `*FormatClause`
  (#87, #49). The field now comes from the embedded `OutputClauses`. The format
  name is `Format.Format.Name` (previously `Format.Literal`) and prints
  unquoted: `SHOW DATABASES FORMAT JSON` used to print as `FORMAT 'JSON'`, which
  ClickHouse rejects.
- Positions of quoted tokens and closing delimiters changed (#92):
  - `StringLiteral.Pos()`/`End()` and `Ident.Pos()`/`End()` of a quoted
    identifier (backticks, double quotes, or a string used as a name) now
    include the quotes. The fields `LiteralPos`/`LiteralEnd` and
    `NamePos`/`NameEnd` still hold the offsets of the content between the
    quotes. Code that slices the source with the fields is unaffected. Code
    that slices it with `Pos()`/`End()` now gets the quotes too.
  - `End()` of nodes closed by `)`, `]` or `}` (`ParamExprList`,
    `FunctionExpr`, `ColumnArgList`, `ArrayParamList`, `MapLiteral`,
    `TableSchemaClause`, `TableArgListExpr`, `ComplexType`, `NestedType`,
    `TypeWithParams`, `CompressionCodec`, `ColumnNamesExpr`,
    `AssignmentValues`, `WindowExpr`, `IndexTypeKwargs`, `QueryParam`,
    `TypedPlaceholder`, `ProjectionSelectStmt`, `TableProjection`) now points
    after the delimiter instead of at it.
  - Statement `End()` values were corrected for many statements. They stopped
    before a trailing quote, `)`, `DESC`, `FINAL`, `COMMENT`, … or included the
    `;`. A statement's `[Pos(), End())` now covers exactly the statement text.
  - Some existing position fields hold corrected values:
    - `CompressionCodec.RightParenPos` is the position of the `)`, as for every
      other `RightParenPos`. It was the end of the `)`.
    - `LiteralPos` of a column `COMMENT` is the string's position. It was the
      `COMMENT` keyword's.
    - `TypedPlaceholder.RightBracePos` is the position of the `}`. It was the
      next token's.
    - `Fill.FillPos` is the `FILL` keyword. It was the token after it.
    - `IsNullExpr.IsPos`/`IsNotNullExpr.IsPos` is the `IS` keyword. It was the
      token after `NULL`. `Pos()` of both is now the start of the operand.
    - `ShowStmt.StatementEnd` and `GrantPrivilegeStmt.StatementEnd` no longer
      point past the statement.
- `ATTACH` and `DETACH` statements keep their verb (#68, #43, #52):
  - `String()`, `PrintVisitor` and `BeautifyVisitor` print `ATTACH …` and
    `DETACH …`. `DETACH TABLE t` used to print as `DROP TABLE t`.
  - `DropStmt.Type()` returns `DETACH TABLE` etc. for a DETACH.
  - Use the new `IsAttach`/`IsDetach` fields to tell them apart from CREATE and
    DROP.
- TimeSeries targets (#42): all parts of one target (table, `INNER UUID`,
  `INNER COLUMNS`, `INNER ENGINE`) are collected into a single
  `TimeSeriesTargetClause`, ordered by the first occurrence of each target.
  Before, each occurrence of a target keyword was its own clause.
- `JoinExpr.Modifiers` starts with `GLOBAL` for a `GLOBAL … JOIN` (#74). The
  parser used to drop it.
- `a LOCAL JOIN b` parses `LOCAL` as the alias of `a`, as ClickHouse does
  (#74). It used to be dropped as a join keyword.
- `FROM t STREAM` parses into `TableExpr.Stream` (#81). `STREAM` used to become
  the implicit alias `AS STREAM`.
- A parenthesised UNION/EXCEPT operand is a **group node** (#79, #93):
  a `SelectQuery` with `HasParen` set, the query inside the parentheses in the
  new `Group` field, no SELECT list of its own, and whatever follows the
  closing `)` in its `UnionAll`/`UnionDistinct`/`Except`. Its `SelectPos` and
  `End()` span the parentheses. `HasParen` is set exactly when `Group` is.
  Code that looked for SELECT items on a `HasParen` node must read them from
  `Group`. The #79 shape (the operand's first SELECT with `HasParen` and its
  inner chain in its own union fields) was never released.
- Statements ending with `FORMAT …` and/or a query-level `SETTINGS …` keep
  those clauses in the embedded `OutputClauses` (#87). The parser used to
  silently drop `FORMAT` on every non-SELECT statement. A `SETTINGS` after
  `FORMAT` on a SELECT goes to the new `SelectQuery.OutputSettings`, separate
  from `SelectQuery.Settings`.

### Changed behavior (printed SQL)

- `PrintVisitor` applied to a subtree now prints the same SQL as the node's
  `String()` for every node type (#98). Before, it crashed or printed wrong SQL
  for some nodes:
  - **crashes**: a privilege with a column list such as `SELECT(x, y)` recursed
    until the stack overflowed; a window-frame `BETWEEN … AND …` and
    `GROUP BY ALL` dereferenced nil;
  - **dropped clauses**: `INTERPOLATE`, and `SETTINGS` on `ADD COLUMN`;
  - **wrong SQL**: `PRECEDING UNBOUNDED`, doubled window-frame parameters and
    `EXTRACT` operands, `MODIFY TTL TTL`, missing spaces in `INDEX`, `INTERVAL`
    added to `1 HOUR`, single-quoted names printed unquoted, and a different
    ENGINE clause order.

  Printing whole statements (the CLI `-format` path) is unchanged.
- `IntervalFrom.Accept` and `WindowFrameParam.Accept` no longer visit their
  child before calling the visitor (#98). A visitor embedding
  `DefaultASTVisitor` used to see an `EXTRACT(… FROM x)` operand twice. It now
  sees it once. `DefaultASTVisitor.VisitWindowFrameParam` visits the parameter,
  so a custom visitor that overrides it no longer gets the parameter visited
  first.

- `DefaultASTVisitor` now recurses into every child that `Walk` reaches, so a
  visitor embedding it sees them (#106). It used to skip the `DISTINCT ON`
  columns, the `ORDER BY` columns of a projection, the string of a `UUID`, the
  name of `CREATE DATABASE`, the `COMMENT` of `CREATE TABLE` and the
  `SETTINGS` of `ALTER … DROP PARTITION`. Both `Walk` and `DefaultASTVisitor`
  now visit the `TargetPair` nodes of `RENAME` (they used to jump to the
  tables inside) and the `StreamClause` of a table expression.
- `Walk`, and so `Find`, `FindAll`, `WalkWithBreak` and `Transform`, now
  visits every node that `Accept` with `DefaultASTVisitor` reaches (#105). It
  used to skip the operands of `EXCEPT`, the rows of `INSERT … VALUES`, TTL
  actions (`DELETE`, `RECOMPRESS CODEC(…)`), the name, `ON CLUSTER` and
  parameters of `CREATE NAMED COLLECTION`, and query parameters in window
  frames. Code that counts or collects nodes, for example the tables of a
  query, now sees them. Map literal keys are now walked (and visited) in
  place: rewriting the `*StringLiteral` of a key changes the AST, where it
  used to change a copy.
- `BeautifyVisitor` emits every statement type (#51). Before, it produced an
  empty string for SET, USE, SHOW, DESCRIBE, DROP, TRUNCATE, RENAME, OPTIMIZE,
  CHECK, SYSTEM, GRANT, DELETE, CREATE/ALTER ROLE, CREATE USER, CREATE
  DATABASE, CREATE DICTIONARY, CREATE FUNCTION and CREATE NAMED COLLECTION,
  and only the inner SELECT for CREATE LIVE VIEW and EXPLAIN.
- One-element tuples keep their trailing comma: `(3,)` no longer prints as
  `(3)` (#70). `ColumnExprList.HasTrailingComma` records it.
- `f(params)(DISTINCT x)` keeps `DISTINCT` (#71).
- `INSERT INTO [TABLE] FUNCTION f(…)` keeps `FUNCTION` (#72).
- `USING` always prints parenthesised, `USING (a, b)` (#74). Comma joins print
  `, ` instead of `,`.
- `PARTITION ID 'x'` keeps `ID` (#76).
- `SELECT … WITH TOTALS` without `GROUP BY` keeps `WITH TOTALS` (#77).
- A sign is printed against a numeric literal, `-1::Int32` rather than
  `- 1::Int32` (#83).
- `TOP n WITH TIES` prints `TOP n` (#90). `String()` used to print only
  `WITH TIES`.
- `PrintVisitor` keeps `PRECEDING`/`FOLLOWING` after an `INTERVAL` window-frame
  bound (#109). The CLI `-format` path printed
  `RANGE BETWEEN INTERVAL 1 DAY AND CURRENT ROW`, which ClickHouse reads as a
  different window.
- `SYSTEM SYNC REPLICA t` keeps `REPLICA` (#62). It used to print as
  `SYSTEM SYNC t`, which ClickHouse rejects. `SystemSyncExpr.End()` is now the
  end of the whole command (`SyncEnd`), not the end of the table.
- `SYSTEM START|STOP DISTRIBUTED SENDS` without a table no longer prints as
  `… DISTRIBUTED SENDS SENDS`, which ClickHouse reads as a table named `SENDS`
  (#108). `SystemCtrlExpr.Type` is now the whole target, e.g. `"MERGES"` or
  `"REPLICATION QUEUES"`, and `SystemCtrlExpr.Cluster` is nil when no table
  is given.
- `WINDOW` is printed after `HAVING`, where ClickHouse expects it (#88). It
  used to be printed right after `FROM`, so a query with `WINDOW` and `WHERE`,
  `GROUP BY` or `HAVING` printed as SQL that ClickHouse rejects. This applies
  to `String()`, `PrintVisitor` and `BeautifyVisitor`.
- `PrintVisitor.VisitOrderByExpr` prints `WITH FILL` (#47). It used to drop the
  clause when an ORDER BY element was printed directly; it now matches
  `OrderExpr.String()`.
- `CREATE TABLE t AS other` and `CREATE DATABASE db ENGINE = …` no longer
  contain a double space (#45, #68).
- `CREATE TABLE` keeps its second, query-level `SETTINGS` clause (#37). It
  used to overwrite the engine's storage `SETTINGS`.
- `CREATE USER`: repeated `HOST` and `SETTINGS` clauses are merged (#37). The
  last one used to win.

### Newly rejected input

- An empty column list: `CREATE TABLE t () …`, `CREATE VIEW v () …`,
  `… TO t ()`, TimeSeries `INNER COLUMNS ()` (#45).
- Repeated clauses:
  - engine clauses in `CREATE TABLE`: a second `ORDER BY`, `PARTITION BY`,
    `PRIMARY KEY`, `SAMPLE BY` or `TTL`, and a third `SETTINGS` (#37);
  - `CREATE USER` clauses: `IDENTIFIED`, `VALID UNTIL`, `DEFAULT ROLE`,
    `DEFAULT DATABASE`, `GRANTEES` (#37);
  - `GROUP BY` modifiers other than `[WITH ROLLUP | WITH CUBE] [WITH TOTALS]`
    (#37);
  - a repeated part of a TimeSeries target (#42).
- TimeSeries `RECENT` not followed by `SAMPLES` (#42).
- `ATTACH OR REPLACE`, `ATTACH`/`DETACH` of `FUNCTION`/`ROLE`/`USER`/`NAMED
  COLLECTION`, `DROP … PERMANENTLY`, and `CREATE DICTIONARY d` without a
  definition (#68).
- A trailing comma anywhere except a bare tuple: `f(a,)`, `[1,]`,
  `quantiles(0.5,)(x)`, `GROUP BY CUBE(a,)`, `ENGINE = MergeTree(x,)` (#70).
- `GLOBAL ARRAY JOIN`, and `GLOBAL` after the join kind, e.g.
  `LEFT GLOBAL JOIN` (#74).
- An alias, `FINAL` or `SAMPLE` after `STREAM` (#81).
- `FORMAT`/`SETTINGS` after statements ClickHouse accepts no output clauses
  for: GRANT, SYSTEM, USE, SET, DELETE, USER/ROLE statements, CREATE FUNCTION,
  CREATE NAMED COLLECTION (#87). The parser used to drop the clause silently.
- `SHOW … FORMAT 'name'` with a quoted format name (#87).
- `INSERT INTO t VALUES (…) FORMAT x` (#87). After `VALUES` everything is row
  data, and ClickHouse fails at runtime with `CANNOT_PARSE_INPUT_ASSERTION_FAILED`.
- A repeated `WITH TIES`, and `LIMIT n BY … WITH TIES` (#90).
- `SYSTEM START|STOP DISTRIBUTED MERGES|FETCHES|TTL MERGES`, which ClickHouse
  rejects (#108). They used to be printed as `… MERGES MERGES`, a different
  statement. Use `SYSTEM START|STOP MERGES|FETCHES|TTL MERGES`.

### Additive

- **New types**:
  - `OutputClauses`, embedded in `ShowStmt`, `DescribeStmt`, `CheckStmt`,
    `ExplainStmt`, `CreateTable`, `CreateView`, `CreateMaterializedView`,
    `CreateLiveView`, `CreateDictionary`, `CreateDatabase`, `AlterTable`,
    `DropStmt`, `DropDatabase`, `TruncateTable`, `RenameStmt` and
    `OptimizeStmt` (#87);
  - `StreamClause` (#81).
- **New fields**:
  - `CreateTable.Settings` (#37);
  - `TimeSeriesTargetClause.InnerUUID`, `EngineShorthand` (#42);
  - `IsAttach` on `CreateTable`, `CreateView`, `CreateMaterializedView`,
    `CreateLiveView`, `CreateDictionary`, `CreateDatabase`; `IsDetach` and
    `Permanently` on `DropStmt` and `DropDatabase`; `DropDatabase.Modifier`
    (#68);
  - `ColumnExprList.HasTrailingComma` (#70);
  - `AlterTableDelete.InPartition` (#76);
  - `SelectQuery.HasParen` (#79), `SelectQuery.Group` (#93) and
    `SelectQuery.OutputSettings` (#87);
  - `TableExpr.Stream` (#81);
  - `LimitClause.WithTies`, `WithTiesEnd` (#90);
  - `OrderExpr.Nulls` (`"FIRST"`/`"LAST"`) and `OrderExpr.Collate` (#47);
  - `WithTimeoutClause.WithTimeoutEnd` (#37);
  - `SystemSyncExpr.SyncEnd`, `Target`, `OnCluster`, `Database`,
    `IfExists`, `Mode`, `From`, `CacheName` (#62). `SystemSyncExpr.Cluster`
    is unchanged and still holds the table of `SYNC REPLICA`.
  - `SystemCtrlExpr.OnCluster` (#108);
  - position fields for the #92 span fixes: `CastExpr.RightParenPos`,
    `SubQuery.RightParenPos`, `UsingClause.UsingEnd`, `OrderExpr.OrderEnd`,
    `NamedCollectionParam.ParamEnd`, `IsNullExpr.NullEnd`,
    `IsNotNullExpr.NullEnd`.
- **New methods**:
  - `VisitStreamClause` on `DefaultASTVisitor` and `PrintVisitor` (#102);
  - `InsertStmt.IsTableFunction()` (#72);
  - `BeautifyVisitor` overrides for the statement types listed above (#51).
- **Newly accepted input**:
  - `CREATE TABLE … SETTINGS … SETTINGS …` (storage plus query level) (#37);
  - TimeSeries `RECENT SAMPLES` / `INNER UUID` / `INNER ENGINE` (#42);
  - `DETACH … PERMANENTLY`, `ATTACH DICTIONARY d`, `DROP/DETACH DATABASE …
    SYNC` (#68);
  - `GLOBAL <kind> JOIN` (#74);
  - `DROP [DETACHED] PARTITION ID`, `DELETE IN PARTITION` (#76);
  - `FROM t STREAM [BOUNDED] [UNORDERED]` (#81);
  - `FORMAT … SETTINGS` and a trailing `SETTINGS` on non-SELECT statements
    (#87);
  - `LIMIT n [OFFSET m] WITH TIES` (#90);
  - ORDER BY elements with `NULLS FIRST|LAST` and `COLLATE 'locale'`, in
    `SELECT` and window specifications (#47);
  - a UNION/EXCEPT that continues after a parenthesised operand, e.g.
    `(SELECT 1) UNION ALL SELECT 2`, a statement starting with `(`,
    `CREATE VIEW … AS (query) UNION …`, `INSERT INTO t (SELECT …) UNION …` and
    `INSERT INTO t WITH … SELECT` (#93).
  - `SYSTEM SYNC REPLICA [ON CLUSTER c] t [IF EXISTS] [STRICT | LIGHTWEIGHT
    [FROM 'r', …] | PULL]`, `SYSTEM SYNC DATABASE REPLICA`, `SYSTEM SYNC
    TRANSACTION LOG`, `SYSTEM SYNC FILE CACHE` and `SYSTEM SYNC FILESYSTEM
    CACHE ['name']`, each with `ON CLUSTER` (#62).
  - `SYSTEM START|STOP` with the targets `MERGES`, `TTL MERGES`, `MOVES`,
    `FETCHES`, `REPLICATED SENDS`, `REPLICATION QUEUES`, `DISTRIBUTED SENDS`,
    `PULLING REPLICATION LOG`, `CLEANUP`, `REDUCE BLOCKING PARTS` and
    `VIRTUAL PARTS UPDATE`, each with an optional `ON CLUSTER` and table, and
    `VIEWS`, `VIEW v` and `REPLICATED VIEW v` (#108).
