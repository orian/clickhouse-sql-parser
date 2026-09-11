
        INSERT INTO sort_wide_64
        SELECT rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64(), rand64()
        FROM numbers(800000)
    