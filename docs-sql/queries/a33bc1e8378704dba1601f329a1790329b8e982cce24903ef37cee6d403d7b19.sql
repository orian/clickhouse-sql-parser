SELECT
    toStartOfHour(time) AS time,
    sum(hits) AS hits,
    round(hits / (60 * 60), 2) AS rate,
    bar(rate * 10, 0, max(rate * 10) OVER (), 25) AS b
FROM wikistat
WHERE path = '"Weird_Al"_Yankovic'
GROUP BY time
LIMIT 10;
