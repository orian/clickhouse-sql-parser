SET enable_time_time64_type = 1;

CREATE TABLE time_examples
(
    `event_id` UInt8,
    `basic_time` Time,
    `precise_time` Time64(3)
)
ENGINE = MergeTree
ORDER BY event_id;
