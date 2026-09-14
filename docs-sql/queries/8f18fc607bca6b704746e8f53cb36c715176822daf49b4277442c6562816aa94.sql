    SELECT DISTINCT arrayJoin(mapKeys(ResourceAttributes)) AS key
    FROM otel_traces
    ORDER BY key;
