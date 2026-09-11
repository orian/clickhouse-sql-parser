SET enable_json_ast_dialect = 1;
SET dialect = 'clickhouse_json';

-- Subsequent queries are parsed as JSON ASTs:
{"type":"SelectWithUnionQuery", ...}
