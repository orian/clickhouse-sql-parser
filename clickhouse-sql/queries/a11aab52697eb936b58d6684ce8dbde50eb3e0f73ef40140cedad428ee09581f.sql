
        SELECT avgIf(key, key != -1)
        FROM
        (
            SELECT materialize(1) AS key
            FROM numbers(100000000)
        )
    