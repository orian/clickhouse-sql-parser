
        SELECT a, b, z FROM limit_by_in_pk_order_dt
        LIMIT 5 BY toStartOfHour(a)
        FORMAT NULL
    