SELECT
    location_id,
    location_name,
    min(min_temp) AS min_temp,  -- Agregar en todas las partes 
    max(max_temp) AS max_temp   -- Agregar en todas las partes
FROM temperature_extremes
GROUP BY location_id, location_name
ORDER BY location_id;
