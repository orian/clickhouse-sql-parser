
        CREATE DICTIONARY simple_key_flat_dictionary
        (
            id UInt64,
            value_int UInt64,
            value_string String,
            value_decimal Decimal64(8),
            value_string_nullable Nullable(String)
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'simple_key_flat_dictionary_source_table'))
        LAYOUT(FLAT(INITIAL_ARRAY_SIZE 50000 MAX_ARRAY_SIZE 5000000))
        LIFETIME(MIN 0 MAX 1000)
    