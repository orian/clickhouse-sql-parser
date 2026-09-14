SELECT
    elapsed,
    initial_user,
    client_name,
    hostname(),
    query_id,
    query
FROM clusterAllReplicas(default, system.processes)
ORDER BY elapsed DESC
