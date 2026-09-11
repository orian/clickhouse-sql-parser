
        INSERT INTO bitmap_cardinality_operands
        SELECT 'odds', groupBitmapState(toUInt32(number * 2 + 1)) FROM numbers(1500000)
    