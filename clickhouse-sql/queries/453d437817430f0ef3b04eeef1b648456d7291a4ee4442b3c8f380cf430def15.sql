
        CREATE DICTIONARY simple_key_direct_dictionary
        (
            id UInt64,
            value_int UInt64,
            value_string String,
            value_decimal Decimal64(8),
            value_string_nullable Nullable(String)
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'simple_key_direct_dictionary_source_table'))
        LAYOUT(DIRECT())
    