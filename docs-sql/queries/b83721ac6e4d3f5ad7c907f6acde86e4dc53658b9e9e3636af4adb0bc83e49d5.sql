SELECT
    toStartOfHour(time) AS hour,
    sum(hits) AS hits    
FROM wikistat
WHERE date(time) = '2015-07-01'
GROUP BY ALL
ORDER BY hour ASC
LIMIT 5;
