SELECT region, sum(amount) AS total
FROM `sales.orders`
GROUP BY region
ORDER BY total DESC;
