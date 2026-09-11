
        INSERT INTO promql_topk_ts (metric_name, tags, time_series)
        SELECT
            'load',
            map('instance', toString(number), 'dc', toString(number % 10)),
            arrayMap(step -> (toDateTime64(100 + step * 10, 3), (sipHash64(number, step) % 1000000) / 1000), range(300))
        FROM numbers_mt(4000)
    