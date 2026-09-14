
        INSERT INTO tuple_compact SELECT number,
            tuple(tuple(randomString(1000), randomString(1000), randomString(1000), randomString(1000), randomString(1000)))
        FROM numbers(140000)
    