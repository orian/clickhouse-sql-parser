
        SELECT number % {cardinality} AS k, number AS v
        FROM numbers_mt(10000000)
        ORDER BY k, v LIMIT 10 BY k
        FORMAT NULL
    