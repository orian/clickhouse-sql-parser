
        CREATE TABLE IF NOT EXISTS interpret_table_03
        (
            `idColumnU64` UInt64,
            `dateColumn` Date,
            `aggCount` AggregateFunction(count),
            `aggCount_2` AggregateFunction(count),
            `aggCount_2_shown` AggregateFunction(count),
            `minDate` SimpleAggregateFunction(min, Date),
            `maxDate` SimpleAggregateFunction(max, Date),
            `maxInt16` SimpleAggregateFunction(max, Int16),
            `minUInt16` SimpleAggregateFunction(min, UInt16),
            `minUInt16_2` SimpleAggregateFunction(min, UInt16),
            `aggCount_3` AggregateFunction(count),
            `aggCount_4` AggregateFunction(count)
        )
        ENGINE = AggregatingMergeTree()
        PARTITION BY toYYYYMM(dateColumn)
        ORDER BY (idColumnU64, dateColumn)
        TTL dateColumn + toIntervalDay(30)
        SETTINGS index_granularity = 1024, ttl_only_drop_parts = 1;
    