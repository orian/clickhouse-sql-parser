SET variant_throw_on_type_mismatch = 0;
CREATE TABLE test2 (v Variant(UInt64, Float64)) ENGINE = Memory;
INSERT INTO test2 VALUES (42::UInt64), (42.42);
SELECT v + 1 AS result, toTypeName(result) FROM test2;
