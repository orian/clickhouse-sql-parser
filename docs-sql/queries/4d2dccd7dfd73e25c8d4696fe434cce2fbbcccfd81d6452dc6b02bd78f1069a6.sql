    SELECT
        ServiceName,
        SpanName,
        mapContains(ResourceAttributes, 'k8s.node.name') AS has_node_info
    FROM otel_traces
    LIMIT 10;
