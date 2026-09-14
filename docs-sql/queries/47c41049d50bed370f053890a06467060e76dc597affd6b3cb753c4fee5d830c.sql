TRUNCATE TABLE test;
INSERT INTO test VALUES (NULL), ('str_1'), ('str_2');
SELECT d, dynamicType(d) FROM test;
