
        INSERT INTO bitmap_contains_32
        SELECT groupBitmapState(toUInt32(number))
        FROM numbers(1000000)
    