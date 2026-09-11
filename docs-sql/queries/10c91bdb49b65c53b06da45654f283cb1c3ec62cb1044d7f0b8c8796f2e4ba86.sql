CREATE TABLE tab (m Map(String, UInt64)) ENGINE = Memory;
INSERT INTO tab VALUES (map('key1', 1, 'key2', 2, 'key3', 3));

SELECT m.keys FROM tab; --   то же самое, что mapKeys(m)
SELECT m.values FROM tab; -- то же самое, что mapValues(m)
