SELECT
    toStartOfInterval(timestamp, toIntervalMillisecond(100)) AS bucket,
    count() AS count
FROM MidJourney.images
WHERE (timestamp >= {start:String}) AND (timestamp <= {end:String})
GROUP BY ALL
ORDER BY bucket ASC
WITH FILL
STEP toIntervalMillisecond(100);
