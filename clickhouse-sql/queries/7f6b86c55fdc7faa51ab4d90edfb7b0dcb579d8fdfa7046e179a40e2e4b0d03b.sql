
        INSERT INTO na_src
        SELECT
            number,
            if(number % 3 = 0, NULL, number),
            if(number % 5 = 0, NULL, toUInt32(number)),
            number % 2,
            toDate(number % 40000),
            toDateTime64(number / 1000, 3),
            toDecimal64(number % 100000, 4),
            toString(number % 100000),
            toString(number % 1000),
            range(number % 5)
        FROM numbers(20000000)
    