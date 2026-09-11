
        INSERT INTO perf_avg(num)
        SELECT toUInt64(UserID / (WatchID + 1)  * 1000000)
        FROM hits_100m_single
        LIMIT 50000000
    