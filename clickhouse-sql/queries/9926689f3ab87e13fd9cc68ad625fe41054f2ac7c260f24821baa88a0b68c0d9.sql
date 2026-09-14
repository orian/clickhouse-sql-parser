
        CREATE TABLE sort_lc (
            lc_col LowCardinality(String),
            uint_col UInt64
        ) ENGINE = MergeTree ORDER BY tuple()
    