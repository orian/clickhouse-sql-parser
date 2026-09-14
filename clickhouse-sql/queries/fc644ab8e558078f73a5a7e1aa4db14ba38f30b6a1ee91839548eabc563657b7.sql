
        INSERT INTO adss_f32_{dim}
        SELECT
            arrayMap(i -> toFloat32((number * 9 + i) % 1000) * 0.001, range({dim})),
            arrayMap(i -> toFloat32((number * 7 + i + 13) % 1000) * 0.001, range({dim}))
        FROM numbers(100000)
    