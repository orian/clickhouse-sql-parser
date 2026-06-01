-- Origin SQL:
CREATE TABLE t
(
    `a` Int32,
    PROJECTION p (SELECT _part_offset ORDER BY a) WITH SETTINGS (index_granularity = 512)
)
ENGINE = MergeTree
ORDER BY a;

CREATE TABLE events
(
    `event_time` DateTime,
    `value` Float64,
    PROJECTION p2
    (
        SELECT
            toStartOfHour(event_time) AS hour,
            sum(value) AS total_value
        GROUP BY hour
    ) WITH SETTINGS (index_granularity = 4096, index_granularity_bytes = 1048576)
)
ENGINE = MergeTree()
ORDER BY event_time;


-- Format SQL:
CREATE TABLE t (`a` Int32, PROJECTION p (SELECT _part_offset ORDER BY a) WITH SETTINGS (index_granularity=512)) ENGINE = MergeTree ORDER BY a;
CREATE TABLE events (`event_time` DateTime, `value` Float64, PROJECTION p2 (SELECT toStartOfHour(event_time) AS hour, sum(value) AS total_value GROUP BY hour) WITH SETTINGS (index_granularity=4096, index_granularity_bytes=1048576)) ENGINE = MergeTree() ORDER BY event_time;
