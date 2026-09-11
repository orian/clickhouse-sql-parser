SELECT count(*)
FROM clusterAllReplicas('default', system.mutations)
WHERE is_done = 0;
