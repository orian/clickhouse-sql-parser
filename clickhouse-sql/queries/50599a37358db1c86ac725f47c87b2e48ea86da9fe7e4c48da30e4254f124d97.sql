
        CREATE TABLE bitmap_cardinality_operands
        (
            name String,
            bm AggregateFunction(groupBitmap, UInt32)
        )
        ENGINE = Memory
    