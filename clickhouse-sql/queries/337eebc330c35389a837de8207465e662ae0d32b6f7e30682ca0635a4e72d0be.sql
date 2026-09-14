
        CREATE DICTIONARY dict_ip_trie
        (
            ip String,
            ver UInt8,
            val Float32
        )
        PRIMARY KEY ip
        SOURCE(CLICKHOUSE(DB 'default' TABLE 'table_ip_trie'))
        LAYOUT(IP_TRIE())
        LIFETIME(300)
    