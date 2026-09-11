
        CREATE TABLE distinct_high_cardinality_lc
        (
            id UInt32,
            lc LowCardinality(String)
        )
        ENGINE = MergeTree
        ORDER BY id;
    