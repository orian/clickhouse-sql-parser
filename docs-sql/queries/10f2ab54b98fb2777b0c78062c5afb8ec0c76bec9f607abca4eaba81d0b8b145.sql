FROM customers
|> AS c
|> LEFT JOIN orders AS o ON c.name = o.customer
|> ARRAY JOIN tags
