
        CREATE TABLE ttl_group_by_src (k UInt64, d DateTime, v1 UInt64, v2 UInt64, v3 UInt64, v4 UInt64, v5 UInt64, v6 UInt64, v7 UInt64, v8 UInt64)
        ENGINE = MergeTree ORDER BY k
        SETTINGS min_bytes_for_wide_part = 0
    