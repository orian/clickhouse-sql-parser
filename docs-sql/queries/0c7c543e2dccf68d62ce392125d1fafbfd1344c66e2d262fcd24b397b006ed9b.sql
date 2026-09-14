CREATE TABLE test (v Variant(UInt64, String, Array(UInt64))) ENGINE = Memory;
INSERT INTO test VALUES (NULL), (42), ('String'), ([1, 2, 3]);
SELECT v::Dynamic AS d, dynamicType(d) FROM test;
