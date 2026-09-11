CREATE TABLE t
(
    id UInt64,
    a UInt32,
    ab_sum UInt64 ALIAS a + 1,
    PROJECTION p (SELECT a ORDER BY ab_sum)
)
ENGINE = MergeTree ORDER BY id;
-- 失败，报错 UNKNOWN_IDENTIFIER
