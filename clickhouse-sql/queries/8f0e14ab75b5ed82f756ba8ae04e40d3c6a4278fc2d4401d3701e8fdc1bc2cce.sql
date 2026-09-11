
        SELECT count()
        FROM
        (
            SELECT
                number + 1 AS a,
                intHash64(number) % 1000000000 + 1 AS b
            FROM numbers({n})
        ) AS l
        LEFT ANTI JOIN
        (
            SELECT
                if(number = 0, toUInt64({n}) + 1, toUInt64(0)) AS a,
                number % 1 AS b
            FROM numbers(1025)
        ) AS r
        ON l.a < r.a AND l.b > r.b
        FORMAT Null
    