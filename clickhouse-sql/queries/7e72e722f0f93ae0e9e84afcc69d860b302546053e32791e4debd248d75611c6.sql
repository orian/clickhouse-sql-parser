
        CREATE TABLE replacing_final_non_intersecting (d DateTime, c1 UInt64, c2 String, c3 LowCardinality(String))
        ENGINE = ReplacingMergeTree()
        ORDER BY d
    