   SELECT
       query_id,
       query_duration_ms,
       read_rows,
       read_bytes,
       memory_usage
   FROM system.query_log
   WHERE type = 'QueryFinish'
     AND query_id = 'your-query-id'
   ORDER BY event_time_microseconds DESC
   LIMIT 1;
