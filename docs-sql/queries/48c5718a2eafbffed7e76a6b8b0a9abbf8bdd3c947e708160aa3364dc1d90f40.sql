CREATE TABLE my_table
(
    `metric_name` String,
    `tags` Map(String, String),
    `time_series` Array(Tuple(DateTime64(3), Float64)),
    `metric_family` String,
    `type` String,
    `unit` String,
    `help` String
)
ENGINE = TimeSeries
SETTINGS version = 1, recent_samples_ttl_seconds = 345600
SAMPLES INNER COLUMNS
(
    `id` Tuple(UInt64, LowCardinality(UUID)),
    `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity = 32768
RECENT SAMPLES INNER COLUMNS
(
    `id` Tuple(UInt64, UUID),
    `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
RECENT SAMPLES INNER ENGINE = MergeTree PARTITION BY toStartOfInterval(toDateTime(timestamp), toIntervalHour(5)) ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1
TAGS INNER COLUMNS
(
    `id` Tuple(UInt64, LowCardinality(UUID)) DEFAULT tuple(sipHash64(metric_name), toLowCardinality(reinterpretAsUUID(sipHash128(tags)))),
    `metric_name` LowCardinality(String),
    `tags` Map(LowCardinality(String), String),
    `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))),
    `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))
)
TAGS INNER ENGINE = AggregatingMergeTree PRIMARY KEY metric_name ORDER BY (metric_name, id) SETTINGS allow_dimensions_outside_sorting_key = 1, index_granularity = 8192
METRICS INNER COLUMNS
(
    `metric_family_name` String,
    `type` LowCardinality(String),
    `unit` LowCardinality(String),
    `help` String
)
METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name
