SELECT
    toDate(time) AS day,
    sum(hits) AS h,
    lagInFrame(h) OVER (ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) AS p,
    h - p AS trend
FROM wikistat
WHERE path = '"Weird_Al"_Yankovic'
GROUP BY ALL
LIMIT 10;
