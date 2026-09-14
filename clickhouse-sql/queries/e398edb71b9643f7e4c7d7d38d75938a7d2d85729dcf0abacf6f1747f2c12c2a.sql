
        SELECT count()
        FROM date_preimage_perf
        WHERE toDate(seen) = toDate('2026-03-02')
        SETTINGS optimize_time_filter_with_preimage = 1
    