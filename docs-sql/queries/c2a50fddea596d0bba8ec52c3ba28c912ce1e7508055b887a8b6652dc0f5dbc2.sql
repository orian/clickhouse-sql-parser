CREATE TABLE hits (EventTime DateTime('UTC')) ENGINE = Memory;

-- One row per page view, spread over the hours of a day.
INSERT INTO hits
SELECT toDateTime('2026-01-01 00:00:00', 'UTC') + toIntervalHour(hour)
FROM (
    SELECT
        number AS hour,
        [293, 181, 115, 85, 69, 78, 113, 171, 278, 391, 458, 494, 510, 523, 540, 528, 539, 524, 506, 521, 522, 542, 494, 400][number + 1] AS count
    FROM numbers(24)
)
ARRAY JOIN range(count);

SELECT
toHour(EventTime) AS h,
count() AS c,
bar(c, 0, 600, 20) AS bar
FROM hits
GROUP BY h
ORDER BY h ASC
