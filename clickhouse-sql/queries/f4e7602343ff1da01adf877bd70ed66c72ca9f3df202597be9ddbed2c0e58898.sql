
        CREATE TABLE test_minmax_index
        (
        x UInt64,
        INDEX idx_x (x) TYPE minmax
        )
        ENGINE = MergeTree
        ORDER BY ()
        SETTINGS index_granularity = 5;
    