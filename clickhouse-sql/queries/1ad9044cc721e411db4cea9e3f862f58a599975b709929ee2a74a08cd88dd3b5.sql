
        CREATE TABLE distinct_low_cardinality_lc
        (
            id UInt32,
            lc LowCardinality(String)
        )
        ENGINE = MergeTree
        ORDER BY id;
    