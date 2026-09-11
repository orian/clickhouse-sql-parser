WITH
    toDate('2024-01-01') AS date,
    toDateTime('2024-01-01 00:00:00') AS date_time,
    '2024-01-01 00:00:00' AS date_time_string
SELECT
    subtractQuarters(date, 1) AS subtract_quarters_with_date,
    subtractQuarters(date_time, 1) AS subtract_quarters_with_date_time,
    subtractQuarters(date_time_string, 1) AS subtract_quarters_with_date_time_string
