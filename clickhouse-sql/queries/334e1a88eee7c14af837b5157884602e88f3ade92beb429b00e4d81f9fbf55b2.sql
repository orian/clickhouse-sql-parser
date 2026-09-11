
        CREATE TABLE t_skip_index_pools
        (
            id UInt64,
            category UInt64,
            v1 UInt64,
            v2 UInt64,
            v3 UInt64,
            INDEX category_idx category TYPE minmax GRANULARITY 1
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS index_granularity = 256
    