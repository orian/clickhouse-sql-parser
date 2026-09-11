SELECT event_time
    ,address
    ,initial_user
    ,initial_address
    ,forwarded_for
    ,query 
FROM clusterAllReplicas('default', system.query_log) 
WHERE user=’compromised_account’
