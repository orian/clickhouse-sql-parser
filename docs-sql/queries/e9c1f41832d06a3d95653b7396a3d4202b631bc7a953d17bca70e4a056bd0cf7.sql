CREATE TABLE events
(
    `timestamp` DateTime,
    `service`   LowCardinality(String),
    `level`     Enum8('DEBUG' = 1, 'INFO' = 2, 'WARN' = 3, 'ERROR' = 4),
    `message`   String,
    `host`      LowCardinality(String),
    `duration_ms` UInt32
)
ENGINE = MergeTree
ORDER BY (service, timestamp)
