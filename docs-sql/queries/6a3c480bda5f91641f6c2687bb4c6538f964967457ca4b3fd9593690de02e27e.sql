SELECT
    type,
    query_duration_ms,
    exception_code,
    exception
FROM system.query_log
WHERE query_id = '6b57dffd-8aac-4be5-b331-fa8b2e70227e'
  AND type IN ('QueryFinish', 'ExceptionBeforeStart', 'ExceptionWhileProcessing')
ORDER BY event_time_microseconds DESC
LIMIT 1;
