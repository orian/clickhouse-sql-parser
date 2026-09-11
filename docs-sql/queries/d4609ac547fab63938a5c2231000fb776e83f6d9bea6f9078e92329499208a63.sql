WITH
    toDate('2025-06-15') AS date,
    toIntervalQuarter(1) AS interval_to_quarter
SELECT date + interval_to_quarter AS result
