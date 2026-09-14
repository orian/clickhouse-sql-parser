
        CREATE TABLE group_by_low_cardinality_array_join
        (
            stack_array Array(LowCardinality(UInt128)),
            stack_flat LowCardinality(UInt128),
            value UInt64
        )
        ENGINE = MergeTree
        ORDER BY tuple()
        SETTINGS min_bytes_for_wide_part = 0, min_rows_for_wide_part = 0
    