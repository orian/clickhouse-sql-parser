SELECT
    url,
    count() AS cnt
FROM hits_clickhouse
WHERE counterid = 38
GROUP BY url
ORDER BY cnt DESC
LIMIT 5
