
        CREATE TABLE adaptive_codec_merge_sizes (s String)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS enable_adaptive_codec_selection = 1
    