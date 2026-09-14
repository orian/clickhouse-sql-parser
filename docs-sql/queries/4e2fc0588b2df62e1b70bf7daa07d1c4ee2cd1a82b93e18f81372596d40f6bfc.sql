SET enable_nullable_tuple_type = 1;

CREATE TABLE test (
    id UInt32,
    data Nullable(Tuple(String, Int64))
) ENGINE = Memory;

INSERT INTO test VALUES (1, ('hello', 42)), (2, NULL);

SELECT * FROM test WHERE data IS NULL;
