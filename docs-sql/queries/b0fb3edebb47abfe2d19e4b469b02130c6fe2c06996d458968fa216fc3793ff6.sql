SELECT
    location_id,
    location_name,
    min_temp,     -- 直接访问 SimpleAggregateFunction 的值
    max_temp      -- 使用 SimpleAggregateFunction 无需调用终结函数
FROM temperature_extremes
ORDER BY location_id;
