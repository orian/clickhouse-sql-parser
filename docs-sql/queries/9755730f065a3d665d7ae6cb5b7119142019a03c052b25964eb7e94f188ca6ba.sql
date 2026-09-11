    SELECT
        ServiceName,
        mapFilter((k, v) -> k LIKE 'k8s.%', ResourceAttributes) AS k8s_attrs
    FROM otel_traces
    WHERE mapContains(ResourceAttributes, 'k8s.pod.name')
    LIMIT 10;
