CREATE TABLE t
(
    a UInt64,
    s String
)
ENGINE = MergeTree
ORDER BY a
SETTINGS storage_policy = 's3_express';
