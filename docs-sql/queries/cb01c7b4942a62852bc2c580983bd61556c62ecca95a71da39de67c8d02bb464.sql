CREATE TABLE raw_page_views
(
    page_id UInt32,
    page_name String,
    response_time_ms UInt32,  -- ページの応答時間（ミリ秒）
    viewed_at DateTime DEFAULT now()
)
ENGINE = MergeTree()
ORDER BY (page_id, viewed_at);
