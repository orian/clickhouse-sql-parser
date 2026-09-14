
        CREATE TABLE map_buckets_compact
        (
            id UInt64, m Map(String, String)
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 1000000000000,
                 map_serialization_version = 'with_buckets', map_serialization_version_for_zero_level_parts = 'with_buckets',
                 map_buckets_strategy = 'constant', max_buckets_in_map = 2, map_buckets_min_avg_size = 0
    