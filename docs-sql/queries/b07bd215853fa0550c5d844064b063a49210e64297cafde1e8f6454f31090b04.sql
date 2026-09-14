SELECT
    client_name,
    count(),
    query_kind,
    toStartOfMinute(event_time) AS event_time_m
FROM system.query_log
WHERE (type = 'QueryStart') AND (event_time > (now() - toIntervalMinute(10)))
GROUP BY
    event_time_m,
    client_name,
    query_kind
ORDER BY
    event_time_m DESC,
    count() ASC
