-- EPHEMERAL column defaults, including the function-call expression form
-- from https://github.com/orian/clickhouse-sql-parser/issues/1
CREATE TABLE t (
    properties String,
    properties_map_ephemeral Map(String, String) EPHEMERAL CAST(JSONExtractKeysAndValues(properties, 'String'), 'Map(String, String)'),
    hidden String EPHEMERAL 'default value',
    bare UInt32 EPHEMERAL,
    bare_with_comment UInt32 EPHEMERAL COMMENT 'no expression',
    properties_map Map(String, String) MATERIALIZED CAST(JSONExtractKeysAndValues(properties, 'String'), 'Map(String, String)')
) ENGINE = Memory;
