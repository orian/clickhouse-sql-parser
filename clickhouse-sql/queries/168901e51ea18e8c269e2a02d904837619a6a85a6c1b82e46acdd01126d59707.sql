
        INSERT INTO sort_wide_16
        SELECT rand64() AS k1, rand64() AS k2,
            rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64(),
            rand64(), rand64(), rand64(), rand64()
        FROM numbers(500000)
    