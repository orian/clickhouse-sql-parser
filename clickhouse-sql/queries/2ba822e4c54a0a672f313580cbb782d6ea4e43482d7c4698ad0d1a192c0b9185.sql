
        CREATE TABLE array_distance_mixed_{array_size}
        (
            id UInt64,
            u8 Array(UInt8),
            f64 Array(Float64)
        )
        ENGINE = Memory
    