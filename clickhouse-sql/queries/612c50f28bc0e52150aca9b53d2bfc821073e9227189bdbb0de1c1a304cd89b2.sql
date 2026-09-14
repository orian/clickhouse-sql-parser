
        CREATE TABLE IF NOT EXISTS interpret_table_06
        (
            `idColumnU64` UInt64,
            `dateColumn` DateTime,
            `aggCount_3` AggregateFunction(count),
            `aggCount` AggregateFunction(count),
            `sagg_max_date` SimpleAggregateFunction(max, DateTime)
        )
        ENGINE = AggregatingMergeTree()
        PARTITION BY toYYYYMM(dateColumn)
        ORDER BY idColumnU64
        TTL dateColumn + toIntervalMonth(6)
        SETTINGS index_granularity = 1024, ttl_only_drop_parts = 1;
    