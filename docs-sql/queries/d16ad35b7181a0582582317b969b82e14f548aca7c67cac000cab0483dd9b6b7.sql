WITH 60e9 AS min_age_ns
SELECT format('{} {}',
    if(s.est_old_objects > 0, s.stack, format('[unattributed];arena_{}_class_{}', toString(b.arena), toString(b.size))),
    toString(toUInt64(if(s.est_old_objects > 0, b.waste * s.est_old_objects / sum(s.est_old_objects) OVER (PARTITION BY b.arena, b.index), b.waste))))
FROM system.jemalloc_arena_bins AS b
LEFT JOIN
(
    SELECT
        arena,
        size_class,
        sum(weight) AS est_old_objects,
        arrayStringConcat(arrayMap(a -> demangle(addressToSymbol(a)), arrayReverse(trace)), ';') AS stack
    FROM system.jemalloc_sampled_allocations
    WHERE age_ns > min_age_ns
    GROUP BY arena, size_class, trace
) AS s ON s.size_class = b.index AND s.arena = b.arena
WHERE b.large = 0 AND b.purpose = '' AND b.waste > 0
SETTINGS allow_introspection_functions=1
