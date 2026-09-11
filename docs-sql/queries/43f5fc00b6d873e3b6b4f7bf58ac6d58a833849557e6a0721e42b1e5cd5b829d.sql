CREATE MATERIALIZED VIEW temperature_extremes_mv
TO temperature_extremes
AS SELECT
    location_id,
    location_name,
    minSimpleState(temperature) AS min_temp,     -- 使用 SimpleState 组合器
    maxSimpleState(temperature) AS max_temp      -- 使用 SimpleState 组合器
FROM raw_temperature_readings
GROUP BY location_id, location_name;
