
        SELECT 1 FROM test.hits GROUP BY EventTime HAVING
        sequenceMatch('(?1)(?t<1)(?2)')(
        EventTime, Age BETWEEN 20 AND 30, Age BETWEEN 35 AND 50, Age >= 0)
        FORMAT Null
    