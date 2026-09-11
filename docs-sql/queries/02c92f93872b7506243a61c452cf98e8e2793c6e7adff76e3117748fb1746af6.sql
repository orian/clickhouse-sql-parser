    CREATE TABLE otel_logs_v2
    (
      -- otel_logs와 동일한 컬럼
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
