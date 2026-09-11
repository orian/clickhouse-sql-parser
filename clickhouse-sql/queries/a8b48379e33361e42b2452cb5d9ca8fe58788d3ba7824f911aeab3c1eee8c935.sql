
        INSERT INTO promql_pk_range_ts (metric_name, tags, time_series)
        SELECT
            'pkr_metric',
            map('instance', toString(number), 'dc', toString(number % 8)),
            arrayMap(step -> (toDateTime64(1000000 + step * 15, 3), (sipHash64(number, step) % 1000000) / 1000), range(256))
        FROM numbers_mt(20000)
    