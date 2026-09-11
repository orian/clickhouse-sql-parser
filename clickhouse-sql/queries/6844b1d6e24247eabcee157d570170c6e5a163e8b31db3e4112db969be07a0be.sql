
        CREATE TABLE map_basic_compact
        (
            id UInt64, m Map(String, String)
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 1000000000000, map_serialization_version = 'basic'
    