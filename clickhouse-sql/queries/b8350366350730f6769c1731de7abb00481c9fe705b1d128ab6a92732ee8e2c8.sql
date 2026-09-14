
        CREATE DICTIONARY IF NOT EXISTS complex_key_{layout_suffix}_dictionary_s{shards}
        (
            id UInt64,
            id_key String,
            value_int UInt64
        )
        PRIMARY KEY id, id_key
        SOURCE(CLICKHOUSE(TABLE 'complex_key_dictionary_source_table'))
        LAYOUT(COMPLEX_KEY_{layout_suffix}(SHARDS {shards}))
        LIFETIME(0)
    