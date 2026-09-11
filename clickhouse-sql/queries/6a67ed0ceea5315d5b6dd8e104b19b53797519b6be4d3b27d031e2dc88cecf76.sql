
        INSERT INTO t_skip_index_pools
        SELECT
            number,
            if(number BETWEEN 45000000 AND 45050000, 9999, number % 64),
            number, number * 2, number * 3
        FROM numbers(90000000)
        SETTINGS max_insert_threads = 8
    