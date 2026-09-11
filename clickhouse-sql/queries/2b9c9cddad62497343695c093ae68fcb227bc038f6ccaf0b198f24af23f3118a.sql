
        INSERT INTO promql_tags_transform (metric_name, tags, time_series)
        SELECT
            if(number < 100000, 'foo', 'bar'),
            map('namespace', 'ns' || toString(intDiv(number % 100000, 1000)), 'pod', 'pod' || toString(number % 100000)),
            [(toDateTime64(100, 3), number::Float64)]
        FROM numbers_mt(200000)
    