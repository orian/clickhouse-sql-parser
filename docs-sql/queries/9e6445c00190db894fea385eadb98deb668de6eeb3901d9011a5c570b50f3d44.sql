SELECT
    database,
    table,
    partition_id,
    name,
    reason,
    count()
FROM clusterAllReplicas(default, system.detached_parts)
GROUP BY database, table, partition_id, name, reason
ORDER BY database, table
