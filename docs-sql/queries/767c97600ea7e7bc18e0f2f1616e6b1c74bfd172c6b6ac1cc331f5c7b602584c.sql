CREATE TABLE t (num UInt32) ENGINE = Memory;
INSERT INTO t VALUES (44), (28), (13), (85);

-- Test data:
-- binary     decimal
-- 00101100 = 44
-- 00011100 = 28
-- 00001101 = 13
-- 01010101 = 85

SELECT groupBitOr(num) FROM t;
