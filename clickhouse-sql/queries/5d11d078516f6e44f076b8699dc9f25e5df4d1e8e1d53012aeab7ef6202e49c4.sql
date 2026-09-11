
        INSERT INTO limit_by_in_pk_order_int
        SELECT number % 100000, number % 97, number % 7, number
        FROM numbers(5000000)
    