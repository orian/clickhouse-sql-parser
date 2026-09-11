
        CREATE TABLE test_sparse_{ratio} (id UInt64, u8 UInt8, u64 UInt64, str String)
        ENGINE = MergeTree ORDER BY id
        SETTINGS ratio_of_defaults_for_sparse_serialization = 0.8
    