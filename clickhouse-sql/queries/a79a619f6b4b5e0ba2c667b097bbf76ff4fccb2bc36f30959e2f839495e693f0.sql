
        SELECT dictGetFloat32('default.dict_ip_trie', 'val', tuple(IPv6StringToNum(ip)))
        FROM table_ip_from_dict
        WHERE ver == 4
        LIMIT 500000 FORMAT Null
    