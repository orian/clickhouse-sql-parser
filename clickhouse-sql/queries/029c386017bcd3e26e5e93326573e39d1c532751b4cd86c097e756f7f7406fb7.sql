
        CREATE DICTIONARY hierarchical_{dictionary_layout}_shards{dictionary_shards}_dictionary
        (
            id UInt64,
            parent_id UInt64 HIERARCHICAL
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'hierarchical_dictionary_source_table'))
        LAYOUT({dictionary_layout}(SHARDS {dictionary_shards}))
        LIFETIME(0);
    