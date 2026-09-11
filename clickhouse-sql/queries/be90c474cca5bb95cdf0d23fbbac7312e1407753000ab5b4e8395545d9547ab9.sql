
        INSERT INTO t_insert_uniq_v2_i64_off
        SELECT * FROM t_insert_uniq_v2_i64_src
        SETTINGS materialize_statistics_on_insert = 0
    