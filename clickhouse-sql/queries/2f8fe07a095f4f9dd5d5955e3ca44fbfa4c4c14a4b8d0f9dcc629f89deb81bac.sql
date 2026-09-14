
    SELECT count()
    FROM numbers_mt(3e7) AS lhs
    INNER JOIN numbers_mt(1e4) AS rhs ON (lhs.number = (rhs.number + 1000000000)) OR (lhs.number = (rhs.number + 2000000000))
    FORMAT Null
    SETTINGS max_threads = 1, join_algorithm = 'hash'
  