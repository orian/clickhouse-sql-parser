SELECT
    location_id,
    location_name,
    min(min_temp) AS min_temp,  -- すべてのパーツにわたって集計
    max(max_temp) AS max_temp   -- すべてのパーツにわたって集計
FROM temperature_extremes
GROUP BY location_id, location_name
ORDER BY location_id;
