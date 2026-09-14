
        CREATE DICTIONARY simple_key_range_hashed_dictionary
        (
            id UInt64,
            value UInt64,
            start UInt64,
            end UInt64
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'simple_key_range_hashed_dictionary_source_table'))
        LAYOUT(RANGE_HASHED())
        RANGE(MIN start MAX end)
        LIFETIME(MIN 0 MAX 1000);
    