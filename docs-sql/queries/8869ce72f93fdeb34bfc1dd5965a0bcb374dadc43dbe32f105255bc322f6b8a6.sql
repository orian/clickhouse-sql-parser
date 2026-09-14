WITH
    initial_query_id = '9bc1b438-542e-450a-8284-0e3e2b3739c8' AS second,
    initial_query_id = '4ce13192-2849-48f9-b426-e04790c4f4aa' AS first
SELECT
    PE.1 AS metric,
    sumIf(PE.2, first) AS v1,
    sumIf(PE.2, second) AS v2,
    10 * log10(v2 / v1) AS dB,
    round(((v2 - v1) / if(v2 > v1, v2, v1)) * 100, 2) AS perc,
    bar(abs(perc), 0, 100, 33) AS bar
FROM clusterAllReplicas(default, system.query_log)
ARRAY JOIN ProfileEvents AS PE
WHERE (first OR second) AND (event_date >= (today() - 3)) AND (type = 2)
GROUP BY metric
HAVING (v1 != v2) AND (abs(perc) >= 0)
ORDER BY
    dB DESC,
    v2 DESC,
    metric ASC
FORMAT PrettyCompactMonoBlock
