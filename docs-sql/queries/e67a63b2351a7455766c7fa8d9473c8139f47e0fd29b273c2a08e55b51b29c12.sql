CREATE TABLE test (v Variant(UInt64, String)) ENGINE = Memory;
INSERT INTO test VALUES (NULL), (42), ('42.42');
SELECT v::Nullable(Float64) FROM test;
