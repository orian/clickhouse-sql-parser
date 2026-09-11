SELECT
    value,
    time,
    round(exp_smooth, 3),
    bar(exp_smooth, 0, 1, 50) AS bar
FROM
(
    SELECT
        (number = 0) OR (number >= 25) AS value,
        number AS time,
        exponentialMovingAverage(10)(value, time) OVER (Rows BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS exp_smooth
    FROM numbers(50)
)
