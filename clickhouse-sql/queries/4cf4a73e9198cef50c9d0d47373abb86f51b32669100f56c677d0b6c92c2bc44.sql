
        INSERT INTO bitmap_cardinality_operands
        SELECT 'dense_a_copy', groupBitmapState(toUInt32(number)) FROM numbers(3000000)
    