SET allow_introspection_functions=1;

-- The arrayMap function allows to process each individual element of the trace array by the demangle function.
-- The result of this processing is shown in the trace_functions column of output.

SELECT
    arrayStringConcat(arrayMap(x -> demangle(addressToSymbol(x)), trace), '\n') AS trace_functions
FROM system.trace_log
LIMIT 1
FORMAT Vertical;
