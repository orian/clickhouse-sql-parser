    SELECT
        toString(trace_id) AS traceID,
        toString(span_id) AS spanID,
        if(toString(parent_span_id)='0', '', toString(parent_span_id)) AS parentSpanID,
        'ClickHouse' AS serviceName,
        operation_name AS operationName,
        start_time_us/1000000 AS startTime,
        (finish_time_us - start_time_us)/1000 AS duration,
        arrayMap(key -> map('key', key, 'value', attribute[key]), mapKeys(attribute)) AS serviceTags
    FROM system.opentelemetry_span_log
    WHERE trace_id = '68a14b27-a61f-596d-3746-2b03d2530e42' ORDER BY startTime ASC
