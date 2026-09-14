CALL clickhouse_perform(
    'server',
    'CREATE TABLE remote_table (id Int32) ENGINE = MergeTree ORDER BY id'
);
