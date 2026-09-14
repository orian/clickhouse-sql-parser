FROM table
|> WHERE x > 1
|> AGGREGATE count() AS c GROUP BY y
|> ORDER BY c DESC
