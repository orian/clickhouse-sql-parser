
        SELECT 1 FROM hits_10m_single GROUP BY EventTime HAVING
        sequenceMatch('(?1)(?t<1)(?2)')(
        EventTime, Age >= 0, Age = -1)
        FORMAT Null
    