
        CREATE TABLE IF NOT EXISTS interpret_table_08
        (
            `idColumnU64` UInt64,
            `dateColumn` Date,
            `aggCount_3` AggregateFunction(count),
            `aggCount_4` AggregateFunction(count)
        )
        ENGINE = AggregatingMergeTree()
        PARTITION BY toYYYYMM(dateColumn)
        ORDER BY (idColumnU64, dateColumn)
        TTL dateColumn + toIntervalDay(30)
        SETTINGS index_granularity = 1024, ttl_only_drop_parts = 1;
    