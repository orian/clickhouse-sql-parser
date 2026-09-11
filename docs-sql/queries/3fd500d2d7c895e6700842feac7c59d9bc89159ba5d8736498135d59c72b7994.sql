CREATE TABLE user_events (
  event_id UUID,
  user_id UInt32,
  event_type LowCardinality(String),
  event_time DateTime
) ENGINE = MergeTree
ORDER BY event_time;

INSERT INTO user_events
SELECT
  generateUUIDv4() AS event_id,
  rand() % 10000 AS user_id,
  arrayElement(['click','view','purchase'], toUInt32(rand()) % 3 + 1) AS event_type,
  now() - INTERVAL rand() % 3600*24 SECOND AS event_time
FROM numbers(1000000);
