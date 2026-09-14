CREATE TABLE events_raw
(
    event_time   DateTime,
    user_id      UInt64,
    country      LowCardinality(String),
    event_type   LowCardinality(String),
    value        Float64
)
ENGINE = MergeTree
PARTITION BY toYYYYMM(event_time)
ORDER BY (event_time, user_id)
TTL event_time + INTERVAL 90 DAY DELETE
