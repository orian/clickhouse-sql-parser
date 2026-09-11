SELECT
    hostName() AS host,
    initial_user,
    query_id,
    elapsed,
    read_rows,
    formatReadableSize(memory_usage) AS memory_usage,
    normalizedQueryHash(query) AS query_hash
FROM clusterAllReplicas(default, system.processes)
ORDER BY elapsed DESC
