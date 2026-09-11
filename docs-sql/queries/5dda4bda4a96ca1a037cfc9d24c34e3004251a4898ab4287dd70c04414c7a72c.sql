CREATE TABLE test (id UInt32, vec QBit(Float32, 8)) ENGINE = Memory;
INSERT INTO test VALUES (1, [1, 2, 3, 4, 5, 6, 7, 8]);
SELECT vec[3], arraySlice(vec, 2, 3) FROM test;
