CREATE TABLE bits
(
    `a` UInt8,
    `b` UInt8
)
ENGINE = Memory;

INSERT INTO bits VALUES (0, 0), (0, 1), (1, 0), (1, 1);

SELECT
    a,
    b,
    bitAnd(a, b)
FROM bits
