
        CREATE TABLE limit_by_in_pk_order_int
        (a Int32, b UInt32, c UInt32, z UInt32)
        ENGINE = MergeTree
        ORDER BY (a, b, c)
    