
        INSERT INTO limit_by_in_pk_order_part_in_order
        SELECT number, number % 97, number
        FROM numbers(50000000)
    