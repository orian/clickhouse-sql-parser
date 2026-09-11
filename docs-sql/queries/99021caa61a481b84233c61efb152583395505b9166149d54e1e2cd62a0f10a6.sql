SELECT
    value,
    time,
    round(exp_smooth, 3),
    bar(exp_smooth, 0, 20, 50) AS bar
FROM
(
    SELECT
        (number % 5) = 0 AS value,
        number AS time,
        exponentialTimeDecayedCount(10)(time) OVER (ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS exp_smooth
    FROM numbers(50)
)
