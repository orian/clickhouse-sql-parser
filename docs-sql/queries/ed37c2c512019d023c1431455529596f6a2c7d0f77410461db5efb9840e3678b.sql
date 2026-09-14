WITH error_rates AS (
  SELECT
    $__timeInterval(Timestamp) as ts,
    ServiceName,
    countIf (SpanKind = 'Server') as request_count,
    countIf (
      SpanKind = 'Server'
      and StatusCode = 'Error'
    ) as error_count,
    error_count / request_count * 100 AS error_percent
  FROM $__sourceTable
  WHERE $__timeFilter(Timestamp)
  GROUP BY ts, ServiceName
)
SELECT ts, ServiceName, error_percent
FROM error_rates
WHERE request_count > 10
