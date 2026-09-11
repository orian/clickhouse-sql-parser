
        CREATE TABLE t_uniq_v2_float64(c1 Float64 STATISTICS(uniq_v2))
        ENGINE = MergeTree ORDER BY tuple() SETTINGS auto_statistics_types = '';
    