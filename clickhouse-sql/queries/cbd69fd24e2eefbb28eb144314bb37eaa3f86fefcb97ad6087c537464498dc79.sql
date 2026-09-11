
        CREATE TABLE limit_by_in_pk_order_multi_part
        (a UInt32, b UInt32, z UInt32)
        ENGINE = MergeTree
        ORDER BY a
        PARTITION BY b % 16
    