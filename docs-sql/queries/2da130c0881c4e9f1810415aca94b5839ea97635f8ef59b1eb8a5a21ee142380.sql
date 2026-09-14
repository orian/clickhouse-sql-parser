SELECT
    hostName() AS host,
    formatReadableSize(max(memory_usage)) AS peak_memory,
    formatReadableSize(avg(memory_usage)) AS avg_memory,
    formatReadableSize(min(memory_usage)) AS min_memory
FROM clusterAllReplicas(default, merge(system, '^query_log'))
WHERE event_date >= today() - 1
GROUP BY hostName()
ORDER BY peak_memory DESC
