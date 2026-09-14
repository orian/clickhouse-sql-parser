SELECT Year, c1/c2
FROM
(
    SELECT
        Year,
        count(*)*100 AS c1
    FROM ontime
    WHERE DepDelay>10
    GROUP BY Year
) q
JOIN
(
    SELECT
        Year,
        count(*) AS c2
    FROM ontime
    GROUP BY Year
) qq USING (Year)
ORDER BY Year;
