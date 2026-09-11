SELECT
    hostName() AS host,
    count() AS table_count
FROM clusterAllReplicas('default', merge(system, '^tables'))
WHERE database = 'default'
GROUP BY hostName()
ORDER BY table_count DESC
