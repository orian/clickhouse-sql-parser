CREATE TABLE test_data (x Float64) ENGINE = Memory;
INSERT INTO test_data VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

SELECT kurtPop(x) FROM test_data;
