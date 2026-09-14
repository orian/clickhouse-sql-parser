SELECT
    toDate(time) AS day,
    sum(hits) AS h,
    sum(h) OVER (ROWS BETWEEN UNBOUNDED PRECEDING AND 0 FOLLOWING) AS c,
    bar(c, 0, 50000, 25) AS b
FROM wikistat
WHERE path = '"Weird_Al"_Yankovic'
GROUP BY ALL
ORDER BY day
LIMIT 10;
