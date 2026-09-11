
        SELECT a, b, c, z FROM limit_by_in_pk_order_int
        LIMIT 5 BY negate(a)
        FORMAT NULL
    