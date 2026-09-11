    CREATE TABLE otel_logs_v2
    (
      -- 与 otel_logs 相同的列
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
