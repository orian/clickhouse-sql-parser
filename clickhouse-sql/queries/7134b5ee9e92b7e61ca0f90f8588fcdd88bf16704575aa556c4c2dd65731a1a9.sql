
    SELECT count()
    FROM numbers_mt(1e8) AS lhs
    INNER JOIN numbers_mt(1e5) AS rhs ON (lhs.number % 100000) = rhs.number
    FORMAT Null
    SETTINGS max_threads = 8, join_algorithm = 'hash'
  