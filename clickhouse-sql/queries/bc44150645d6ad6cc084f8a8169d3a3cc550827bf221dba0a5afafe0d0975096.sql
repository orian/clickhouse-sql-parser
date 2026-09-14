
        CREATE TABLE complex_key_hashed_dictionary_source_table
        (
            id UInt64,
            id_key String,
            value_int UInt64,
            value_string String,
            value_decimal Decimal64(8),
            value_string_nullable Nullable(String)
        ) ENGINE = Memory;
    