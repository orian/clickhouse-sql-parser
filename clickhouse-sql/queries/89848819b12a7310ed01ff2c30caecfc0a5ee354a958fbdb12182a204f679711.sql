INSERT INTO test_lwu_merge
        SELECT
            number,
            rand(),
            rand(),
            randomPrintableASCII(8),
            randomPrintableASCII(8),
        FROM numbers_mt(50000000) SETTINGS max_insert_threads=8
    