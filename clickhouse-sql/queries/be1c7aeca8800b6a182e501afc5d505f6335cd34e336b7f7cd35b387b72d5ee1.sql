
        SELECT dictGetFloat32('default.dict_ip_trie', 'val', tuple(randomFixedString(16)))
        FROM numbers(500000) FORMAT Null
    