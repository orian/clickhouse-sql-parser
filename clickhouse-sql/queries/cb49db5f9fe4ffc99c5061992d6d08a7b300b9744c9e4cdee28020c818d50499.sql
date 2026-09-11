
        CREATE TABLE test_set (k UInt32, x UInt32, INDEX idx (x) TYPE set(10) GRANULARITY 1) ENGINE = MergeTree ORDER BY k SETTINGS index_granularity = 111;
    