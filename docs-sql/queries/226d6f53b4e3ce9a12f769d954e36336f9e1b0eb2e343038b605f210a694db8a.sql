CREATE TABLE default.`.inner_id.recentsamples.xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
(
    `id` Tuple(UInt64, UUID),
    `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
ENGINE = MergeTree
PARTITION BY toStartOfInterval(toDateTime(timestamp), toIntervalHour(5))
ORDER BY (id, timestamp)
TTL toDateTime(timestamp) + toIntervalSecond(345600)
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1
