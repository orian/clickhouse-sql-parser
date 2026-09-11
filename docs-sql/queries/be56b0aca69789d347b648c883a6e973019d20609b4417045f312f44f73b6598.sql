SELECT * FROM clusterAllReplicas('cluster_name', 'system', 'mutations')
WHERE is_done = 0 AND table = 'tmp';

-- или

SELECT * FROM clusterAllReplicas('cluster_name', 'system.mutations')
WHERE is_done = 0 AND table = 'tmp';
