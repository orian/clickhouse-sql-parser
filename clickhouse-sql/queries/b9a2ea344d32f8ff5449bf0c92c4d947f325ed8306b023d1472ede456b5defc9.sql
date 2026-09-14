
        SELECT count()
        FROM (SELECT * FROM prop_perf_orders WHERE k = 15000000) AS o
        INNER JOIN prop_perf_lineitem AS l ON o.k = l.k
    