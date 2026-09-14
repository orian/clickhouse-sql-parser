
        SELECT * FROM (SELECT toNullable(number) AS k FROM numbers_mt(5e8)) AS lhs
        LEFT JOIN (SELECT toNullable(number) AS k FROM numbers_mt(1e3)) AS rhs USING (k)
        FORMAT Null
        SETTINGS max_threads=8, join_algorithm='parallel_hash'
    