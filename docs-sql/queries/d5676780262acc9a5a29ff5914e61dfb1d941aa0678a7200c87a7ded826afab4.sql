CREATE TABLE test (d Dynamic(max_types=2)) ENGINE = Memory;
INSERT INTO test VALUES (NULL), (42), ('Hello, World!'), ([1, 2, 3]);
SELECT d, isDynamicElementInSharedData(d) FROM test;
