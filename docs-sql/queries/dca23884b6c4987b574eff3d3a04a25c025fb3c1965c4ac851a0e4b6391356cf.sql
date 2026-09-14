-- 快速拉取特定 trace
SELECT *
FROM otel_traces
WHERE TraceId = 'aeea7f401feb75fc5af8eb25ebc8e974'
  AND Timestamp >= now() - INTERVAL 1 DAY
ORDER BY Timestamp;

-- Trace 范围内的聚合
SELECT
  toStartOfMinute(Timestamp) AS t,
  count() AS spans
FROM otel_traces
WHERE TraceId = 'aeea7f401feb75fc5af8eb25ebc8e974'
  AND Timestamp >= now() - INTERVAL 1 DAY
GROUP BY t
ORDER BY t;
