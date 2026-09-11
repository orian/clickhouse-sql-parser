CREATE TABLE hackernews_views (
    id UInt32,
    author String,
    views UInt64,
    sign Int8
)
ENGINE = CollapsingMergeTree(sign)
PRIMARY KEY (id, author)
