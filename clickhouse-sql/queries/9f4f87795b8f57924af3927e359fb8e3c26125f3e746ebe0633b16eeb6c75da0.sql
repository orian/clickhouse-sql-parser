
        SELECT count() FROM
            (SELECT key FROM phj_left LIMIT 1000000) AS l
        INNER JOIN
            (SELECT key FROM phj_left LIMIT 1000000) AS r
        USING (key)
        FORMAT Null
    