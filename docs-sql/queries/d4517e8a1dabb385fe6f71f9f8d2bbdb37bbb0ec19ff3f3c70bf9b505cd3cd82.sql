    SELECT
        ServiceName,
        SpanName,
        ResourceAttributes['host.name']             AS host,
        ResourceAttributes['k8s.pod.name']          AS pod,
        ResourceAttributes['deployment.environment'] AS env
    FROM otel_traces
    LIMIT 10;
