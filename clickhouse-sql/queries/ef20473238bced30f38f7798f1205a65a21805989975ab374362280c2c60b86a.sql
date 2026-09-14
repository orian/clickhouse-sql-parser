
        WITH
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'dense_a') AS a,
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'dense_b') AS b
        SELECT sum(bitmapAndCardinality(materialize(a), b))
        FROM numbers(5000)
        FORMAT Null
    