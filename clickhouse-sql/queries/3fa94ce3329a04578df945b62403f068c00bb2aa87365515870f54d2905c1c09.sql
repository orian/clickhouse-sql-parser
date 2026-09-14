
        CREATE DICTIONARY IF NOT EXISTS complex_key_{layout_suffix}_dictionary_l0_{load_factor}
        (
            id UInt64,
            id_key String,
            value_int UInt64
        )
        PRIMARY KEY id, id_key
        SOURCE(CLICKHOUSE(TABLE 'complex_key_dictionary_source_table'))
        LAYOUT(COMPLEX_KEY_{layout_suffix}(MAX_LOAD_FACTOR 0.{load_factor}))
        LIFETIME(0)
    