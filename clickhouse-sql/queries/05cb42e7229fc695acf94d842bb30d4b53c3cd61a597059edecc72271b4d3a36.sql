
        INSERT INTO vrow_filtered
        SELECT number, number >= 63000000 FROM numbers(64000000)
        SETTINGS max_block_size = 1000000, max_insert_block_size = 1000000,
                 min_insert_block_size_rows = 1000000, min_insert_block_size_bytes = 0
    