SET allow_introspection_functions=1;

-- The arrayJoin function will split array to rows

SELECT
    ta, addressToLineWithInlines(arrayJoin(trace) AS ta)
FROM system.trace_log
WHERE
    query_id = '5e173544-2020-45de-b645-5deebe2aae54';
