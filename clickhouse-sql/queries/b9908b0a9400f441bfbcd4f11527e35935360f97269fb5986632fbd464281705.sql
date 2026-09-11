
        INSERT INTO sort_lc
        SELECT
            toString(rand() % 1000) AS lc_col,
            rand64() AS uint_col
        FROM numbers(2000000)
    