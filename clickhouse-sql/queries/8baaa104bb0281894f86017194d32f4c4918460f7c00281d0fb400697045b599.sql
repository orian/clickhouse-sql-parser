
        SELECT count()
        FROM date_preimage_perf
        WHERE toStartOfMonth(d) = toDate('2026-03-01')
        SETTINGS optimize_time_filter_with_preimage = 1
    