FROM orders |> ORDER BY amount DESC;
FROM orders |> SELECT customer, amount |> ORDER BY ALL;
FROM points |> ORDER BY x WITH FILL FROM 1 TO 10 INTERPOLATE (y AS y + 1)
