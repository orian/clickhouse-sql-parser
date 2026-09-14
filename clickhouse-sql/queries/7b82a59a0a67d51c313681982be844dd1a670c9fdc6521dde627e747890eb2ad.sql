
        SELECT timeSeriesThrowDuplicateSeriesIf(materialize(CAST('-0.0', 'Float64')), toUInt64(0))
        FROM numbers(10000000)
        SETTINGS max_threads = 1, max_block_size = 500000
        FORMAT Null
    