CREATE TABLE sales(
    transaction_id UInt32,
    amount Decimal(10,2),
    is_successful UInt8
) ENGINE = MergeTree
ORDER BY ();

INSERT INTO sales VALUES
    (1, 100.50, 1),
    (2, 200.75, 1),
    (3, 150.25, 0),
    (4, 300.00, 1),
    (5, 250.50, 0),
    (6, 175.25, 1);

SELECT
    avgIf(amount, is_successful = 1) AS avg_successful_sale
FROM sales;
