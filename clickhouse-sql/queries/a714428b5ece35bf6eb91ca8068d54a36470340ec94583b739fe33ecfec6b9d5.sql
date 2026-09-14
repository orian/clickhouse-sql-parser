
        INSERT INTO promql_set_operator_presence_mask_ts (metric_name, tags, time_series)
        SELECT
            if(number < 4000, 'foo', 'bar'),
            map('instance', toString(number % 4000)),
            arrayMap(step -> (toDateTime64(100 + step * 10, 3), (sipHash64(number, step) % 1000000) / 1000), range(300))
        FROM numbers_mt(8000)
        UNION ALL
        SELECT
            if(number < 4000, 'foo_sparse', 'bar_sparse'),
            map(
                'instance', toString(intDiv(number, 4) % 1000),
                'replica', toString(number % 4)),
            arrayMap(
                step -> (toDateTime64(100 + step * 10, 3), (sipHash64(number, step) % 1000000) / 1000),
                arrayFilter(step -> sipHash64(number, step) % 2 = 0, range(300)))
        FROM numbers_mt(8000)
    