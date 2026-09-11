
        INSERT INTO phj_left
        SELECT number % 1000, toString(number)
        FROM numbers(5000000)
    