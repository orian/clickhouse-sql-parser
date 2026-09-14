
        select CounterID, UserID, count(*) user_hits
        from hits_100m_single
        where CounterID < 10000
        group by CounterID, UserID
        order by user_hits desc
        limit 10 by CounterID
        format Null
    