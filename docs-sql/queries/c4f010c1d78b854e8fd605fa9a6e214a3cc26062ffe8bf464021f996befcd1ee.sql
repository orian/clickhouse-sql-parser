SELECT
    event_date,
    count() AS total_count,
    sum(if(query LIKE '%async%', 1, 0)) AS async_count,
    sum(if(query LIKE '%INSERT%', 1, 0)) AS insert_count
FROM clusterAllReplicas(default, merge(system, '^query_log'))
WHERE event_date >= today() - 7
GROUP BY event_date
ORDER BY event_date DESC
