    CREATE TABLE otel_logs_v2
    (
      -- otel_logs と同じカラム
    )
    ENGINE = MergeTree
    ORDER BY (TenantId, ServiceName, Timestamp);
