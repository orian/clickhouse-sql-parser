
        CREATE TABLE bench
        ENGINE = AggregatingMergeTree()
        ORDER BY key
        SETTINGS index_granularity = 8192
        AS
        SELECT CAST(reinterpretAsString(number), 'SimpleAggregateFunction(any, String)') AS key
        FROM numbers_mt(5e6)
        SETTINGS max_insert_threads = 16
    