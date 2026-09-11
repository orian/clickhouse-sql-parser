WITH 60e9 AS min_age_ns
SELECT
    b.arena,
    b.size AS reg_size,
    b.waste AS class_waste,
    s.est_old_objects,
    s.stack
FROM system.jemalloc_arena_bins AS b
INNER JOIN
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
WHERE b.large = 0 AND b.purpose = '' AND b.waste > 1048576
ORDER BY b.waste DESC, s.est_old_objects DESC
SETTINGS allow_introspection_functions=1
