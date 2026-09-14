
        CREATE DICTIONARY hierarchical_flat_dictionary
        (
            id UInt64,
            parent_id UInt64 HIERARCHICAL
        )
        PRIMARY KEY id
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'hierarchical_dictionary_source_table'))
        LAYOUT(FLAT(max_array_size 1000001))
        LIFETIME(0);
    