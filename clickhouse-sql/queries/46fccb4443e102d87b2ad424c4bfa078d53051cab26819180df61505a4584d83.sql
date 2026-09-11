
        CREATE TABLE array_join_filter_fusion (id UInt64, arr Array(String), payload String)
        ENGINE = MergeTree ORDER BY id
    