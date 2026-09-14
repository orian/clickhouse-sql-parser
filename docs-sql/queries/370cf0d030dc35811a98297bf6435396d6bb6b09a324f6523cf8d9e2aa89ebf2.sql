CREATE TABLE page_performance
(
    page_id UInt32,
    page_name String,
    avg_response_time AggregateFunction(avg, UInt32)  -- avg 계산에 필요한 상태를 저장합니다
)
ENGINE = AggregatingMergeTree()
ORDER BY page_id;
