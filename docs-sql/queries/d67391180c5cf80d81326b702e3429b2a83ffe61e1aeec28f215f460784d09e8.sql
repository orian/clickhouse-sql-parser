SELECT clickhouse_raw_query(
    'CREATE TABLE t1 (x String) ENGINE = Memory',
    'host=localhost port=8123'
);
