INSERT INTO nullint64
                SELECT IF(rand() % 5 == 0, NULL::Nullable(UInt64), cityHash64(number)) AS x
                FROM numbers_mt(200000000)
                SETTINGS max_threads = 8
    