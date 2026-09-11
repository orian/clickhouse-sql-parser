
        CREATE TABLE optimized_select_final (t DateTime, x Int32)
        ENGINE = ReplacingMergeTree()
        PARTITION BY toYYYYMM(t) ORDER BY x
    