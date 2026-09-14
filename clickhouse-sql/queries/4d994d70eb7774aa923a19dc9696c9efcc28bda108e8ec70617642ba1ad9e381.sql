
        CREATE TABLE lazy_mat_final_limit (k UInt64, v UInt64, payload String)
        ENGINE = ReplacingMergeTree
        ORDER BY k
    