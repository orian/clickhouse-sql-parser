WITH traces AS (
    SELECT * FROM system.trace_log
    WHERE event_date >= today() AND trace_type = 'Instrumentation' AND handler = 'profile'
    ORDER BY event_time, entry_type
)
SELECT
    format(
        '{{"traceEvents": [{}\n]}}',
        arrayStringConcat(
            groupArray(
                format(
                    '\n{{"name": "{}", "cat": "clickhouse", "ph": "{}", "ts": {}, "pid": 1, "tid": {}, "args": {{"query_id": "{}", "cpu_id": {}, "stack": [{}]}}}},',
                    function_name,
                    if(entry_type = 0, 'B', 'E'),
                    timestamp_ns/1000,
                    toString(thread_id),
                    query_id,
                    cpu_id,
                    arrayStringConcat(arrayMap((x, y) -> concat('"', x, ': ', y, '", '), lines, symbols))
                )
            )
        )
    )
FROM traces;
