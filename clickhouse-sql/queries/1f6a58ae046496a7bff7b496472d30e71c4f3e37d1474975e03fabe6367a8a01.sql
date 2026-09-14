
        CREATE TABLE mt_comp_parts
        ENGINE = MergeTree
        ORDER BY (c1, c2)
        SETTINGS min_rows_for_wide_part = 1000000000 AS
        SELECT *
        FROM generateRandom('c1 UInt32, c2 UInt64, s1 String, arr1 Array(UInt32), c3 UInt64, s2 String', 0, 5, 6)
        LIMIT 50000000
        SETTINGS max_insert_threads = 8
    