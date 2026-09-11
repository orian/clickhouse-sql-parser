
        WITH
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'sparse_a') AS a,
            (SELECT bm FROM bitmap_cardinality_operands WHERE name = 'sparse_b') AS b
        SELECT sum(bitmapAndCardinality(materialize(a), b))
        FROM numbers(2000)
        FORMAT Null
    