SELECT
    uuid,
    sum(count) AS count
FROM events
GROUP BY ALL
ORDER BY count DESC
LIMIT 10
