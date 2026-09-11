
        CREATE TABLE ttl_group_by (k UInt64, d DateTime, v1 UInt64, v2 UInt64, v3 UInt64, v4 UInt64, v5 UInt64, v6 UInt64, v7 UInt64, v8 UInt64)
        ENGINE = MergeTree ORDER BY k
        TTL d + INTERVAL 1 SECOND GROUP BY k
            SET d = max(d), v1 = sum(v1), v2 = sum(v2), v3 = sum(v3), v4 = sum(v4), v5 = sum(v5), v6 = sum(v6), v7 = sum(v7), v8 = sum(v8)
        SETTINGS min_bytes_for_wide_part = 0, max_number_of_merges_with_ttl_in_pool = 0
    