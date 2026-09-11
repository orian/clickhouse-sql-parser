
        CREATE TABLE t_minmax_stats_float64(c1 Float64 STATISTICS(minmax), c2 Float64 STATISTICS(minmax), c3 Float64 STATISTICS(minmax), c4 Float64 STATISTICS(minmax), c5 Float64 STATISTICS(minmax), c6 Float64 STATISTICS(minmax), c7 Float64 STATISTICS(minmax), c8 Float64 STATISTICS(minmax), c9 Float64 STATISTICS(minmax), c10 Float64 STATISTICS(minmax))
        ENGINE = MergeTree ORDER BY tuple() SETTINGS auto_statistics_types = '';
    