    CREATE TABLE otel_logs_v2
    (
      -- نفس الأعمدة الموجودة في otel_logs
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
