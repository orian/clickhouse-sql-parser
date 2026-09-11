-- A query can use multiple arrayJoin functions. In this case, the transformation is performed multiple times and the rows are multiplied.

SELECT
    sum(1) AS impressions,
    arrayJoin(cities) AS city,
    arrayJoin(browsers) AS browser
FROM
(
    SELECT
        ['Istanbul', 'Berlin', 'Bobruisk'] AS cities,
        ['Firefox', 'Chrome', 'Chrome'] AS browsers
)
GROUP BY
    2,
    3
ORDER BY
    city,
    browser
