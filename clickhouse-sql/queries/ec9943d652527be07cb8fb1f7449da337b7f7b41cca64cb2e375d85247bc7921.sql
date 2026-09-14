
        INSERT INTO t_insert_uniq_v2_i64_on
        SELECT * FROM t_insert_uniq_v2_i64_src
        SETTINGS materialize_statistics_on_insert = 1
    