SELECT
    toStartOfInterval(timestamp, toIntervalMillisecond(100)) AS bucket,
    count() AS count,
    sum(count) OVER (ORDER BY bucket) AS cumulative,
    bar(cumulative, 0, 10, 10) AS barChart
FROM MidJourney.images
WHERE (timestamp >= {start:String}) AND (timestamp <= {end:String})
GROUP BY ALL
ORDER BY bucket ASC
WITH FILL
FROM toDateTime64({start:String}, 3)
TO toDateTime64({end:String}, 3) + INTERVAL 100 millisecond
STEP toIntervalMillisecond(100)
INTERPOLATE (cumulative, barChart);
