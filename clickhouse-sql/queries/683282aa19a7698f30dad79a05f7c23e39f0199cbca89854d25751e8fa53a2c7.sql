
        SELECT 1 FROM test.hits WHERE RefererCategories != [] GROUP BY ClientIP, RequestNum HAVING
        sequenceMatch('(?1)(?2)(?1)(?2)(?1)')(
        EventTime, hasAny(RefererCategories, [3849, 2, 3, 4, 5, 6, 7]), hasAny(RefererCategories, [1, 2]))
        FORMAT Null
    