    DROP TABLE IF EXISTS hackernews;

    CREATE TABLE hackernews
    (
        `id` UInt32,
        `deleted` UInt8,
        `type` Enum('story' = 1, 'comment' = 2, 'poll' = 3, 'pollopt' = 4, 'job' = 5),
        `by` LowCardinality(String),
        `time` DateTime,
        `text` String,
        `dead` UInt8,
        `parent` UInt32,
        `poll` UInt32,
        `kids` Array(UInt32),
        `url` String,
        `score` Int32,
        `title` String,
        `parts` Array(UInt32),
        `descendants` Int32
    )
        ENGINE = MergeTree
    ORDER BY id
