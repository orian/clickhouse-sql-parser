
        SELECT number % {cardinality} AS user, number AS ts
        FROM numbers_mt(10000000)
        ORDER BY user ASC, ts DESC LIMIT 10 BY user
        FORMAT NULL
    