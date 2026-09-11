
        SELECT
            {function}(WatchID),
            {function}(CounterID),
            sum(toUInt256(ClientIP)),
            {function}(ClientIP),
            {function}(IPNetworkID),
            {function}(SearchEngineID)
        FROM hits_100m_single
        GROUP BY intHash32(UserID) % {group_scale}
        FORMAT Null
    