
        SELECT avgIf(key, key != -1)
        FROM
        (
            SELECT materialize(toNullable(1)) AS key
            FROM numbers(100000000)
        )
    