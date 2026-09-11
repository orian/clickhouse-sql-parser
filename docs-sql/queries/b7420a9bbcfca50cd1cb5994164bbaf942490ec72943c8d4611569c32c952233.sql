SELECT
    modification_time,
    rows,
    formatReadableSize(bytes_on_disk),
    *
FROM clusterAllReplicas(default, system.parts)
WHERE (database = 'default') AND active AND (level = 0)
ORDER BY modification_time DESC
LIMIT 100
