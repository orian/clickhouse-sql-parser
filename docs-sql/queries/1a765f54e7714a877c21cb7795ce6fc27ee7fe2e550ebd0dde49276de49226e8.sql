CREATE TABLE otel_logs_merge
AS otel_logs
ENGINE = Merge(currentDatabase(), 'otel_logs*')
