    CREATE TABLE otel_logs_v2
    (
      -- mesmas colunas de otel_logs
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
