-- 在 CREATE TABLE 中定义
CREATE TABLE events
(
  ts DateTime,
  user_id UInt64,
  value UInt32,
  INDEX ts_minmax ts TYPE minmax GRANULARITY 1
)
ENGINE=MergeTree
ORDER BY ts;

-- 或稍后添加并物化
ALTER TABLE events ADD INDEX ts_minmax ts TYPE minmax GRANULARITY 1;
ALTER TABLE events MATERIALIZE INDEX ts_minmax;

-- 利用索引加速的查询
SELECT count() FROM events WHERE ts >= now() - 3600;

-- 验证索引使用情况
EXPLAIN indexes = 1
SELECT count() FROM events WHERE ts >= now() - 3600;
