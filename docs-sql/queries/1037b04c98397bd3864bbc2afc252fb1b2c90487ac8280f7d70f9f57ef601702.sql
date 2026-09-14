        SELECT
            Timestamp,
            SpanName,
            Duration / 1e6 AS duration_ms
        FROM otel_traces
        WHERE ResourceAttributes['service.name'] = 'cartservice'
        ORDER BY Timestamp
        LIMIT 10;
