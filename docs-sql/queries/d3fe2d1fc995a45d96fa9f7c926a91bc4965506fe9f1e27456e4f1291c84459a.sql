SELECT ProfileEvents['UserTimeMicroseconds'] + ProfileEvents['SystemTimeMicroseconds'] AS cpu_us,
       ProfileEvents['NetworkSendBytes'] AS sent_bytes,
       query_duration_ms
FROM system.query_log
WHERE query_id = 'your-query-id' AND type = 'QueryFinish';
