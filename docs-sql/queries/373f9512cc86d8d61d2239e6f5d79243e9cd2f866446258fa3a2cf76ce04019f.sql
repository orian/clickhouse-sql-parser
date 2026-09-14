CREATE FUNCTION isBusinessHours AS (ts) ->
toDayOfWeek(ts) BETWEEN 1 AND 5
AND toHour(ts) BETWEEN 9 AND 17;
