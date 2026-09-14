
        CREATE TABLE test_strings_{cardinality}
        (
            col1 LowCardinality(String),
            col2 LowCardinality(String)
        )
        ENGINE = MergeTree
        ORDER BY col1
    