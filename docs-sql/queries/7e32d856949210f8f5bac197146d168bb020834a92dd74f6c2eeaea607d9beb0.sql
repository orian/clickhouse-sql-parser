SELECT
    location_id,
    location_name,
    min_temp,     -- Acessando diretamente os valores de SimpleAggregateFunction
    max_temp      -- Não é necessária uma função de finalização com SimpleAggregateFunction
FROM temperature_extremes
ORDER BY location_id;
