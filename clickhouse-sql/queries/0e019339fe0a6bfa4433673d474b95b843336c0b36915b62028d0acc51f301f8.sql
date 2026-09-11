
        CREATE DICTIONARY join_hashed_dictionary (key UInt64, value String)
        PRIMARY KEY key 
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'join_dictionary_source_table'))
        LIFETIME(MIN 0 MAX 1000)
        LAYOUT(HASHED());
    