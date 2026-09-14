SELECT
    count(),
    arrayStringConcat(arrayMap((symbol, line) -> concat(symbol, '\n    ', line), any(symbols), any(lines)), '\n') AS sym
FROM system.trace_log
WHERE (query_id = '<query_id>') AND (event_date = today())
GROUP BY trace
ORDER BY count() DESC
LIMIT 10
