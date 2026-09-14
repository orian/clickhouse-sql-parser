FROM orders
|> AGGREGATE sum(amount) AS total GROUP BY customer
|> WHERE total > 100
