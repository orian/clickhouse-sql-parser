FROM orders
|> SELECT customer
|> UNION ALL (FROM customers |> SELECT name)
|> DISTINCT
