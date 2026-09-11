
        INSERT INTO bitmap_cardinality_operands
        SELECT 'sparse_a', groupBitmapState(toUInt32(number * 100)) FROM numbers(300000)
    