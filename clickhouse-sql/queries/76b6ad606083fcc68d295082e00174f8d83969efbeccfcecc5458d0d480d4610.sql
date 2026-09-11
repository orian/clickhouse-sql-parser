
        CREATE TABLE t_uniq_v2_uint64(c1 UInt64 STATISTICS(uniq_v2))
        ENGINE = MergeTree ORDER BY tuple() SETTINGS auto_statistics_types = '';
    