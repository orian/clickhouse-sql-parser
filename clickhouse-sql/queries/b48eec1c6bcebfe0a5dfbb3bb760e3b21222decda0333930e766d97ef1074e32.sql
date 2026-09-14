
        SELECT a, b, v FROM test_limit_by_{cardinality}
        ORDER BY a, b LIMIT 10 BY a
        FORMAT NULL
    