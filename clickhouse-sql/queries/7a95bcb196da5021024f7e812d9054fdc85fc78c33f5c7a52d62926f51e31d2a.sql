
        CREATE TABLE test_limit_by_{cardinality}
        (a UInt32, b UInt32, v UInt64)
        ENGINE = MergeTree
        ORDER BY (a, b)
    