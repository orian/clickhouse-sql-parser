
        SELECT a, b, c, z FROM limit_by_in_pk_order_int
        WHERE z > 5 LIMIT 5 BY a
        FORMAT NULL
    