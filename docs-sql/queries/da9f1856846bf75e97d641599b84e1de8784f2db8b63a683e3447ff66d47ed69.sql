FROM orders
|> WHERE cancelled = 0
|> AGGREGATE sum(amount) AS total GROUP BY customer
|> ORDER BY total DESC
|> LIMIT 3
