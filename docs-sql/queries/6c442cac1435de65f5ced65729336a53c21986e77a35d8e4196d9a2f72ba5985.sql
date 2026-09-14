CREATE TABLE t (val UInt32) ENGINE = Memory;
INSERT INTO t VALUES (1), (1), (2), (3);

SELECT quantile(val) FROM t;
