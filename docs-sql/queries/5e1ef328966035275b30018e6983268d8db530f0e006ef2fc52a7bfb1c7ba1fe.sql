SELECT path, sum(hits) AS hits
FROM wikistat_top
WHERE month = '2015-05-01'
GROUP BY ALL
ORDER BY hits DESC
LIMIT 10;
