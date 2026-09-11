CREATE TABLE default.`.inner_id.tags.xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
(
    `id` Tuple(UInt64, LowCardinality(UUID)) DEFAULT tuple(sipHash64(metric_name), toLowCardinality(reinterpretAsUUID(sipHash128(tags)))),
    `metric_name` LowCardinality(String),
    `tags` Map(LowCardinality(String), String),
    `min_time` SimpleAggregateFunction(min, Nullable(DateTime64(3))),
    `max_time` SimpleAggregateFunction(max, Nullable(DateTime64(3)))
)
ENGINE = AggregatingMergeTree
PRIMARY KEY metric_name
ORDER BY (metric_name, id)
SETTINGS allow_dimensions_outside_sorting_key = 1, index_granularity = 8192
