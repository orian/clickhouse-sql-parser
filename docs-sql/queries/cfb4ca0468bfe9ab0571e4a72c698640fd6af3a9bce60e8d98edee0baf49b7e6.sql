CREATE TABLE page_performance
(
    page_id UInt32,
    page_name String,
    avg_response_time AggregateFunction(avg, UInt32)  -- 存储 avg 计算所需的状态
)
ENGINE = AggregatingMergeTree()
ORDER BY page_id;
