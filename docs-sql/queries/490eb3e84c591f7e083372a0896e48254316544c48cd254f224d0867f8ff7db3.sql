CREATE TABLE db1.tags_table
(
    `id` UInt64,
    `tags_string` Array(String),
    `tags_int` Array(UInt64),
)
ENGINE = MergeTree
ORDER BY id;
