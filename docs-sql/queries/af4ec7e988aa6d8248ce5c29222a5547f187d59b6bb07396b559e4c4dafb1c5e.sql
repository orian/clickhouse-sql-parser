SELECT *
FROM system.mutations
WHERE database = 'otel'
  AND table = 'otel_traces'
ORDER BY create_time DESC;
