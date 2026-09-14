SET allow_introspection_functions=1;

-- The arrayMap function allows to process each individual element of the trace array by the addressToSymbols function.
-- The result of this processing is shown in the trace_symbols column of output.

SELECT
    arrayStringConcat(arrayMap(x -> addressToSymbol(x), trace), '\n') AS trace_symbols
FROM system.trace_log
LIMIT 1
FORMAT Vertical;
