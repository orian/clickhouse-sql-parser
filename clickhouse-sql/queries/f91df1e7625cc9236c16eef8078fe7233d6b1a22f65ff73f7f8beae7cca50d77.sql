
        CREATE TABLE t_insert_uniq_v2_f64_off (c1 Float64)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS auto_statistics_types = 'uniq_v2'
    