INSERT INTO nullfloat32
                SELECT IF(rand() % 5 == 0, NULL::Nullable(Float32), toFloat32(number)) as x
                FROM numbers_mt(200000000)
                SETTINGS max_threads = 8
    