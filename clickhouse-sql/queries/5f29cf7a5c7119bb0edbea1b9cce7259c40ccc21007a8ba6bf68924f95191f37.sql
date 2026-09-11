
        CREATE TABLE t_materialize_uniq_v2_f64 (c1 Float64)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS auto_statistics_types = 'uniq_v2'
    