
        SELECT *
        FROM bench
        GROUP BY key
        SETTINGS optimize_aggregation_in_order = 1, max_threads = 16
        FORMAT Null
    