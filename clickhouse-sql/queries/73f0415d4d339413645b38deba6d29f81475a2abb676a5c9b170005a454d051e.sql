
        SELECT avg(key)
        FROM
        (
            SELECT materialize(toNullable(1)) AS key
            FROM numbers(100000000)
        )
    