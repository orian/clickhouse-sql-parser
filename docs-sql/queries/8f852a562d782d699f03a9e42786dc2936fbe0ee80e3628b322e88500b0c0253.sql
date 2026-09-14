CREATE TABLE temperature_extremes
(
    location_id UInt32,
    location_name String,
    min_temp SimpleAggregateFunction(min, Int32),  -- 最低気温を格納
    max_temp SimpleAggregateFunction(max, Int32)   -- 最高気温を格納
)
ENGINE = AggregatingMergeTree()
ORDER BY location_id;
