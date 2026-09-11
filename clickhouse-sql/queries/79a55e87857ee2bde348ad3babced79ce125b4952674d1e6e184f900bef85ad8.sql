
        INSERT INTO t_streaming_read_round
        SELECT number, toDateTime64('2020-01-01 00:00:00', 3) + number
        FROM numbers(100000000)
        SETTINGS max_insert_threads = 1, max_insert_block_size = 2000000, min_insert_block_size_rows = 2000000, min_insert_block_size_bytes = 0
    