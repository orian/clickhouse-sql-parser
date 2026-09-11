
        INSERT INTO tuple_wide SELECT number, intDiv(number, 2048) % 2,
            tuple(randomString(1000), randomString(1000), randomString(1000), randomString(1000), randomString(1000))
        FROM numbers(150000)
    