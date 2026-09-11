
        SELECT
            {function}(WatchID),
            {function}(CounterID),
            {function}(ClientIP),
            {function}(IPNetworkID),
            {function}(SearchEngineID)
        FROM hits_100m_single
        GROUP BY intHash32(UserID) % {group_scale}
        FORMAT Null
    