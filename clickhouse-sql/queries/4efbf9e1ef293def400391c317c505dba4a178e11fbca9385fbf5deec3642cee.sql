
        CREATE TABLE t_basic_stats_float64(c1 Float64 STATISTICS(basic), c2 Float64 STATISTICS(basic), c3 Float64 STATISTICS(basic), c4 Float64 STATISTICS(basic), c5 Float64 STATISTICS(basic), c6 Float64 STATISTICS(basic), c7 Float64 STATISTICS(basic), c8 Float64 STATISTICS(basic), c9 Float64 STATISTICS(basic), c10 Float64 STATISTICS(basic))
        ENGINE = MergeTree ORDER BY tuple() SETTINGS auto_statistics_types = '';
    