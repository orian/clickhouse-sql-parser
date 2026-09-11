
        INSERT INTO sort_mixed
        SELECT
            rand64() AS uint_col,
            randomPrintableASCII(rand() % 32 + 1) AS str_col
        FROM numbers(2000000)
    