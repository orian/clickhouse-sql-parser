SELECT Timestamp, RequestPath, Status, RemoteAddress, UserAgent
FROM otel_logs_v2
WHERE Status = 500
FORMAT `Null`
