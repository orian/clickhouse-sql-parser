
        CREATE TABLE t_minmax_stats_uint64(c1 UInt64 STATISTICS(minmax), c2 UInt64 STATISTICS(minmax), c3 UInt64 STATISTICS(minmax), c4 UInt64 STATISTICS(minmax), c5 UInt64 STATISTICS(minmax), c6 UInt64 STATISTICS(minmax), c7 UInt64 STATISTICS(minmax), c8 UInt64 STATISTICS(minmax), c9 UInt64 STATISTICS(minmax), c10 UInt64 STATISTICS(minmax))
        ENGINE = MergeTree ORDER BY tuple() SETTINGS auto_statistics_types = '';
    