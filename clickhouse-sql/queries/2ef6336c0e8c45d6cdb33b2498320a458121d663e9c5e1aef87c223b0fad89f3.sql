
        CREATE TABLE t_projection_index_pools
        (
            id UInt64,
            category String,
            v1 UInt64,
            v2 UInt64,
            v3 UInt64,
            PROJECTION category_proj INDEX category TYPE basic
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS index_granularity = 256
    