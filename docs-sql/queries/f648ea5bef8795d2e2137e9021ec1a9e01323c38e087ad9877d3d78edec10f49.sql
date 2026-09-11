    SELECT
        Timestamp,
        ServiceName,
        SpanName,
        ResourceAttributes['host.name']    AS host,
        ResourceAttributes['k8s.pod.name'] AS pod,
        SpanAttributes['error.type']       AS error_type,
        SpanAttributes['error.message']    AS error_message
    FROM otel_traces
    WHERE StatusCode = 'STATUS_CODE_ERROR';
