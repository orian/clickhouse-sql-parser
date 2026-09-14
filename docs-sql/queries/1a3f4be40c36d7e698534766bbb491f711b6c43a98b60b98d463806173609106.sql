SELECT
    logger_name,
    count() AS c
FROM single_day_log
WHERE (thread_name = 'TCPHandler')
    AND (instance_type = 'm6i.4xlarge')
    AND hasToken(message, 'error')
GROUP BY logger_name
ORDER BY c DESC
LIMIT 5
