CREATE TABLE temperature_extremes
(
    location_id UInt32,
    location_name String,
    min_temp SimpleAggregateFunction(min, Int32),  -- 存储最低温度
    max_temp SimpleAggregateFunction(max, Int32)   -- 存储最高温度
)
ENGINE = AggregatingMergeTree()
ORDER BY location_id;
