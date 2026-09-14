
        SELECT if(number % 10 = 0, number % {cardinality}, 0) AS k, number AS v
        FROM numbers_mt(10000000)
        ORDER BY k, v LIMIT 10 BY k
        FORMAT NULL
    