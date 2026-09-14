
        SELECT countIf(key IS NOT NULL)
        FROM
        (
            SELECT materialize(toNullable(1)) AS key
            FROM numbers(100000000)
        )
    