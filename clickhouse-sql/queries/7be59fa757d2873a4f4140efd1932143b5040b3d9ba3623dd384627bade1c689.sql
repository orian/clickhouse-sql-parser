
        INSERT INTO test_minmax_index
        SETTINGS max_insert_threads = 64, max_insert_block_size = 65536, min_insert_block_size_rows = 1, min_insert_block_size_bytes = 1
        SELECT * FROM numbers(1000000);
    