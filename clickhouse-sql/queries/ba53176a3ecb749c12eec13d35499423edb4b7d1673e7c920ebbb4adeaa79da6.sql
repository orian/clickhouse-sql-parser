
        SELECT
            {function}(value_1),
            {function}(value_2),
            sum(toUInt256(value_3)),
            {function}(value_3)
        FROM {table}
        GROUP BY key
        FORMAT Null
    