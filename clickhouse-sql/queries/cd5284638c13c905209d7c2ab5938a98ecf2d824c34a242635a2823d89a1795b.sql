
        CREATE DICTIONARY IF NOT EXISTS simple_key_{layout_suffix}_dictionary_l0_{load_factor}
        (
            id UInt64,
            value_int UInt64
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(TABLE 'simple_key_dictionary_source_table'))
        LAYOUT({layout_suffix}(MAX_LOAD_FACTOR 0.{load_factor}))
        LIFETIME(0)
    