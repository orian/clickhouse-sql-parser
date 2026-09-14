
        INSERT INTO t_insert_uniq_v2_i64_src
        SELECT toInt64(number) FROM numbers(20000000)
    