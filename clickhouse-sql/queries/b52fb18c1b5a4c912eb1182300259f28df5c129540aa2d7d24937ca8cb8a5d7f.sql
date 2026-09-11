
        SELECT
            {function}If(value_1, predicate),
            {function}If(value_2, predicate),
            sumIf(toUInt256(value_3), predicate),
            {function}If(value_3, predicate),
            {function}If(value_4, predicate),
            {function}If(value_5, predicate)
        FROM {table}
        GROUP BY key
        FORMAT Null
    