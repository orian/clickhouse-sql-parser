
        INSERT INTO sort_string_full
        SELECT randomPrintableASCII(rand() % 100 + 1) AS str_col
        FROM numbers(3000000)
    