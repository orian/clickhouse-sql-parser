
        CREATE TABLE adaptive_codec_merge (x UInt64)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS enable_adaptive_codec_selection = 1
    