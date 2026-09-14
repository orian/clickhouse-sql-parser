SELECT query_id, query, query_duration_ms, read_rows, memory_usage
FROM system.query_log
WHERE query_id = '<query_id from run_results.json>'
  AND type = 'QueryFinish'
