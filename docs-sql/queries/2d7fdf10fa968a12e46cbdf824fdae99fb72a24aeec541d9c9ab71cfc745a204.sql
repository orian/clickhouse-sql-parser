CREATE TABLE test (t UInt8) ENGINE = Memory;
SELECT avgWeighted(t, t) FROM test
