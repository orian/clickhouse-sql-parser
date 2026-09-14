
        CREATE DICTIONARY complex_key_direct_dictionary
        (
            id UInt64,
            id_key String,
            value_int UInt64,
            value_string String,
            value_decimal Decimal64(8),
            value_string_nullable Nullable(String)
        )
        PRIMARY KEY id, id_key
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'complex_key_direct_dictionary_source_table'))
        LAYOUT(COMPLEX_KEY_DIRECT())
    