
        CREATE TABLE lc_many_marks_point_lookups
        (
            k String,
            ts UInt32,
            lc1 LowCardinality(String),
            lc2 LowCardinality(String),
            lc3 LowCardinality(String),
            lc4 LowCardinality(String),
            lc5 LowCardinality(String),
            lc6 LowCardinality(String),
            v UInt64
        )
        ENGINE = MergeTree
        ORDER BY (k, ts)
        SETTINGS index_granularity = 8, min_bytes_for_wide_part = 0, min_rows_for_wide_part = 0
    