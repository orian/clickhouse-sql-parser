CREATE TABLE tab (m Map(String, UInt64)) ENGINE = Memory;
INSERT INTO tab VALUES (map('key1', 1, 'key2', 2, 'key3', 3));

SELECT m.keys FROM tab; --   mapKeys(m) と同じ
SELECT m.values FROM tab; -- mapValues(m) と同じ
