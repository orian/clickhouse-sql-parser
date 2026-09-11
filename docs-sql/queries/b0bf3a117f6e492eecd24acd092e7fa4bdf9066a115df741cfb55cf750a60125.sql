CREATE TABLE test_table (CounterID UInt32, UserID UInt32, ver UInt32) ENGINE = Memory;
INSERT INTO test_table VALUES (1, 3, 3), (1, 1, 1), (1, 2, 2), (1, 5, 5), (1, 4, 4);

SELECT *, generateSerialID('id1') FROM test_table
