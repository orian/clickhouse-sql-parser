
        INSERT INTO lc_many_marks_point_lookups
        SELECT
            'K' || toString(number % 200),
            number % 5000,
            ['alpha', 'beta', 'gamma'][(number % 3) + 1],
            'constant',
            'c3', 'c4', 'c5', 'c6',
            number
        FROM numbers(2000000)
        SETTINGS max_insert_threads = 1, max_insert_block_size = 2000000
    