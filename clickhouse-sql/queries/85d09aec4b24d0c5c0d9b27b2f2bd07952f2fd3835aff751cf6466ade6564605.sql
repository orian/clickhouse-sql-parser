
        SELECT toString(number % {cardinality}) AS k1, toString(number % 97) AS k2, number AS v
        FROM numbers_mt(10000000)
        LIMIT 10 BY k1, k2
        FORMAT NULL
    