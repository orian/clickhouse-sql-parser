
        INSERT INTO jit_test
            SELECT
                number,
                number,
                number,
                number,
                number,
                number,
                number,
                number,
                number,
                number
            FROM
                system.numbers_mt
            LIMIT 200000000
    