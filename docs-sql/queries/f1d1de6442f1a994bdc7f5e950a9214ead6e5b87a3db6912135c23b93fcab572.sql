CREATE TABLE page_performance
(
    page_id UInt32,
    page_name String,
    avg_response_time AggregateFunction(avg, UInt32)  -- Armazena o estado necessário para o cálculo de avg
)
ENGINE = AggregatingMergeTree()
ORDER BY page_id;
