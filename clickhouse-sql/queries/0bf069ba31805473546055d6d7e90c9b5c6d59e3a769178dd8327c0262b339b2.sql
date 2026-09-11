
    SELECT count()
    FROM numbers_mt(1e7) AS lhs
    RIGHT JOIN numbers_mt(1e5) AS rhs ON (lhs.number % 100000) = rhs.number
    FORMAT Null
    SETTINGS max_threads = 1, join_algorithm = 'hash'
  