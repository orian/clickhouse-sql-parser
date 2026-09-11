SELECT
    value,
    time,
    round(exp_smooth, 3),
    bar(exp_smooth, 0, 5, 50) AS bar
FROM
    (
    SELECT
    (number = 0) OR (number >= 25) AS value,
    number AS time,
    exponentialTimeDecayedAvg(10)(value, time) OVER (ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS exp_smooth
    FROM numbers(50)
    )
