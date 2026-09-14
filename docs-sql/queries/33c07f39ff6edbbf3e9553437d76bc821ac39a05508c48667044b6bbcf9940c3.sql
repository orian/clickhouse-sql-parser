SELECT
  ServiceName,
  count() AS request_count
FROM otel_traces
WHERE $__timeFilter(Timestamp)
  AND $__filters
GROUP BY ServiceName
