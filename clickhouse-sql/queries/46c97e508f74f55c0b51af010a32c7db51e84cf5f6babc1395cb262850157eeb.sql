
        INSERT INTO t_materialize_uniq_v2_i64
        SELECT * FROM t_insert_uniq_v2_i64_src
        SETTINGS materialize_statistics_on_insert = 0
    