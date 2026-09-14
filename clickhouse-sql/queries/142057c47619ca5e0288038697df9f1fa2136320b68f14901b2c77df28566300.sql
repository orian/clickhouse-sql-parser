
        CREATE TABLE t_insert_uniq_v2_i64_on (c1 Int64)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS auto_statistics_types = 'uniq_v2'
    