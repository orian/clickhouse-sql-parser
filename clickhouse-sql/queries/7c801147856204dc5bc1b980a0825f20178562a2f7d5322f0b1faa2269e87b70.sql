
        CREATE TABLE dict_ip_trie_table
        (
            `ip` String,
            `ver` UInt8,
            `val` Float32
        ) ENGINE = Dictionary(default.dict_ip_trie)
    