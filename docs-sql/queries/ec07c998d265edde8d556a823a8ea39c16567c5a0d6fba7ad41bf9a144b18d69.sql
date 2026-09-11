SELECT message
FROM clusterAllReplicas('default', system.text_log)
WHERE query_id = 'ad40c712-d25d-45c4-b1a1-a28ba8d4019c'
ORDER BY event_time_microseconds ASC
