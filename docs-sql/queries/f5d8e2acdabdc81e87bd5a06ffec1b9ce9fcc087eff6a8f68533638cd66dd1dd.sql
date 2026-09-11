CREATE TABLE temperature_extremes
(
    location_id UInt32,
    location_name String,
    min_temp SimpleAggregateFunction(min, Int32),  -- Armazena a temperatura mínima
    max_temp SimpleAggregateFunction(max, Int32)   -- Armazena a temperatura máxima
)
ENGINE = AggregatingMergeTree()
ORDER BY location_id;
