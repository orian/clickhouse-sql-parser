SELECT one
FROM
  (
    SELECT main, sum(two) AS one
    FROM tabl
    WHERE
      x > 1
  )
WHERE
  main LIKE '%olsztyn%'
ORDER BY
  one DESC
LIMIT 10
