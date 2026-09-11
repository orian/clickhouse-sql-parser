
        CREATE TABLE string_compact
        (
            id UInt64, s String
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 1000000000000, string_serialization_version = 'with_size_stream'
    