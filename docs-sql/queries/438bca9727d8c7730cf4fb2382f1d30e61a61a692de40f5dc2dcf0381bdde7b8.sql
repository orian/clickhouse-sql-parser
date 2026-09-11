SELECT
  $__timeInterval(TimestampTime) AS ts,
  count() AS count
FROM otel_logs
WHERE $__timeFilter(TimestampTime)
  AND $__filter(ServiceName, $service)
  AND $__conditionalAll(SeverityText NOT IN ($severity), $severity)
GROUP BY ts
ORDER BY ts ASC
