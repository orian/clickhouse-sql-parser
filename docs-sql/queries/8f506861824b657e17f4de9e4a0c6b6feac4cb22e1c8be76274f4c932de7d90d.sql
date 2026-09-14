SELECT
    location_id,
    location_name,
    min(min_temp) AS min_temp,  -- Агрегация по всем частям 
    max(max_temp) AS max_temp   -- Агрегация по всем частям
FROM temperature_extremes
GROUP BY location_id, location_name
ORDER BY location_id;
