CREATE TABLE t (
    n Int32,
    val UInt32
) ENGINE = Memory;

INSERT INTO t VALUES (0, 3), (1, 2), (2, 1), (5, 4);

SELECT quantileInterpolatedWeighted(n, val) FROM t;
