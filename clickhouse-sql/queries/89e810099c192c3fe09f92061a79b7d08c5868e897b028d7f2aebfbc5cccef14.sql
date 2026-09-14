
        CREATE TABLE map_buckets_wide
        (
            id UInt64, skip UInt8, m Map(String, String),
            INDEX idx_skip skip TYPE minmax GRANULARITY 1
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 0, index_granularity = 2048, index_granularity_bytes = 0,
                 map_serialization_version = 'with_buckets', map_serialization_version_for_zero_level_parts = 'with_buckets',
                 map_buckets_strategy = 'constant', max_buckets_in_map = 2, map_buckets_min_avg_size = 0
    