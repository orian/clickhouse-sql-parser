
        CREATE TABLE json_compact
        (
            id UInt64, json JSON(max_dynamic_paths = 0)
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 1000000000000,
                 object_shared_data_serialization_version = 'map', object_shared_data_serialization_version_for_zero_level_parts = 'map'
    