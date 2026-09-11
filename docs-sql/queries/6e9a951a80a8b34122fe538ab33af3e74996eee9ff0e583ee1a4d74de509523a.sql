SELECT
    toDate('2022-01-01') AS e,
    toDate('2021-12-29') AS s,
    dateDiff('day', s, e) AS day_diff,
    dateDiff('month', s, e) AS month_diff,
    dateDiff('year', s, e) AS year_diff
