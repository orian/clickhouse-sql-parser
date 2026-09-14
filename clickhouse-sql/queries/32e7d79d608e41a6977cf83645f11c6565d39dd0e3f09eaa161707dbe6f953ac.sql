
        CREATE TABLE read_in_reverse_order_final (x UInt64, y UInt64)
        ENGINE = ReplacingMergeTree()
        ORDER BY x
    