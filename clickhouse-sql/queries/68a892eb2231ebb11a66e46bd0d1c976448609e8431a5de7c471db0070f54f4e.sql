
        INSERT INTO array_distance_f32_{array_size}
        SELECT
            number,
            arrayMap(i -> toFloat32((number * 9 + i) % 1000) * 0.001, range({array_size})),
            arrayMap(i -> toFloat32((number * 7 + i + 13) % 1000) * 0.001, range({array_size}))
        FROM numbers(100000)
    