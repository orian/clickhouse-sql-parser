SELECT
    hostname,
    type,
    query_id,
    initial_query_id,
    is_initial_query,
    query
FROM system.query_log
WHERE initial_query_id = '<query_id_of_initial_query>'
ORDER BY event_time_microseconds;
