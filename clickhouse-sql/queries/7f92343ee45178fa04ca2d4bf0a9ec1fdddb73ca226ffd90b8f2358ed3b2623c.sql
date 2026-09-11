
        CREATE TABLE limit_by_in_pk_order_part_in_order
        (a Int32, b UInt32, z UInt32)
        ENGINE = MergeTree
        ORDER BY (a, b)
        PARTITION BY a % 100
    