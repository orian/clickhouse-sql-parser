SELECT event_time, query_kind, query, exception_code, exception
FROM system.query_log
WHERE user = '{fivetran_user}'
ORDER BY event_time DESC
LIMIT 100;
