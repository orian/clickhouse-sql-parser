
        INSERT INTO limit_by_in_pk_order_dt
        SELECT toDateTime(number % 100000), number % 97, number
        FROM numbers(5000000)
    