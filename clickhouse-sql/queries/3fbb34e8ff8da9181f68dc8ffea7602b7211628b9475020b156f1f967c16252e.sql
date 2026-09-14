
        SELECT count() FROM (SELECT if(number % 7 = 0, NULL, number % 50) AS k, toUInt8(number % 3) AS cond FROM numbers_mt(2000000)) AS lhs
        LEFT JOIN (SELECT toNullable(number % 25) AS a, toNullable(number % 30) AS b FROM numbers(2000)) AS rhs
            ON (lhs.k = rhs.a AND lhs.cond) OR (lhs.k = rhs.b)
        FORMAT Null
        SETTINGS max_threads=8, join_algorithm='hash', join_use_nulls=1, max_joined_block_size_rows=1024
    