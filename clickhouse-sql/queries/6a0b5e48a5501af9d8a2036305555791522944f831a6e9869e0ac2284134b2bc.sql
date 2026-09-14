
        SELECT a, b, z FROM limit_by_in_pk_order_part_in_order
        LIMIT 5 BY a
        FORMAT NULL
        SETTINGS allow_limit_by_partitions_independently = 1
    