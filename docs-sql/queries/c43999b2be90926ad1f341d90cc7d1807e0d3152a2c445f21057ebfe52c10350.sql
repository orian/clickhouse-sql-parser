CREATE TABLE temperature_extremes
(
    location_id UInt32,
    location_name String,
    min_temp SimpleAggregateFunction(min, Int32),  -- Хранит минимальную температуру
    max_temp SimpleAggregateFunction(max, Int32)   -- Хранит максимальную температуру
)
ENGINE = AggregatingMergeTree()
ORDER BY location_id;
