SELECT avg(Duration), max(Duration), count(), quantiles(0.95,0.99)(Duration), toStartOfMinute(Timestamp) as time, ServiceName, StatusCode
FROM otel_traces
GROUP BY time, ServiceName, StatusCode
