
        CREATE TABLE tuple_wide
        (
            id UInt64, skip UInt8,
            t Tuple(a String, b String, c String, d String, e String),
            INDEX idx_skip skip TYPE minmax GRANULARITY 1
        )
        ENGINE = MergeTree ORDER BY id
        SETTINGS min_bytes_for_wide_part = 0, index_granularity = 2048, index_granularity_bytes = 0,
                 string_serialization_version = 'single_stream'
    