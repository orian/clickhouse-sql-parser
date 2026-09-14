SELECT
  $__timeInterval(TimestampTime) AS ts,
  ServiceName,
  count() AS count
FROM otel_logs
WHERE $__timeFilter(TimestampTime)
  AND $__filters
GROUP BY ServiceName, ts
ORDER BY ts ASC
