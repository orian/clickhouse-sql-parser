
        INSERT INTO phj_asof_left
        SELECT number % 500, toDateTime('2020-01-01') + number, toString(number)
        FROM numbers(2000000)
    