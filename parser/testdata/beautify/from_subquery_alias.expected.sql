SELECT *
FROM
  (
    SELECT a
    FROM t
  ) AS sub
WHERE sub.a < 10
