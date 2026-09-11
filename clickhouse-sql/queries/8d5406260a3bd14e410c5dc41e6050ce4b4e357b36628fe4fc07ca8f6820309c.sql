
        CREATE TABLE tuple_compact
        (
            id UInt64,
            t Tuple(inner Tuple(a String, b String, c String, d String, e String))
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 1000000000000, string_serialization_version = 'single_stream'
    