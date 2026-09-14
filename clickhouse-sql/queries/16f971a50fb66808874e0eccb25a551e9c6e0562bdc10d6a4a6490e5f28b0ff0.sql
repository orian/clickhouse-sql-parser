
        INSERT INTO t_projection_index_pools
        SELECT
            number,
            if(number BETWEEN 45000000 AND 45050000, 'needle', concat('hay_', toString(number % 64))),
            number, number * 2, number * 3
        FROM numbers(90000000)
        SETTINGS max_insert_threads = 8
    