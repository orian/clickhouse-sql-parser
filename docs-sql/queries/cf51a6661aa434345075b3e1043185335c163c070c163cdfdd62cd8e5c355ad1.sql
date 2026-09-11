CREATE TABLE test.Orders
(
    OrderId UInt64,
    OrderName String,
    OrderDate DateTime
) ENGINE = MergeTree
ORDER BY ();
