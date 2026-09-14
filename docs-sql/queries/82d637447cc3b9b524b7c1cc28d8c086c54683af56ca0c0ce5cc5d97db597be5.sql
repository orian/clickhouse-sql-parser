SELECT *
FROM otel_traces
WHERE (ServiceName = 'accountingservice') AND (TraceId = '4512e822ca3c0c68bbf5d4a263f9943d')
FORMAT Null
SETTINGS use_query_condition_cache = 0, use_skip_indexes = 0;
