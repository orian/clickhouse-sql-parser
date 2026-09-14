INSERT INTO nullfloat64_stats
                SELECT IF(rand() % 5 == 0, NULL::Nullable(Float64), toFloat64(cityHash64(number) % 1000000)) AS x
                FROM numbers_mt(200000000)
                SETTINGS max_threads = 8
    