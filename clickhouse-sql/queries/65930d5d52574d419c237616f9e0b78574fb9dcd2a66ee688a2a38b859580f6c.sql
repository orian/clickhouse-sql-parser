
        INSERT INTO table_ip_trie
        SELECT
            IPv6NumToString(ipv6) || '/' || toString(rand() % 65 + 64) as ip,
            6 as ver,
            val
        FROM generateRandom('ipv6 FixedString(16), val Float32', 0, 30, 30)
        LIMIT 2500000
    