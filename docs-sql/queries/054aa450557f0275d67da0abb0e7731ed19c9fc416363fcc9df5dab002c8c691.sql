WITH
    toDate('2025-06-15') AS date,
    toIntervalMinute(12) AS interval_to_minutes
SELECT date + interval_to_minutes AS result
