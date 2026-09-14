
        CREATE TABLE IF NOT EXISTS interpret_table_02
        (
            `idColumnU64` UInt64,
            `dateColumn` DateTime,
            `agg_uniq_u128_01` AggregateFunction(uniq, UInt128),
            `agg_uniq_u128_02` AggregateFunction(uniq, UInt128),
            `aggCount` AggregateFunction(count),
            `agg_uniq_u128_03` AggregateFunction(uniq, UInt128),
            `agg_uniq_u128_04` AggregateFunction(uniq, UInt128),
            `aggCount_3` AggregateFunction(count),
            `aggCount_4` AggregateFunction(count),
            `agg_topk_01` AggregateFunction(topKWeighted(2), UInt128, UInt64)
        )
        ENGINE = AggregatingMergeTree()
        PARTITION BY toYYYYMM(dateColumn)
        ORDER BY idColumnU64
        TTL dateColumn + toIntervalMonth(6)
        SETTINGS index_granularity = 1024, ttl_only_drop_parts = 1;
    