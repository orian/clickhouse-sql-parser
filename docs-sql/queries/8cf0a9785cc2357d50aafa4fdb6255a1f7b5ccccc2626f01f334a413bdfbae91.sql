SELECT
    normalized_query_hash,
    count() AS executions,
    quantile(0.5)(query_duration_ms) AS median_duration_ms,
    max(query_duration_ms) AS max_duration_ms,
    formatReadableSize(avg(read_bytes)) AS avg_read_bytes,
    formatReadableSize(max(memory_usage)) AS max_memory,
    any(query) AS example_query
FROM merge('system', '^query_log')
WHERE type = 'QueryFinish'
  AND is_initial_query = 1
  AND query_kind = 'Select'
  AND event_time >= now() - INTERVAL 1 HOUR
  AND has(databases, 'nyc_taxi')
GROUP BY normalized_query_hash
HAVING executions >= 2
ORDER BY median_duration_ms DESC
LIMIT 10
