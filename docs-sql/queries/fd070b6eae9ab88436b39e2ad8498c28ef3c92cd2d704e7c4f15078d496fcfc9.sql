SELECT
    tempMax / 10 AS maxTemp,
    location,
    name,
    date
FROM blogs.noaa
WHERE tempMax > 500
ORDER BY
    tempMax DESC,
    date ASC
LIMIT 5
