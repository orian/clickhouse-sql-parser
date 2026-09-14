
        WITH (SELECT bm FROM bitmap_contains_32) AS bm
        SELECT countIf(bitmapContains(bm, toUInt64(4294967296 + number % 1000000)))
        FROM numbers(10000000)
        FORMAT Null
    