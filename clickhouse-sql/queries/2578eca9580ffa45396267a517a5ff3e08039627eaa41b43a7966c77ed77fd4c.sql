
    SELECT max(r.k1)
    FROM (SELECT number % 10000 AS a, number % 10000 AS b FROM numbers_mt(1e7)) AS l
    RIGHT JOIN (SELECT number % 1000 AS k1, number % 1000 AS k2 FROM numbers(1e5)) AS r
    ON l.a = r.k1 OR l.b = r.k2
    FORMAT Null
  