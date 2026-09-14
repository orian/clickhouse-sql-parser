
        SELECT countIf(key IS NOT NULL)
        FROM
        (
            SELECT materialize(CAST(NULL, 'Nullable(Int8)')) AS key
            FROM numbers(100000000)
        )
    