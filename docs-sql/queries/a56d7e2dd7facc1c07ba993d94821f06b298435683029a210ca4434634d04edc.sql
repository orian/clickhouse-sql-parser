WITH
    toDate('2024-01-01') AS date,
    toDateTime('2024-01-01 00:00:00') AS date_time,
    '2024-01-01 00:00:00' AS date_time_string
SELECT
    subtractMonths(date, 1) AS subtract_months_with_date,
    subtractMonths(date_time, 1) AS subtract_months_with_date_time,
    subtractMonths(date_time_string, 1) AS subtract_months_with_date_time_string
