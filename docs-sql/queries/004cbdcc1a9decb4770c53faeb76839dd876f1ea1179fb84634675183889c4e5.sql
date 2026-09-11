SELECT
   hostname(),
   *
FROM clusterAllReplicas('default', system.events)
WHERE event ILIKE '%ParallelReplicas%'
