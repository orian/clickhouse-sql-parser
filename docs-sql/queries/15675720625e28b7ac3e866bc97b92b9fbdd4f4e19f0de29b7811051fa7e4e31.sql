SET variant_throw_on_type_mismatch = 0;
CREATE TABLE test3 (v Variant(Array(UInt32), UInt32)) ENGINE = Memory;
INSERT INTO test3 VALUES ([1,2,3]), (42);
SELECT v + 10 AS result, toTypeName(result) FROM test3;
