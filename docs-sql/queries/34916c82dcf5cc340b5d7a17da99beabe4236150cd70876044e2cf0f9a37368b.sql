-- Note the ARRAY JOIN syntax in the `SELECT` query below, which provides broader possibilities.
-- ARRAY JOIN allows you to convert multiple arrays with the same number of elements at a time.

SELECT
    sum(1) AS impressions,
    city,
    browser
FROM
(
    SELECT
        ['Istanbul', 'Berlin', 'Bobruisk'] AS cities,
        ['Firefox', 'Chrome', 'Chrome'] AS browsers
)
ARRAY JOIN
    cities AS city,
    browsers AS browser
GROUP BY
    2,
    3
ORDER BY
    2,
    3
