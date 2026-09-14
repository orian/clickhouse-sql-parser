
        CREATE TABLE lazy_final_perf
        (
            part UInt8,
            key UInt64,
            version UInt64,
            category UInt8,
            payload String
        )
        ENGINE = ReplacingMergeTree(version)
        PARTITION BY part
        ORDER BY key
        SETTINGS index_granularity = 256
    