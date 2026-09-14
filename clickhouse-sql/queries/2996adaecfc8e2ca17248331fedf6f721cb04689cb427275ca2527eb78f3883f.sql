
        INSERT INTO array_distance_const_exact_chunks
        SELECT
            arrayMap(i -> toFloat32(((number * 3 + i) % 251) + 1), range(16)) AS w_f32_16,
            arrayMap(i -> toFloat64(((number * 5 + i) % 251) + 1), range(8)) AS w_f64_8,
            arrayMap(i -> toBFloat16(((number * 7 + i) % 251) + 1), range(32)) AS w_bf16_32
        FROM numbers(1000000)
    