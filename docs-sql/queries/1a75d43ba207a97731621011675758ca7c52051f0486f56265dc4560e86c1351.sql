SET enable_time_time64_type = 1;

CREATE TABLE events
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id;

SET enable_time_time64_type = 0;

-- 設定が無効になっていても、これは動作する。
-- テーブルはすでに存在するため、挿入はブロックされない。
INSERT INTO events VALUES (1, '14:30:25');

-- ただし、Time型で新しいテーブルを作成しようとすると失敗する。
CREATE TABLE events_new
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id; -- ERR: TYPE_TIME_TIME64_IS_NOT_ENABLED
