    CREATE TABLE events_rollup_1h
    (
        bucket_start  DateTime,            -- start of the hour
        country       LowCardinality(String),
        event_type    LowCardinality(String),
        users_uniq    AggregateFunction(uniqExact, UInt64),
        value_sum     AggregateFunction(sum, Float64),
        value_avg     AggregateFunction(avg, Float64),
        events_count  AggregateFunction(count)
    )
    ENGINE = AggregatingMergeTree
    PARTITION BY toYYYYMM(bucket_start)
    ORDER BY (bucket_start, country, event_type)
