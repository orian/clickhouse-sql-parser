
        CREATE TABLE bitmap_contains_32
        (
            bm AggregateFunction(groupBitmap, UInt32)
        )
        ENGINE = Memory
    