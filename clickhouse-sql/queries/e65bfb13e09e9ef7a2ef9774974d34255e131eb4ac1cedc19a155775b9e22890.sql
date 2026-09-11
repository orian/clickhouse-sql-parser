
        INSERT INTO {table}
            SELECT
                number % 1000000,
                number,
                number,
                number,
                number,
                number,
                if (number % 2 == 0, 1, 0)
            FROM
                system.numbers_mt
            LIMIT 10000000
    