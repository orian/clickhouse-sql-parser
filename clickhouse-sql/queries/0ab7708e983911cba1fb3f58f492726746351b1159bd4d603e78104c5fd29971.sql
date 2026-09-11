
        INSERT INTO phj_asof_right
        SELECT number % 500, if(number % 30 = 0, NULL, toDateTime('2020-01-01') + number * 2), toString(number)
        FROM numbers(2000000)
    