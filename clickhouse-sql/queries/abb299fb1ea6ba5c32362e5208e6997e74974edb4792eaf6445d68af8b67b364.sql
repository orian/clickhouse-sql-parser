
        SELECT avg(o_totalprice)
        FROM orders
        JOIN (SELECT * FROM customer JOIN nation ON c_nationkey = n_nationkey WHERE n_name = 'ETHIOPIA') AS cn 
        ON c_custkey = o_custkey
        SETTINGS enable_join_runtime_filters=1
    