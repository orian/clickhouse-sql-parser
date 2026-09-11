SELECT event_time, signal, signal_description
FROM system.crash_log
WHERE signal = 9 AND signal_description LIKE '%OOM Canary%'
ORDER BY event_time DESC;
