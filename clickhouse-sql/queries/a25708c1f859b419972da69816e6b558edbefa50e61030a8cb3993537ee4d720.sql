
        CREATE TABLE sort_mixed (
            uint_col UInt64,
            str_col String
        ) ENGINE = MergeTree ORDER BY tuple()
    