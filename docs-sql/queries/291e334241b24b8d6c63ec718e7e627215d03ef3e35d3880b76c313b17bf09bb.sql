CREATE TABLE raw_page_views
(
    page_id UInt32,
    page_name String,
    response_time_ms UInt32,  -- 페이지 응답 시간(밀리초)
    viewed_at DateTime DEFAULT now()
)
ENGINE = MergeTree()
ORDER BY (page_id, viewed_at);
