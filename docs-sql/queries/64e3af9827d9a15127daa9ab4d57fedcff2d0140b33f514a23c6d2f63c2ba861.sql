SELECT
    a * 2,
    b,
    count(c),
FROM t
GROUP BY ALL
