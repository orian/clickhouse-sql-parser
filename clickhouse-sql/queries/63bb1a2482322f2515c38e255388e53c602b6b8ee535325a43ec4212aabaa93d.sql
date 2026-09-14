
        CREATE TABLE narrow_col (w UInt16)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS index_granularity = 8192
    