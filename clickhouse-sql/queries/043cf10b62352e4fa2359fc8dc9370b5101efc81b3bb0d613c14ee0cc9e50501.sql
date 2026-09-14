
    SELECT count()
    FROM (SELECT toString(number) AS key FROM numbers_mt(10000000)) AS l
    FULL JOIN (SELECT toString(number + 100000000) AS key FROM numbers_mt(10000000)) AS r
    ON l.key = r.key
  