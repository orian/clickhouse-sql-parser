
        INSERT INTO limit_by_in_pk_order_low
        SELECT number % 5, intDiv(number, 5), number % 7, number
        FROM numbers(5000000)
    