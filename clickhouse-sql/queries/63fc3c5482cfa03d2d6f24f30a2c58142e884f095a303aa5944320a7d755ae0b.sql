
    CREATE TABLE enum16_wide_sparse
    (
        `c` Enum16('' = -30000, 'a' = -1, 'A' = 0, 'AB' = 7, 'aBc' = 300, 'ABCD' = 1000, 'xAyz' = 5000, 'Foo A Bar' = 30000)
    )
    ENGINE = MergeTree
    ORDER BY tuple()
    SETTINGS index_granularity = 8192
    