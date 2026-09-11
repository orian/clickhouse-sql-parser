SELECT log_time, log_message
FROM logs
WHERE $__timeFilter(log_time)
