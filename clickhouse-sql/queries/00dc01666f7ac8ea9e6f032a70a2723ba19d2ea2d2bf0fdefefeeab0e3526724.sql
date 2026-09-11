
        SELECT timeSeriesThrowDuplicateSeriesIf(materialize(toUInt8(0)), toUInt64(0))
        FROM numbers(10000000)
        SETTINGS max_threads = 1, max_block_size = 500000
        FORMAT Null
    