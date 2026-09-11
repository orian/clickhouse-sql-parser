
        INSERT INTO promql_row_cost_ts (metric_name, tags, time_series)
        SELECT
            'rowcost_metric',
            map('instance', toString(number), 'dc', toString(number % 8)),
            arrayMap(step -> (toDateTime64(1000000 + step * 15, 3), toFloat64(step % 977)), range(4096))
        FROM numbers_mt(12800)
    