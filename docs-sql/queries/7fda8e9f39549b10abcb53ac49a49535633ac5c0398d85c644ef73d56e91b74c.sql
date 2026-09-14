SELECT
    toStartOfMinute(Timestamp),
    ServiceName,
    StatusCode,
    count() AS count,
    avg(Duration),
    max(Duration),
    quantiles(0.95, 0.99)(Duration)
FROM otel_traces
GROUP BY
    time,
    ServiceName,
    StatusCode
