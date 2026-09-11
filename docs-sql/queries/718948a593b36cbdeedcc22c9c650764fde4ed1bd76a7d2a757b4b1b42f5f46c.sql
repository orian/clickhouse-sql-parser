-- CREATE TABLEで定義する
CREATE TABLE events
(
  ts DateTime,
  user_id UInt64,
  value UInt32,
  INDEX ts_minmax ts TYPE minmax GRANULARITY 1
)
ENGINE=MergeTree
ORDER BY ts;

-- または後から追加してマテリアライズする
ALTER TABLE events ADD INDEX ts_minmax ts TYPE minmax GRANULARITY 1;
ALTER TABLE events MATERIALIZE INDEX ts_minmax;

-- 索引を活用するクエリ
SELECT count() FROM events WHERE ts >= now() - 3600;

-- 使用状況を確認する
EXPLAIN indexes = 1
SELECT count() FROM events WHERE ts >= now() - 3600;
