CREATE TABLE test (v Variant(String, UInt64)) ENGINE = Memory;
INSERT INTO test VALUES ('hello'), (42), ('foo');

-- Default (throw on mismatch): length() does not accept UInt64, so the query throws.
SELECT length(v) FROM test;  -- throws ILLEGAL_TYPE_OF_ARGUMENT

-- With throw disabled: incompatible rows return NULL.
SET variant_throw_on_type_mismatch = false;
SELECT v, length(v) FROM test ORDER BY v::String NULLS LAST;
