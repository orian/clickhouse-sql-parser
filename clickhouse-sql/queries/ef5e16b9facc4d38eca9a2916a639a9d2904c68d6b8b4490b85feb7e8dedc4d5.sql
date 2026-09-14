
        INSERT INTO perf_lc_fixed_str_groupby SELECT ('number key ' || toString(number % 400)) AS a, ('number key ' || toString(number % 20)) AS b FROM numbers(30000000)
    