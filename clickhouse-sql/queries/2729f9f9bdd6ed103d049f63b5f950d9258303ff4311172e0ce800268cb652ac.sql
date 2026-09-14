
        CREATE TABLE limit_by_in_pk_order_merge_1
        (a Int32, b UInt32, z UInt32)
        ENGINE = MergeTree
        ORDER BY (a, b)
    