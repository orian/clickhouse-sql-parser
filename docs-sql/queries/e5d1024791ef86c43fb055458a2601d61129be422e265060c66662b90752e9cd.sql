SELECT
    location_id,
    location_name,
    min(min_temp) AS min_temp,  -- Agrega entre todas as partes 
    max(max_temp) AS max_temp   -- Agrega entre todas as partes
FROM temperature_extremes
GROUP BY location_id, location_name
ORDER BY location_id;
