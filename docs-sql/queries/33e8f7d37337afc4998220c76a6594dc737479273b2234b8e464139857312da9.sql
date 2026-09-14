SYSTEM FLUSH LOGS query_log;

SELECT query_id, type, query, query_duration_ms
FROM system.query_log
WHERE query_id IN (SELECT query_id FROM system.session_query_ids)
ORDER BY event_time_microseconds;
