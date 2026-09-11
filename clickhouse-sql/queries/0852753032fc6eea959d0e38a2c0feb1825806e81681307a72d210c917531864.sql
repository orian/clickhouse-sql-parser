
        SELECT sum(l.k)
        FROM (SELECT * FROM prop_perf_orders WHERE k BETWEEN 7000000 AND 7010000) AS o
        LEFT JOIN prop_perf_lineitem AS l ON o.k = l.k
    