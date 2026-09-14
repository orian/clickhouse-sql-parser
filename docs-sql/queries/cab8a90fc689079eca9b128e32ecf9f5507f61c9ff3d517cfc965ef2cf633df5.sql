FROM orders |> AGGREGATE count() AS c, sum(amount) AS total GROUP BY customer;
FROM orders |> AGGREGATE count() AS c;
