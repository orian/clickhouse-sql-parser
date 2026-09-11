select event_time
    ,type
    ,user
    ,auth_type
    ,client_address 
FROM clusterAllReplicas('default',system.session_log) 
WHERE type='LoginFailure' 
LIMIT 100
