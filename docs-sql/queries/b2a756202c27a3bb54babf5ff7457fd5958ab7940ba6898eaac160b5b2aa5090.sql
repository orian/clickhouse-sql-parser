CREATE TABLE t (x UInt8) ENGINE = Memory;
INSERT INTO t VALUES (1), (2), (1), (3), (2), (4);

SELECT groupUniqArray(x) FROM t;
