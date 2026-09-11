CREATE TABLE page_performance
(
    page_id UInt32,
    page_name String,
    avg_response_time AggregateFunction(avg, UInt32)  -- avg計算に必要なstateを格納する
)
ENGINE = AggregatingMergeTree()
ORDER BY page_id;
