
        INSERT INTO bitmap_cardinality_operands
        SELECT 'sparse_b', groupBitmapState(toUInt32(number * 100 + 15000000)) FROM numbers(300000)
    