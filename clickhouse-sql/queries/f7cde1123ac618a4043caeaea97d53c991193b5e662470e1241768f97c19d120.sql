
        SELECT avg(o_totalprice)
        FROM orders
        JOIN customer
        ON c_custkey = o_custkey
        SETTINGS enable_join_runtime_filters=0
    