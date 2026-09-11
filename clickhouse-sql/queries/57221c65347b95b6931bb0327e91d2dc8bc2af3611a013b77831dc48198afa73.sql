
        SELECT 1 FROM test.hits GROUP BY EventTime HAVING
        sequenceMatch('(?1)(?t<1)(?2)')(
        EventTime, Age >= 0, Age = -1)
        FORMAT Null
    