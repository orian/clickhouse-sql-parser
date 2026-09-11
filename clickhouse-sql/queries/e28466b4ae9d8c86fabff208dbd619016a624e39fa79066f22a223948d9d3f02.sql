
        SELECT a, b, v FROM test_limit_by_{cardinality}
        ORDER BY a, b LIMIT 10 BY b
        FORMAT NULL
    