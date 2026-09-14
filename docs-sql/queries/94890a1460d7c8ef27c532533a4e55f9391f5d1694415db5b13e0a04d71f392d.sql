-- Define in CREATE TABLE
CREATE TABLE events
(
  ts DateTime,
  user_id UInt64,
  value UInt32,
  INDEX ts_minmax ts TYPE minmax GRANULARITY 1
)
ENGINE=MergeTree
ORDER BY ts;

-- Or add later and materialize
ALTER TABLE events ADD INDEX ts_minmax ts TYPE minmax GRANULARITY 1;
ALTER TABLE events MATERIALIZE INDEX ts_minmax;

-- Query that benefits from the index
SELECT count() FROM events WHERE ts >= now() - 3600;

-- Verify usage
EXPLAIN indexes = 1
SELECT count() FROM events WHERE ts >= now() - 3600;
