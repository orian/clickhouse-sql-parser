
        INSERT INTO limit_by_in_pk_order_merge_1
        SELECT number % 100000, number % 97, number
        FROM numbers(5000000)
    