SELECT addr, isIPv4String(addr)
FROM(
SELECT ['0.0.0.0', '127.0.0.1', '::ffff:127.0.0.1'] AS addr
)
ARRAY JOIN addr;
