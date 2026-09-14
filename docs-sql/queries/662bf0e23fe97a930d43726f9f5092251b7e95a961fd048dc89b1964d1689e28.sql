CREATE TABLE logs
(
    `code` LowCardinality(String),
    `timestamp` DateTime64(3)
)
ENGINE = MergeTree
ORDER BY (code, toUnixTimestamp(timestamp))
