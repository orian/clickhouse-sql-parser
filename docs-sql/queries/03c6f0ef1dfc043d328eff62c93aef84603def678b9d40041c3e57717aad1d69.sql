        SELECT
            ServiceName,
            key,
            value
        FROM otel_traces
        ARRAY JOIN
            mapKeys(ResourceAttributes)  AS key,
            mapValues(ResourceAttributes) AS value
        WHERE ServiceName = 'cartservice'
        LIMIT 20;
