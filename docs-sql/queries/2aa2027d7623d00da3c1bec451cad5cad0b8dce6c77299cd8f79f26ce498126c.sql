CREATE TABLE otel_logs_23_01_2025 AS otel_logs
PRIMARY KEY (SeverityText, ServiceName, Timestamp)
ORDER BY (SeverityText, ServiceName, Timestamp)
