
        CREATE TABLE dynamic_compact_read
        (
            key Dynamic,
            blob1 Dynamic,
            blob2 Dynamic,
            grp_kind String,
            tenant String,
            flag_a Bool,
            flag_b Bool,
            rid String
        )
        ENGINE = MergeTree
        ORDER BY (flag_a, flag_b, grp_kind, tenant, rid)
        SETTINGS min_bytes_for_wide_part = 9223372036854775807
    