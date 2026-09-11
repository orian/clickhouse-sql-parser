CREATE TABLE source
(
    `message` String
)
ENGINE = MergeTree
ORDER BY tuple();

CREATE TABLE target
(
    `message` String,
    `from` String,
    `now` DateTime64(9),
    `sleep` UInt8
)
ENGINE = MergeTree
ORDER BY tuple();

CREATE MATERIALIZED VIEW mv_2 TO target
AS SELECT
    message,
    'mv2' AS from,
    now64(9) as now,
    sleep(1) as sleep
FROM source;

CREATE MATERIALIZED VIEW mv_3 TO target
AS SELECT
    message,
    'mv3' AS from,
    now64(9) as now,
    sleep(1) as sleep
FROM source;

CREATE MATERIALIZED VIEW mv_1 TO target
AS SELECT
    message,
    'mv1' AS from,
    now64(9) as now,
    sleep(1) as sleep
FROM source;
