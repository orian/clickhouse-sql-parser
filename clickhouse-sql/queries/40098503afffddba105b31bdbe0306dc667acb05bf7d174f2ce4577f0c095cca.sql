
        CREATE TABLE date_preimage_perf
        (
            id UInt64,
            seen DateTime('UTC'),
            d Date,
            INDEX idx_seen seen TYPE minmax GRANULARITY 1,
            INDEX idx_d d TYPE minmax GRANULARITY 1
        )
        ENGINE = MergeTree
        ORDER BY id
        SETTINGS index_granularity = 256
    