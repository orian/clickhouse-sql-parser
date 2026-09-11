CREATE TABLE page_performance
(
    page_id UInt32,
    page_name String,
    avg_response_time AggregateFunction(avg, UInt32)  -- Хранит состояние, необходимое для вычисления avg
)
ENGINE = AggregatingMergeTree()
ORDER BY page_id;
