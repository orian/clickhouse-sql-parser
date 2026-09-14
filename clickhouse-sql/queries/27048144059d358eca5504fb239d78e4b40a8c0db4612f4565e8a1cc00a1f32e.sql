
        CREATE TABLE IF NOT EXISTS interpret_table_05
        (
            `idColumnU64` UInt64,
            `dateColumn` Date,
            `agg_uniq_u128_01` AggregateFunction(uniq, UInt128),
            `agg_uniq_u128_02` AggregateFunction(uniq, UInt128),
            `agg_uniq_u128_03` AggregateFunction(uniq, UInt128),
            `agg_uniq_u128_04` AggregateFunction(uniq, UInt128),
            `aggCount_3` AggregateFunction(count),
            `aggCount_4` AggregateFunction(count),
            `aggCount` AggregateFunction(count)
        )
        ENGINE = AggregatingMergeTree()
        PARTITION BY toYYYYMM(dateColumn)
        ORDER BY (idColumnU64, dateColumn)
        TTL dateColumn + toIntervalDay(30)
        SETTINGS index_granularity = 1024, ttl_only_drop_parts = 1;
    