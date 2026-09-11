CREATE TABLE hackernews
(
    `id` Int32,
    `doc_id` Int32,
    `text` String,
    `vector` Array(Float32),
    `node_info` Tuple(
        start Nullable(UInt64),
        end Nullable(UInt64)),
    `metadata` String,
    `type` Enum8('story' = 1, 'comment' = 2, 'poll' = 3, 'pollopt' = 4, 'job' = 5),
    `by` LowCardinality(String),
    `time` DateTime,
    `title` String,
    `post_score` Int32,
    `dead` UInt8,
    `deleted` UInt8,
    `length` UInt32
)
ENGINE = MergeTree
ORDER BY id;
