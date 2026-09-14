SELECT
    location_id,
    location_name,
    min_temp,     -- Acceso directo a los valores de SimpleAggregateFunction
    max_temp      -- No se necesita función de finalización con SimpleAggregateFunction
FROM temperature_extremes
ORDER BY location_id;
