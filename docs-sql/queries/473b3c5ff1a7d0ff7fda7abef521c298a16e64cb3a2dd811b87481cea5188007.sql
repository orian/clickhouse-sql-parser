SELECT
    query_start_time,
    query_duration_ms,
    query
FROM system.user_query_log
ORDER BY query_start_time DESC
LIMIT 10;
