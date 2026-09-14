
        SELECT count() FROM
            (SELECT number % 10 AS k FROM numbers_mt(100000)) AS l
        INNER JOIN
            (SELECT number % 10 AS k FROM numbers_mt(100000)) AS r
        USING (k)
        FORMAT Null
    