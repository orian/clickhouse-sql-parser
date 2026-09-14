-- The arrayMap function in this example processing each individual element of the trace array by the addressToLine function.
-- The result of this processing is seen in the trace_source_code_lines column of output.

SELECT
    arrayStringConcat(arrayMap(x -> addressToLine(x), trace), '\n') AS trace_source_code_lines
FROM system.trace_log
LIMIT 1
FORMAT Vertical;
