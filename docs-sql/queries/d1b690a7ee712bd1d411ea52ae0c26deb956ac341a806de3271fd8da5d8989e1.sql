CREATE TABLE t
(
    `int` UInt8,
    `float` Float32,
    `dec` Decimal32(2)
)
ENGINE = Memory;

INSERT INTO t VALUES (1, 1.1, 1.10), (2, 2.2, 2.20), (4, 4.4, 4.40), (7, 7.77, 7.77);

SELECT
    groupArrayMovingAvg(int) AS I,
    groupArrayMovingAvg(float) AS F,
    groupArrayMovingAvg(dec) AS D
FROM t;
