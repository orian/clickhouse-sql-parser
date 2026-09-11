
        select * from
            (select EventTime,
                count(*) over (partition by
                    floor((toUInt32(EventTime) + 10 + 1) / 20)) as c
            from hits_10m_single)
        format Null
    