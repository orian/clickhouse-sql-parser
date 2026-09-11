CREATE TABLE test_table
(
    key UInt64,
    value_int Nullable(UInt32),
    value_string Nullable(String),
    value_date Nullable(Date)
)
ENGINE = CoalescingMergeTree()
ORDER BY key
