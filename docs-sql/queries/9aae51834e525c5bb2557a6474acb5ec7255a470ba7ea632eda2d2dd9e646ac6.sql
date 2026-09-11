    CREATE TABLE otel_logs_v2
    (
      -- mêmes colonnes que otel_logs
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
