
        CREATE DICTIONARY complex_key_range_hashed_dictionary
        (
            id UInt64,
            id_key String,
            value UInt64,
            start UInt64,
            end UInt64
        )
        PRIMARY KEY id, id_key
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'complex_key_range_hashed_dictionary_source_table'))
        LAYOUT(COMPLEX_KEY_RANGE_HASHED())
        RANGE(MIN start MAX end)
        LIFETIME(MIN 0 MAX 1000);
    