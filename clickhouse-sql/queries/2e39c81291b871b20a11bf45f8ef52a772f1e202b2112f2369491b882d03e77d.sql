
        INSERT INTO bitmap_cardinality_operands
        SELECT 'evens', groupBitmapState(toUInt32(number * 2)) FROM numbers(1500000)
    