SELECT
    normalizedQueryHash(query) AS query_hash,
    count() AS execution_count,
    any(query) AS example_query
FROM clusterAllReplicas(default, merge(system, '^query_log'))
WHERE event_date >= today() - 1
GROUP BY normalizedQueryHash(query)
ORDER BY execution_count DESC
LIMIT 20
