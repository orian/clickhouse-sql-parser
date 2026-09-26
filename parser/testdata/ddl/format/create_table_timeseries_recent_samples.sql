-- Origin SQL:
-- Introspected DDL from ClickHouse 26.8.11.7 (system.tables.create_table_query).
CREATE TABLE ts.custom (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS tags_to_columns = {'instance':'instance'}, recent_samples_ttl_seconds = 345600 SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity = 32768 RECENT SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1 TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `instance` String, `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree PRIMARY KEY metric_name ORDER BY tuple(metric_name, id) SETTINGS index_granularity = 8192, allow_dimensions_outside_sorting_key = 1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name;
CREATE TABLE ts.ext (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS recent_samples_ttl_seconds = 345600 SAMPLES ts.s TAGS INNER COLUMNS (`id` UUID DEFAULT reinterpretAsUUID(sipHash128(tags)), `metric_name` LowCardinality(String), `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree PRIMARY KEY metric_name ORDER BY tuple(metric_name, id) SETTINGS index_granularity = 8192, allow_dimensions_outside_sorting_key = 1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name RECENT SAMPLES INNER COLUMNS (`id` UUID, `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree PARTITION BY toStartOfInterval(toDateTime(timestamp), toIntervalHour(5)) ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;
CREATE TABLE ts.no_recent (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS recent_samples_ttl_seconds = 0 SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity = 32768 TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree PRIMARY KEY metric_name ORDER BY tuple(metric_name, id) SETTINGS index_granularity = 8192, allow_dimensions_outside_sorting_key = 1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name;
CREATE TABLE ts.plain (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS recent_samples_ttl_seconds = 345600 SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity = 32768 TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree PRIMARY KEY metric_name ORDER BY tuple(metric_name, id) SETTINGS index_granularity = 8192, allow_dimensions_outside_sorting_key = 1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name RECENT SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree PARTITION BY toStartOfInterval(toDateTime(timestamp), toIntervalHour(5)) ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1;
-- SHOW CREATE TABLE with show_table_uuid_in_table_create_query_if_not_nil = 1 (INNER UUID form).
CREATE TABLE ts.custom UUID '16ec4261-141c-46c1-8054-3ca5169dd169'
(
    `metric_name` String,
    `tags` Map(String, String),
    `time_series` Array(Tuple(
        DateTime64(3),
        Float64)),
    `metric_family` String,
    `type` String,
    `unit` String,
    `help` String
)
ENGINE = TimeSeries
SETTINGS tags_to_columns = {'instance':'instance'}, recent_samples_ttl_seconds = 345600
SAMPLES INNER UUID '11111111-1111-1111-1111-111111111111'
SAMPLES INNER COLUMNS
(
    `id` Tuple(
        UInt64,
        UUID),
    `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
SAMPLES INNER
ENGINE = MergeTree
ORDER BY (id, timestamp)
SETTINGS index_granularity = 32768
RECENT SAMPLES INNER UUID 'ba8a4b6b-9ca4-4462-a6cd-76897f587681'
RECENT SAMPLES INNER COLUMNS
(
    `id` Tuple(
        UInt64,
        UUID),
    `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
RECENT SAMPLES INNER
ENGINE = MergeTree
ORDER BY (id, timestamp)
TTL toDateTime(timestamp) + toIntervalSecond(345600)
SETTINGS index_granularity = 8192, ttl_only_drop_parts = 1
TAGS INNER UUID '6d397e50-015c-4dbb-bc10-6370e96446c6'
TAGS INNER COLUMNS
(
    `id` Tuple(
        UInt64,
        UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))),
    `metric_name` LowCardinality(String),
    `instance` String,
    `tags` Map(LowCardinality(String), String),
    `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))),
    `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))
)
TAGS INNER
ENGINE = AggregatingMergeTree
PRIMARY KEY metric_name
ORDER BY tuple(metric_name, id)
SETTINGS index_granularity = 8192, allow_dimensions_outside_sorting_key = 1
METRICS INNER UUID '0f11e5fa-a487-4edf-9194-e5289294101c'
METRICS INNER COLUMNS
(
    `metric_family_name` String,
    `type` LowCardinality(String),
    `unit` LowCardinality(String),
    `help` String
)
METRICS INNER
ENGINE = ReplacingMergeTree
ORDER BY metric_family_name
;
-- Hand-written forms accepted by ClickHouse 26.8 formatQuery.
CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES d.s TAGS d.t METRICS d.m RECENT SAMPLES d.r;
CREATE TABLE d.ts ENGINE = TimeSeries RECENT SAMPLES ENGINE = MergeTree ORDER BY id;
CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES INNER ENGINE = MergeTree ORDER BY id;
CREATE TABLE d.ts ENGINE = TimeSeries DATA d.s TAGS d.t SAMPLES INNER ENGINE = Memory;
CREATE TABLE d.ts ENGINE = TimeSeries recent samples d.r;


-- Format SQL:
CREATE TABLE ts.custom (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS tags_to_columns={'instance': 'instance'}, recent_samples_ttl_seconds=345600 SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity=32768 RECENT SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity=8192, ttl_only_drop_parts=1 TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `instance` String, `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree ORDER BY tuple(metric_name, id) PRIMARY KEY metric_name SETTINGS index_granularity=8192, allow_dimensions_outside_sorting_key=1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name;
CREATE TABLE ts.ext (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS recent_samples_ttl_seconds=345600 SAMPLES ts.s TAGS INNER COLUMNS (`id` UUID DEFAULT reinterpretAsUUID(sipHash128(tags)), `metric_name` LowCardinality(String), `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree ORDER BY tuple(metric_name, id) PRIMARY KEY metric_name SETTINGS index_granularity=8192, allow_dimensions_outside_sorting_key=1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name RECENT SAMPLES INNER COLUMNS (`id` UUID, `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) PARTITION BY toStartOfInterval(toDateTime(timestamp), toIntervalHour(5)) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity=8192, ttl_only_drop_parts=1;
CREATE TABLE ts.no_recent (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS recent_samples_ttl_seconds=0 SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity=32768 TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree ORDER BY tuple(metric_name, id) PRIMARY KEY metric_name SETTINGS index_granularity=8192, allow_dimensions_outside_sorting_key=1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name;
CREATE TABLE ts.plain (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS recent_samples_ttl_seconds=345600 SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity=32768 TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree ORDER BY tuple(metric_name, id) PRIMARY KEY metric_name SETTINGS index_granularity=8192, allow_dimensions_outside_sorting_key=1 METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name RECENT SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) PARTITION BY toStartOfInterval(toDateTime(timestamp), toIntervalHour(5)) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity=8192, ttl_only_drop_parts=1;
CREATE TABLE ts.custom UUID '16ec4261-141c-46c1-8054-3ca5169dd169' (`metric_name` String, `tags` Map(String, String), `time_series` Array(Tuple(DateTime64(3), Float64)), `metric_family` String, `type` String, `unit` String, `help` String) ENGINE = TimeSeries SETTINGS tags_to_columns={'instance': 'instance'}, recent_samples_ttl_seconds=345600 SAMPLES INNER UUID '11111111-1111-1111-1111-111111111111' SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) SETTINGS index_granularity=32768 RECENT SAMPLES INNER UUID 'ba8a4b6b-9ca4-4462-a6cd-76897f587681' RECENT SAMPLES INNER COLUMNS (`id` Tuple(UInt64, UUID), `timestamp` DateTime64(3) CODEC(DoubleDelta, ZSTD(1)), `value` Float64 CODEC(ZSTD(3))) RECENT SAMPLES INNER ENGINE = MergeTree ORDER BY (id, timestamp) TTL toDateTime(timestamp) + toIntervalSecond(345600) SETTINGS index_granularity=8192, ttl_only_drop_parts=1 TAGS INNER UUID '6d397e50-015c-4dbb-bc10-6370e96446c6' TAGS INNER COLUMNS (`id` Tuple(UInt64, UUID) DEFAULT tuple(sipHash64(metric_name), reinterpretAsUUID(sipHash128(tags))), `metric_name` LowCardinality(String), `instance` String, `tags` Map(LowCardinality(String), String), `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))), `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))) TAGS INNER ENGINE = AggregatingMergeTree ORDER BY tuple(metric_name, id) PRIMARY KEY metric_name SETTINGS index_granularity=8192, allow_dimensions_outside_sorting_key=1 METRICS INNER UUID '0f11e5fa-a487-4edf-9194-e5289294101c' METRICS INNER COLUMNS (`metric_family_name` String, `type` LowCardinality(String), `unit` LowCardinality(String), `help` String) METRICS INNER ENGINE = ReplacingMergeTree ORDER BY metric_family_name;
CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES d.s TAGS d.t METRICS d.m RECENT SAMPLES d.r;
CREATE TABLE d.ts ENGINE = TimeSeries RECENT SAMPLES ENGINE = MergeTree ORDER BY id;
CREATE TABLE d.ts ENGINE = TimeSeries SAMPLES INNER ENGINE = MergeTree ORDER BY id;
CREATE TABLE d.ts ENGINE = TimeSeries DATA d.s DATA INNER ENGINE = Memory TAGS d.t;
CREATE TABLE d.ts ENGINE = TimeSeries recent samples d.r;
