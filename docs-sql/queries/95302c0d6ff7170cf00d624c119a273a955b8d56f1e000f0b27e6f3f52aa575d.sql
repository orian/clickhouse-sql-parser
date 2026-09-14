CREATE TABLE temperature_extremes
(
    location_id UInt32,
    location_name String,
    min_temp SimpleAggregateFunction(min, Int32),  -- 최저 온도 저장
    max_temp SimpleAggregateFunction(max, Int32)   -- 최고 온도 저장
)
ENGINE = AggregatingMergeTree()
ORDER BY location_id;
