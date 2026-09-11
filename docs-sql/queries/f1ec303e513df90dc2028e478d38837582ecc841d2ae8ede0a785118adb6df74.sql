SELECT
    location_id,
    location_name,
    min(min_temp) AS min_temp,  -- 모든 파트에 대해 집계 
    max(max_temp) AS max_temp   -- 모든 파트에 대해 집계
FROM temperature_extremes
GROUP BY location_id, location_name
ORDER BY location_id;
