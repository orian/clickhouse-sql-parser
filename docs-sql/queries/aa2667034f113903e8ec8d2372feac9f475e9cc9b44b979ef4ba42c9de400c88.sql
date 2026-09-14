CREATE MATERIALIZED VIEW temperature_extremes_mv
TO temperature_extremes
AS SELECT
    location_id,
    location_name,
    minSimpleState(temperature) AS min_temp,     -- Uso del combinador SimpleState
    maxSimpleState(temperature) AS max_temp      -- Uso del combinador SimpleState
FROM raw_temperature_readings
GROUP BY location_id, location_name;
