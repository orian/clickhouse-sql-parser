
        CREATE TABLE test_lwu_steps
        (
            id UInt64,
            a String,
            b UInt64,
            c UInt64,
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS enable_block_number_column = 1, enable_block_offset_column = 1;
    