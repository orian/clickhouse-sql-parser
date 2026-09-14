CREATE TABLE otel_logs_v2
(
  -- same columns as otel_logs
)
ENGINE = MergeTree
ORDER BY (TenantId, ServiceName, Timestamp);
