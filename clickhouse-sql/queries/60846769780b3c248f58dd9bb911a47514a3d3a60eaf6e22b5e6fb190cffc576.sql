
        SELECT a, b, v FROM test_limit_by_{cardinality}
        ORDER BY a, b, v + 1 LIMIT 10 BY a
        FORMAT NULL
    