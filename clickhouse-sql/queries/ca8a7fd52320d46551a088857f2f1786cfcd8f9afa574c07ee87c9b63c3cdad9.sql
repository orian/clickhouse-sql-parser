
        CREATE TABLE test_set_index_check_in_range (a UInt32, b UInt32)
        ENGINE = MergeTree ORDER BY (a, b) SETTINGS index_granularity = 128
    