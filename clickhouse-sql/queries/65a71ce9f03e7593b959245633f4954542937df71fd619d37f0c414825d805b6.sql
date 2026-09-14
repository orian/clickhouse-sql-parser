
        INSERT INTO t_insert_uniq_v2_f64_src
        SELECT toFloat64(number) / 10.0 FROM numbers(20000000)
    