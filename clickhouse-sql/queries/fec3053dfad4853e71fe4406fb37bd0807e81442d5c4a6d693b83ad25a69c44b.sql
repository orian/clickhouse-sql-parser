
    SELECT count()
    FROM (SELECT number % 1000 AS a FROM numbers(1e5)) AS l
    RIGHT JOIN (SELECT number % 100000 AS k1, number % 100000 AS k2 FROM numbers_mt(5e6)) AS r
    ON l.a = r.k1 OR l.a = r.k2
    FORMAT Null
  