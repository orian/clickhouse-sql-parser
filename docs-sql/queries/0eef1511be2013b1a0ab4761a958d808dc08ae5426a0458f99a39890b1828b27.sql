SELECT count()
FROM cluster(cluster, view(SELECT * FROM t SETTINGS max_execution_time = 10));
