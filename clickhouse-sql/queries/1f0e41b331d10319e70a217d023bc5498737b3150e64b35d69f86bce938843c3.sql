
        INSERT INTO t_materialize_uniq_v2_f64
        SELECT * FROM t_insert_uniq_v2_f64_src
        SETTINGS materialize_statistics_on_insert = 0
    