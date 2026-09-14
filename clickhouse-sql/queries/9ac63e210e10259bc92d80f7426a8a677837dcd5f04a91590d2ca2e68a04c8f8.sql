
        INSERT INTO points
        SELECT
            number AS id,
            tuple(
                toInt32(toInt64(intHash64(number * 41) % 2001) - 1000),
                toInt32(toInt64(intHash64(number * 43) % 2001) - 1000)
            ) AS pt
        FROM numbers(500_000);
    