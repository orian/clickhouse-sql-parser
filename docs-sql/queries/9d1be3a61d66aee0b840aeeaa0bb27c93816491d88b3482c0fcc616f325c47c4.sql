-- CREATE TABLE에서 정의
CREATE TABLE events
(
  ts DateTime,
  user_id UInt64,
  value UInt32,
  INDEX ts_minmax ts TYPE minmax GRANULARITY 1
)
ENGINE=MergeTree
ORDER BY ts;

-- 또는 나중에 추가하고 구체화
ALTER TABLE events ADD INDEX ts_minmax ts TYPE minmax GRANULARITY 1;
ALTER TABLE events MATERIALIZE INDEX ts_minmax;

-- 인덱스를 활용하는 쿼리
SELECT count() FROM events WHERE ts >= now() - 3600;

-- 사용 여부 확인
EXPLAIN indexes = 1
SELECT count() FROM events WHERE ts >= now() - 3600;
