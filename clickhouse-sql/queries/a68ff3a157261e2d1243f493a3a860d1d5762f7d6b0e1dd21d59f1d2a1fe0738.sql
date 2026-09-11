
        SELECT dictGetFloat32('default.dict_ip_trie', 'val', tuple(rand32()))
        FROM numbers(500000) FORMAT Null
    