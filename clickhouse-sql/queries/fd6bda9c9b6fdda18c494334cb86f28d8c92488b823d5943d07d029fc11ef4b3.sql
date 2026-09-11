
        SELECT count()
        FROM (SELECT * FROM prop_perf_orders WHERE k IN (10, 1000000, 7000000, 15000000, 22000000, 29999999)) AS o
        INNER JOIN prop_perf_lineitem AS l ON o.k = l.k
    