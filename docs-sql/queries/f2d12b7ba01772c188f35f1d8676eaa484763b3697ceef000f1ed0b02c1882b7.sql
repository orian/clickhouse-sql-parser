CREATE TABLE t (val UInt32) ENGINE = Memory;
INSERT INTO t VALUES (1), (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

SELECT quantiles(0.25, 0.5, 0.75, 0.9)(val) FROM t;
