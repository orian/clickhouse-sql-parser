SELECT hostName(), query_id, elapsed, query
FROM clusterAllReplicas(my_cluster, system.processes)
WHERE query_id = '6b57dffd-8aac-4be5-b331-fa8b2e70227e';
