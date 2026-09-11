
        INSERT INTO {table}
        SELECT
            number,
            rand64(0),
            rand64(1),
            rand64(2),
            rand64(3),
            rand64(4)
        FROM
            system.numbers_mt
        LIMIT 10000000
    