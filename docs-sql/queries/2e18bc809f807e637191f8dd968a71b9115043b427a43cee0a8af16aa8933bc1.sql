SELECT
    toDate(time) AS date,
    sum(hits) AS hits
FROM wikistat
GROUP BY ALL
ORDER BY date ASC
LIMIT 5;
