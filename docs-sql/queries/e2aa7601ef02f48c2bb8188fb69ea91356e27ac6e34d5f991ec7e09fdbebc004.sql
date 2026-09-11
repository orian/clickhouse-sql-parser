CREATE TABLE page_performance
(
    page_id UInt32,
    page_name String,
    avg_response_time AggregateFunction(avg, UInt32)  -- Almacena el estado necesario para el cálculo del promedio
)
ENGINE = AggregatingMergeTree()
ORDER BY page_id;
