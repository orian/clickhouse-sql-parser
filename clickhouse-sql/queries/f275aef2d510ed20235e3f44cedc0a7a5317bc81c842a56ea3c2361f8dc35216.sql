
        CREATE DICTIONARY dict_const_path
        (
            id UInt64,
            value     UInt32,
            value_str String
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(TABLE 'dict_src_const_path'))
        LIFETIME(0)
        LAYOUT(HASHED());
    