CREATE TABLE key_val (key1 String, key2 String, value UInt32) ENGINE = Memory;
INSERT INTO key_val VALUES ('Hello, ', 'World!', 1), ('Hello, ', 'World!', 2), ('Hello, ', 'World', 3);

SELECT concatAssumeInjective(key1, key2) AS key, sum(value) FROM key_val GROUP BY key ORDER BY key;
