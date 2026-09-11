WITH
    toDate('2024-06-15') AS date,
    toIntervalYear(1) AS interval_to_year
SELECT date + interval_to_year AS result
