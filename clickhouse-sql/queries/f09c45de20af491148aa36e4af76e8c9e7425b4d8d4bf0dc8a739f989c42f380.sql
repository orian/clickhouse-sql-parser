
        CREATE TABLE permute_cache_wide
        (
            c1 String,
            c2 String,
            c3 String,
            c4 String,
            c5 String,
            ts DateTime64(6),
            payload String,
            value UInt64,
            INDEX idx_c1 (c1) TYPE bloom_filter,
            INDEX idx_c2 (c2) TYPE bloom_filter,
            INDEX idx_c3 (c3) TYPE bloom_filter,
            INDEX idx_c4 (c4) TYPE bloom_filter,
            INDEX idx_c5 (c5) TYPE bloom_filter,
            INDEX idx_ts (ts) TYPE minmax
        )
        ENGINE = MergeTree
        ORDER BY (c1, c2, c3, c4, c5, ts)
        SETTINGS min_rows_for_wide_part = 0, min_bytes_for_wide_part = 0
    