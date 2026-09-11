    CREATE TABLE otel_logs_temp AS otel_logs
    PRIMARY KEY (SeverityText, ServiceName, Timestamp)
    ORDER BY (SeverityText, ServiceName, Timestamp)
