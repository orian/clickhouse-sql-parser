SELECT
    timeDiff(toDateTime('2021-12-29'), toDateTime('2022-01-01')) AS time_diff_result,
    dateDiff('second', toDateTime('2021-12-29'), toDateTime('2022-01-01')) AS date_diff_result
