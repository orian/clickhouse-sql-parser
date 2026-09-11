
        SELECT * FROM numbers_mt(5e8) AS lhs
        LEFT JOIN (SELECT number AS k FROM numbers_mt(1e3)) AS rhs ON lhs.number = rhs.k AND (lhs.number % 2 = 0)
        FORMAT Null
        SETTINGS max_threads=8, join_algorithm='hash'
    