CREATE TABLE test (d Dynamic(max_types=3)) ENGINE = Memory;
INSERT INTO test VALUES (NULL), (42), (43), ('42.42'), (true);
SELECT d::Dynamic(max_types=5) as d2, dynamicType(d2) FROM test;
