CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- Default (throw on mismatch): length() does not accept integers, so the query throws.
SELECT length(d) FROM test;  -- throws ILLEGAL_TYPE_OF_ARGUMENT

-- With throw disabled: incompatible rows return NULL.
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
