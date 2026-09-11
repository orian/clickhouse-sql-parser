
        SELECT count() FROM (SELECT toNullable(number) AS k FROM numbers_mt(1e8)) AS lhs
        INNER JOIN (SELECT toNullable(number) AS k FROM numbers_mt(1e3)) AS rhs USING (k)
        FORMAT Null
        SETTINGS max_threads=1, join_algorithm='hash'
    