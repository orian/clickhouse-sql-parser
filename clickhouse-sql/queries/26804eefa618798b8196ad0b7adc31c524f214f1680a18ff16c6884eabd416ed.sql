
        SELECT count()
        FROM (SELECT * FROM prop_perf_orders WHERE k BETWEEN 15000000 AND 15001000) AS o
        INNER JOIN prop_perf_lineitem AS l ON o.k = l.k
    