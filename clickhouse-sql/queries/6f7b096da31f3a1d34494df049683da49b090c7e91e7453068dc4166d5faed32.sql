
        CREATE TABLE variant_compact
        (
            id UInt64, v Variant(String, UInt64)
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 1000000000000
    