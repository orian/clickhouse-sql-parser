
        INSERT INTO bitmap_cardinality_operands
        SELECT 'dense_a', groupBitmapState(toUInt32(number)) FROM numbers(3000000)
    