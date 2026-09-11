
        SELECT count() FROM
            (SELECT number AS key FROM numbers_mt(100)) AS l
        RIGHT JOIN
            (SELECT number AS key FROM numbers_mt(5000000)) AS r
        USING (key)
        FORMAT Null
    