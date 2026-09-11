CREATE TABLE test (t UInt8) ENGINE = Memory;

SELECT avg(t) FROM test;
