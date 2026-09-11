DROP TABLE IF EXISTS test_dynamic;
CREATE TABLE test_dynamic(d Dynamic) ENGINE = Memory;
INSERT INTO test_dynamic VALUES (42), (NULL), ('Hello'), ([1, 2, 3]), ('2020-01-01'), (map(1, 2)), (43), ([4, 5]), (NULL), ('World'), (map(3, 4));

SELECT distinctDynamicTypes(d) FROM test_dynamic;
