    CREATE TABLE otel_logs_v2
    (
      -- mismas columnas que otel_logs
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
