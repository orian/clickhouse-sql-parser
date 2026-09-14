-- Paid time off vs. observance days by year
SELECT
    toYear(date) AS year,
    countIf(isPaidTimeOff) AS paid_days,
    count() AS all_days
FROM publicholidays
GROUP BY year
ORDER BY year;
