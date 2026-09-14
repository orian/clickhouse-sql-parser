CREATE MATERIALIZED VIEW temperature_extremes_mv
TO temperature_extremes
AS SELECT
    location_id,
    location_name,
    minSimpleState(temperature) AS min_temp,     -- SimpleState combinator 사용
    maxSimpleState(temperature) AS max_temp      -- SimpleState combinator 사용
FROM raw_temperature_readings
GROUP BY location_id, location_name;
