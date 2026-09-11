
        SELECT number % {cardinality} AS grp, number AS v
        FROM numbers_mt(10000000)
        ORDER BY v LIMIT 10 BY grp % 64
        FORMAT NULL
    