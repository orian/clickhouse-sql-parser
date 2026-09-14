CREATE TABLE t (
    n Int32,
    val UInt32
) ENGINE = Memory;

-- Insert the sample data
INSERT INTO t VALUES
(0, 3),
(1, 2),
(2, 1),
(5, 4);

SELECT quantileExactWeighted(n, val) FROM t;
