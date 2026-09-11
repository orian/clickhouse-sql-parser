SELECT
    database,
    table,
    count() AS part_count,
    formatReadableSize(sum(bytes_on_disk)) AS total_size
FROM clusterAllReplicas(default, system.parts)
WHERE active = 1 AND database = 'default'
GROUP BY database, table
ORDER BY part_count DESC
