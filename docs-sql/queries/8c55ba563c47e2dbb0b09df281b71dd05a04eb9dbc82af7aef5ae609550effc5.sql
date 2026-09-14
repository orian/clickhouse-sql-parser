SELECT
    location_id,
    location_name,
    min(min_temp) AS min_temp,  -- 聚合所有 parts 的数据
    max(max_temp) AS max_temp   -- 聚合所有 parts 的数据
FROM temperature_extremes
GROUP BY location_id, location_name
ORDER BY location_id;
