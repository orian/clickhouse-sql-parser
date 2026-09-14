
        SELECT if(number % 10 = 0, NULL, number % {cardinality}) AS k, number AS v
        FROM numbers_mt(10000000)
        LIMIT 10 BY k
        FORMAT NULL
    