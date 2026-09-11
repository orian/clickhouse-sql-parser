EXPLAIN ESTIMATE
SELECT
    toStartOfHour(Timestamp) AS hour,
    ServiceName,
    avgMerge(avg__Duration) AS avg__Duration
FROM otel_v2.otel_traces_1m
GROUP BY
    hour,
    ServiceName
ORDER BY hour DESC
