
        SELECT * FROM (SELECT 1 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 2 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 3 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 4 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 5 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 6 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 7 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 8 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 9 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
        UNION ALL
        SELECT * FROM (SELECT 10 AS k, id, L2DistanceTransposed(vec, arrayMap(i -> (i + k) * 2, range({vector_size})), {p}) AS dist FROM tab_{type} ORDER BY dist LIMIT 10)
    