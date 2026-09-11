SELECT
  count() AS total_errors
FROM otel_logs
WHERE $__timeFilter(TimestampTime)
  AND SeverityText = 'error'
  AND $__filters
