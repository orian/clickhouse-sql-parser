
        CREATE TABLE IF NOT EXISTS interpret_table_04
        (
            `idColumnU64` UInt64,
            `dateColumn` DateTime,
            `u128_id_02` UInt128,
            `ls_01` LowCardinality(String),
            `agg_count_01` AggregateFunction(count),
            `agg_count_02` AggregateFunction(count),
            `agg_smax_datetime_01` SimpleAggregateFunction(max, DateTime),
            `agg_smax_datetime_02` SimpleAggregateFunction(max, DateTime),
            `agg_count_03` AggregateFunction(count),
            `agg_count_04` AggregateFunction(count)
        )
        ENGINE = AggregatingMergeTree()
        PARTITION BY toYYYYMM(dateColumn)
        ORDER BY (idColumnU64, u128_id_02, ls_01)
        TTL dateColumn + toIntervalMonth(6)
        SETTINGS index_granularity = 1024, ttl_only_drop_parts = 1;
    