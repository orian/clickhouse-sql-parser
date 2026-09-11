
        INSERT INTO perf_lc_num (num)
        SELECT toUInt8(number)
        FROM numbers(10000000)
    