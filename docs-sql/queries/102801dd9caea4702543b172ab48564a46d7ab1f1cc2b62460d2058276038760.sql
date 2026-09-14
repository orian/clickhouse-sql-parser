SELECT addr, cutIPv6(IPv6StringToNum(addr), 0, 0) FROM (SELECT ['127.0.0.1', '1111::ffff'] AS addr) ARRAY JOIN addr;
