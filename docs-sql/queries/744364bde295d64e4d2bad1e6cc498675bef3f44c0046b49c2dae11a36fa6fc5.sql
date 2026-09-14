FROM orders
|> AGGREGATE sum(amount) AS total GROUP BY customer
|> AS agg
|> JOIN orders AS o ON agg.customer = o.customer
