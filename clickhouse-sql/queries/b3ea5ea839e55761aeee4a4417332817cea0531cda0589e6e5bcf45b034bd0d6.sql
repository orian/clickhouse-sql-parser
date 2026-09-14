
        INSERT INTO table_ip_from_dict
        SELECT splitByChar('/', ip )[1] as ip, ver FROM dict_ip_trie_table
    