
        CREATE TABLE test_set_index
        (
            x UInt64,
            INDEX idx_x (x) TYPE set(10)
        )
        ENGINE = MergeTree
        ORDER BY ()
        SETTINGS index_granularity = 5;
    