SELECT
    hostName() AS host,
    database,
    table,
    count() AS queue_size,
    sum(if(is_currently_executing = 1, 1, 0)) AS executing_count
FROM clusterAllReplicas(default, system.replication_queue)
GROUP BY hostName(), database, table
HAVING queue_size > 0
ORDER BY queue_size DESC
