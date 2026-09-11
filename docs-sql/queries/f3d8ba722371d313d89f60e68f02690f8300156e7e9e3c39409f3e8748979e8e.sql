SELECT
    location_id,
    location_name,
    min_temp,  
    max_temp
FROM temperature_extremes
ORDER BY location_id;
