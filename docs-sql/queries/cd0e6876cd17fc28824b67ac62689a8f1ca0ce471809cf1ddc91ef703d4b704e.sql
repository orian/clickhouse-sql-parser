SELECT
    count(),
    arrayStringConcat(arrayMap(x -> concat(demangle(addressToSymbol(x)), '\n    ', addressToLine(x)), trace), '\n') AS sym
FROM system.stack_trace
GROUP BY trace
ORDER BY count() DESC
LIMIT 10
FORMAT Vertical
SETTINGS allow_introspection_functions = 1;
