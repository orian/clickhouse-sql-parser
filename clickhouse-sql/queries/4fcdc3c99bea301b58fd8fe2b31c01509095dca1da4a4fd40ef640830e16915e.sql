
        SELECT count()
        FROM prop_perf_orders AS o
        INNER JOIN prop_perf_lineitem AS l ON o.k = l.k
        WHERE l.k % 1000 = 0
    