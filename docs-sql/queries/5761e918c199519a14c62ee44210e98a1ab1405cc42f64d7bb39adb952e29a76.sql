SELECT region, SUM(amount) AS sum, AVG(amount) AS mean
FROM file('sales.csv', 'CSVWithNames')
WHERE amount > 1000
GROUP BY region
ORDER BY sum DESC
LIMIT 10
