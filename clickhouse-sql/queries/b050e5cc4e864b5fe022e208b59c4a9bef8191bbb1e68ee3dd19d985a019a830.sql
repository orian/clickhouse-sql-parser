
        WITH
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'evens') AS a,
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'odds') AS b
        SELECT sum(bitmapHasAny(materialize(a), b))
        FROM numbers(2000)
        FORMAT Null
    