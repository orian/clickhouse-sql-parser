CREATE TABLE default.`.inner_id.samples.xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
(
    `id` Tuple(UInt64, UUID),
    `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
ENGINE = MergeTree
ORDER BY (id, timestamp)
SETTINGS index_granularity = 32768
