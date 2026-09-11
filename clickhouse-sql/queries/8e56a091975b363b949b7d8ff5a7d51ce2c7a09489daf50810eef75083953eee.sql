INSERT INTO test_lwu_join
        SELECT
            number,
            rand(),
            rand(),
            randomPrintableASCII(8),
            randomPrintableASCII(8),
        FROM numbers_mt(20000000) SETTINGS max_insert_threads=1
    