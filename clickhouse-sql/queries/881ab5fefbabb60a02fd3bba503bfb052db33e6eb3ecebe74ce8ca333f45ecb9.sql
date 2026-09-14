
        CREATE TABLE a
        (
            d Date,
            os String,
            n UInt64
        )
        ENGINE = MergeTree
        PARTITION BY d
        ORDER BY (d, n)
    