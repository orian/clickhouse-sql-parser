SELECT
    event_date,
    event_type,
    table,
    error,
    COUNT() AS error_count
FROM clusterAllReplicas(default, merge(system, '^part_log'))
WHERE database = 'default'
GROUP BY
    event_date,
    event_type,
    error,
    table
ORDER BY
    event_date DESC,
    error_count DESC
