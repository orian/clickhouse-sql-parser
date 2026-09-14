
    SELECT count()
    FROM numbers_mt(5e6) AS lhs
    RIGHT JOIN numbers_mt(1e5) AS rhs ON (lhs.number = (rhs.number + 1000000000)) OR ((lhs.number % 100000) = rhs.number)
    FORMAT `Null`
  