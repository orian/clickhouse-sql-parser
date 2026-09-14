SET enable_time_time64_type = 1;

CREATE TABLE events
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id;

SET enable_time_time64_type = 0;

-- 即使该设置现已禁用，此操作仍然有效。
-- 表已存在，因此插入操作不会被阻止。
INSERT INTO events VALUES (1, '14:30:25');

-- 但使用 Time 类型创建新表将会失败。
CREATE TABLE events_new
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id; -- ERR: TYPE_TIME_TIME64_IS_NOT_ENABLED
