SELECT
    hostName() AS host,
    name,
    value
FROM clusterAllReplicas(default, system.settings)
WHERE changed = 1
ORDER BY hostName(), name
