-- использование MATERIALIZED столбца
CREATE TABLE t
(
    id UInt64,
    a UInt32,
    ab_sum UInt64 MATERIALIZED a + 1,
    PROJECTION p (SELECT a ORDER BY ab_sum)
)
ENGINE = MergeTree ORDER BY id;

-- использование встроенного выражения
CREATE TABLE t
(
    id UInt64,
    a UInt32,
    PROJECTION p (SELECT a ORDER BY a + 1)
)
ENGINE = MergeTree ORDER BY id;
