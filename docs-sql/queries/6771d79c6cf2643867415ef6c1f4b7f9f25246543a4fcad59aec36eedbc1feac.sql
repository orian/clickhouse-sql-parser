INSERT INTO test VALUES (42), (43), ('str_1');
SELECT d, dynamicType(d) FROM test;
