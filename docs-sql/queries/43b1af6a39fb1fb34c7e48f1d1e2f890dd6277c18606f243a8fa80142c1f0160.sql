    INSERT INTO otel_traces_backfill
    SELECT
        Timestamp,
        ServiceName,
        StatusCode,
        Duration
    FROM otel_traces
    WHERE Timestamp < (
        SELECT min(Timestamp) FROM otel_traces_1m
    );
