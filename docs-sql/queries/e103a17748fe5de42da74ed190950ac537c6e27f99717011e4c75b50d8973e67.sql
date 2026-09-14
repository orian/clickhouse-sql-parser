SELECT
    date,
    count()
FROM sensors
GROUP BY date
ORDER BY date ASC;
