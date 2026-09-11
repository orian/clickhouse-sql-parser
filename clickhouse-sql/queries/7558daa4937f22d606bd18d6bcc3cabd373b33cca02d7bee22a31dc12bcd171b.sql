
        CREATE TABLE test_in_skip_idx
        (
            a UInt64,
            s String,
            INDEX idx s TYPE bloom_filter GRANULARITY 1
        )
        ENGINE = MergeTree() ORDER BY a
    