    CREATE TABLE otel_logs_v2
    (
      -- те же столбцы, что и в otel_logs
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
