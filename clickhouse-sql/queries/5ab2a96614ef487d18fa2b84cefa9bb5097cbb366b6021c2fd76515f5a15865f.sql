
        INSERT INTO bitmap_cardinality_operands
        SELECT 'dense_b', groupBitmapState(toUInt32(number + 1500000)) FROM numbers(3000000)
    