
        SELECT 1 FROM test.hits WHERE RefererCategories != [] GROUP BY ClientIP, RequestNum HAVING
        sequenceMatch('(?1)(?3)(?1)(?3)')(
        EventTime, hasAny(RefererCategories, [9]), hasAny(RefererCategories, [3849, 2, 3, 4, 5, 6, 7]), hasAll(RefererCategories, [1, 9]), hasAny(RefererCategories, [1, 2326, 5496]))
        FORMAT Null
    