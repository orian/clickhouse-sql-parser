SELECT
    location_id,
    location_name,
    min_temp,     -- Прямой доступ к значениям SimpleAggregateFunction
    max_temp      -- Функция финализации для SimpleAggregateFunction не требуется
FROM temperature_extremes
ORDER BY location_id;
