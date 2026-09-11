SET predicate_statistics_sample_rate = 1;

SELECT *
FROM hits
WHERE URL LIKE '%/product/%' AND EventDate >= today() - 7
FORMAT Null;

SYSTEM FLUSH LOGS predicate_statistics_log;

SELECT
    query_id,
    predicate_expression,
    round(filter_selectivity, 3) AS step_selectivity,
    round(total_selectivity, 3) AS query_selectivity,
    index_names,
    index_selectivities
FROM system.predicate_statistics_log
WHERE table = 'hits'
ORDER BY event_time DESC
LIMIT 10;
