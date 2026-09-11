
        CREATE TABLE test_lwu_join_2
        (
            id UInt64,
            uint_1p UInt64,
            uint_10p UInt64,
            str_1p String,
            str_10p String,
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS apply_patches_on_merge = 0, enable_block_number_column = 1, enable_block_offset_column = 1;
    