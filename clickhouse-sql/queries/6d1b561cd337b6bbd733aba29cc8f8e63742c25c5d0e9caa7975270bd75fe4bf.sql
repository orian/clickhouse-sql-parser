
        CREATE TABLE dict_src_const_path
        (
            id UInt64,
            value     UInt32,
            value_str String
        )
        ENGINE = MergeTree
        ORDER BY id;
    