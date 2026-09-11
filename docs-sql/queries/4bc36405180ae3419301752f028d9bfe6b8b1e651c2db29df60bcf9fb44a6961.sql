CREATE TABLE test (id UInt32, vec QBit(Float32, 8)) ENGINE = Memory;
INSERT INTO test VALUES (1, [0, 0, 0, 0, 0, 0, 0, 0]);
INSERT INTO test VALUES (1, [-0, -0, -0, -0, -0, -0, -0, -0]);
SELECT bin(vec.1) FROM test;
