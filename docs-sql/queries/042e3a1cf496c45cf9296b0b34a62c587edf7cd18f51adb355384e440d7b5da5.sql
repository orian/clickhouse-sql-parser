SELECT addr, isIPv6String(addr)
FROM(SELECT ['::', '1111::ffff', '::ffff:127.0.0.1', '127.0.0.1'] AS addr)
ARRAY JOIN addr;
