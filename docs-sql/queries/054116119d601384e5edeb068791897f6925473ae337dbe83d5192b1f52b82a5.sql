SELECT *
FROM system.query_log
WHERE type = 'QueryFinish'
  AND is_initial_query = 1
ORDER BY query_start_time DESC
LIMIT 1
FORMAT Vertical;
