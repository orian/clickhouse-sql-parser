SELECT
    anyIf(transaction_id, amount < 200) AS tid_lt_200,
    anyIf(transaction_id, amount > 200) AS tid_gt_200
FROM sales;
