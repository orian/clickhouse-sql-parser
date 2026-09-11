
        INSERT INTO array_distance_mixed_{array_size}
        SELECT
            number,
            arrayMap(i -> toUInt8((number + i) % 256), range({array_size})),
            arrayMap(i -> toFloat64((number * 3 + i) % 997) * 0.5, range({array_size}))
        FROM numbers(100000)
    