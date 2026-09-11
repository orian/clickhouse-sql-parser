
        SELECT a, b, c, z FROM limit_by_in_pk_order_low
        WHERE a = 1 LIMIT 5 BY b
        FORMAT NULL
    