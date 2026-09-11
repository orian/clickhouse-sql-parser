
        WITH
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'dense_a') AS a,
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'dense_a_copy') AS b
        SELECT sum(bitmapHasAny(materialize(a), b))
        FROM numbers(2000)
        FORMAT Null
    