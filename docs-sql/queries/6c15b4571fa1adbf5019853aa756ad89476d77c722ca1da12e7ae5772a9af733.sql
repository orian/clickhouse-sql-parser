WITH 'a9241258-a0c4-4776-a00b-e6a1d9bec4a1' AS my_query_id
SELECT
    concat(substring(hostName(), length(hostName()), 1), leftPad(greatest(attribute['clickhouse.thread_id'], attribute['thread_number']), 5, '0')) AS group,
    operation_name,
    start_time_us,
    finish_time_us,
    sipHash64(operation_name) AS color,
    attribute
FROM system.opentelemetry_span_log
WHERE (trace_id IN (
    SELECT trace_id
    FROM system.opentelemetry_span_log
    WHERE (attribute['clickhouse.query_id']) = my_query_id
)) AND (operation_name != 'query') AND (operation_name NOT LIKE 'Query%')
ORDER BY
    hostName() ASC,
    group ASC,
    parent_span_id ASC,
    start_time_us ASC
INTO OUTFILE 'trace.json'
FORMAT JSON
SETTINGS output_format_json_named_tuples_as_objects = 1
