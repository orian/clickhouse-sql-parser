
        SELECT count() FROM (SELECT toNullable(number) AS k FROM numbers_mt(2e8)) AS lhs
        LEFT JOIN (SELECT number*2 AS a, number*3 AS b FROM numbers_mt(1e3)) AS rhs ON lhs.k = rhs.a OR lhs.k = rhs.b
        FORMAT Null
        SETTINGS max_threads=8, join_algorithm='hash'
    