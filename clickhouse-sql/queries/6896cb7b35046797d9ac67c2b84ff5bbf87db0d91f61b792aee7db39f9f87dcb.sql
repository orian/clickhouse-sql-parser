
        select *
        from (
            select CounterID, UserID, count(*) user_hits,
                count()
                    over (partition by CounterID order by user_hits desc
                        rows unbounded preceding)
                    user_rank
            from hits_100m_single
            where CounterID < 10000
            group by CounterID, UserID
        )
        where user_rank <= 10
        format Null
    