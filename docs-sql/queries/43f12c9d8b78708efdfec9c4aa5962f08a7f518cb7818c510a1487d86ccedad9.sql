CREATE TABLE otel_logs_30_01_2025 AS otel_logs
PRIMARY KEY (SeverityNumber, ServiceName, TimestampTime)
ORDER BY (SeverityNumber, ServiceName, TimestampTime)
