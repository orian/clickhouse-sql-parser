
        INSERT INTO perf_lc_str (str)
        SELECT concat('asdf', toString(number % 10000))
        FROM numbers(100000000)
    