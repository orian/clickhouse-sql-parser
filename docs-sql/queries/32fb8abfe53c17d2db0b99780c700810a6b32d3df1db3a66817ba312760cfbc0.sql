CREATE TABLE tab (m Map(String, UInt64)) ENGINE = Memory;
INSERT INTO tab VALUES (map('key1', 1, 'key2', 2, 'key3', 3));

SELECT m.keys FROM tab; --   与 mapKeys(m) 相同
SELECT m.values FROM tab; -- 与 mapValues(m) 相同
