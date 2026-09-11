
        CREATE TABLE test_group_by_strings (
            key1_8 String,
            key2_8 String,
            key1_16 String,
            key2_16 String,
            key1_32 String,
            key2_32 String,
            key1_64 String,
            key2_64 String,
            key1_512 String,
            key2_512 String,
            value UInt64
        ) ENGINE = MergeTree ORDER BY (key1_8, key2_8)
    