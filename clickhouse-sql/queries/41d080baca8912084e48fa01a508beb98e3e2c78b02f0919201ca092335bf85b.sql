
        INSERT INTO limit_by_in_pk_order_multi_part
        SELECT number % 100000, number, number
        FROM numbers(16000000)
    