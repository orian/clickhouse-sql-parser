SELECT kolmogorovSmirnovTest('less', 'exact')(value, num)
FROM
(
    SELECT
        randNormal(0, 10) AS value,
        0 AS num
    FROM numbers(10000)
    UNION ALL
    SELECT
        randNormal(0, 10) AS value,
        1 AS num
    FROM numbers(10000)
)
