SELECT
    request,
    count(*)
FROM error_log
GROUP BY request
