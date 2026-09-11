    CREATE TABLE events_rollup_1d
    (
        bucket_start Date,
        country      LowCardinality(String),
        event_type   LowCardinality(String),
        users_uniq   AggregateFunction(uniqExact, UInt64),
        value_sum    AggregateFunction(sum, Float64),
        value_avg    AggregateFunction(avg, Float64),
        events_count AggregateFunction(count)
    )
    ENGINE = AggregatingMergeTree
    PARTITION BY toYYYYMM(bucket_start)
    ORDER BY (bucket_start, country, event_type);
