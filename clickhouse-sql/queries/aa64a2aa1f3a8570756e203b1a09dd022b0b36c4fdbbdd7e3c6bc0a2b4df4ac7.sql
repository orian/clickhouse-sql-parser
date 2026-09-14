
        INSERT INTO adss_mixed_{dim}
        SELECT
            arrayMap(i -> toUInt8((number + i) % 256), range({dim})),
            arrayMap(i -> toFloat64((number * 3 + i) % 997) * 0.5, range({dim}))
        FROM numbers(100000)
    