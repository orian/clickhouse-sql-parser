TRUNCATE TABLE test;
INSERT INTO test VALUES (NULL), ([1, 2]), ([3, 4]);
SELECT d, dynamicType(d) FROM test;
