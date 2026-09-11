
        CREATE DICTIONARY IF NOT EXISTS simple_key_{layout_suffix}_dictionary_s{shards}
        (
            id UInt64,
            value_int UInt64
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(TABLE 'simple_key_dictionary_source_table'))
        LAYOUT({layout_suffix}(SHARDS {shards}))
        LIFETIME(0)
    