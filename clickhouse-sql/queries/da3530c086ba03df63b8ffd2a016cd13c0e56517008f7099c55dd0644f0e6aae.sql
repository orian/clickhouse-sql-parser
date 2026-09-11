
        CREATE TABLE lazy_final_limit (k UInt64, v UInt64, c UInt8, payload String)
        ENGINE = ReplacingMergeTree
        ORDER BY k
    