SELECT 
    region,
    category,
    SUM(amount) AS amount_sum,
    AVG(amount) AS amount_mean,
    COUNT(amount) AS amount_count
FROM file('sales.csv', 'CSVWithNames')
WHERE date >= '2024-01-01'
  AND amount > 100
GROUP BY region, category
HAVING amount_sum > 10000
ORDER BY amount_sum DESC
LIMIT 20
