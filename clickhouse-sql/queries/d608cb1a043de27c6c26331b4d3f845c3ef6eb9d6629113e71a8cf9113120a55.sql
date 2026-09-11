
        WITH (WatchID % 2 == 0) AS predicate
        SELECT
            {function}If(WatchID, predicate),
            {function}If(CounterID, predicate),
            {function}If(ClientIP, predicate),
            {function}If(IPNetworkID, predicate),
            {function}If(SearchEngineID, predicate)
        FROM hits_100m_single
        GROUP BY intHash32(UserID) % {group_scale}
        FORMAT Null
    