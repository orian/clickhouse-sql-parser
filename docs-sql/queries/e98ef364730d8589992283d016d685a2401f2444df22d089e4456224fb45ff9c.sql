SELECT
  $__timeInterval(TimestampTime) AS ts,
  lower(SeverityText),
  count() AS count
FROM otel_logs
WHERE $__timeFilter(TimestampTime)
  AND lower(SeverityText) IN ('error', 'warn')
  AND $__filters
GROUP BY SeverityText, ts
ORDER BY ts ASC
