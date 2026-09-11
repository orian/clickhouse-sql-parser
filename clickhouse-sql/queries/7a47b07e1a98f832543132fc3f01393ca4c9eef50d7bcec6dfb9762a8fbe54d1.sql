
        SELECT
            WatchID
        FROM hits_100m_single
        ORDER BY WatchID ASC, CounterID DESC, ClientIP ASC, IPNetworkID DESC, SearchEngineID ASC
        LIMIT 2000000
        FORMAT Null
    