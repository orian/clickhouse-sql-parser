
        INSERT INTO table_ip_trie
        SELECT
            IPv4NumToString(ipv4) || '/' || toString(rand() % 17 + 16) as ip,
            4 as ver,
            val
        FROM generateRandom('ipv4 UInt32, val Float32', 0, 30, 30)
        LIMIT 200000
    