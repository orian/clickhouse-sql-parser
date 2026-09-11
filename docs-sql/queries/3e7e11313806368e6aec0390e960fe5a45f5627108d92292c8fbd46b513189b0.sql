SELECT * FROM clusterAllReplicas('default',system.kafka_consumers)
ORDER BY assignments.partition_id ASC;
