
        CREATE TABLE optimized_select_final (d Date, key UInt64, value String)
        ENGINE = ReplacingMergeTree()
        PARTITION BY toYYYYMM(d) ORDER BY key
    