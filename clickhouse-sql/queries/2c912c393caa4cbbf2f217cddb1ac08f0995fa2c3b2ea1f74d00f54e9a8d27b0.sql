
        SELECT number
        FROM numbers(10)
        WHERE (number IN
        (
            SELECT sum(number)
            FROM numbers(100000000)
            GROUP BY bitAnd(number, 15)
        )) OR (number IN
        (
            SELECT sum(number)
            FROM numbers(100000000)
            GROUP BY bitAnd(number, 17)
        ))
    