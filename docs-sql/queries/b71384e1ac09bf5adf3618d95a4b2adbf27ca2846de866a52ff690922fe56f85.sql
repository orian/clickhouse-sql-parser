WITH toUInt64(123456789) AS selected_hash
SELECT
    event_time,
    query_id,
    query_duration_ms,
    read_rows,
    read_bytes,
    memory_usage,
    query
FROM merge('system', '^query_log')
WHERE type = 'QueryFinish'
  AND is_initial_query = 1
  AND normalized_query_hash = selected_hash
  AND event_time >= now() - INTERVAL 1 HOUR
ORDER BY event_time DESC
LIMIT 10;
