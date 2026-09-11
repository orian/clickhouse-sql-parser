
        INSERT INTO text_index_lazy_wide
        SELECT number,
            concat('common',
                if(number % 2 = 0, ' half', ''),
                if(number % 3 = 0, ' third', ''),
                if(number % 4 = 0, ' quarter', ''),
                if(number % 5 = 0, ' fifth', ''),
                if(number % 10 = 0, ' dense', ''),
                if(number % 15 = 0, ' medium', ''),
                if(number % 50 = 7, ' sparse', ''),
                ' filler', toString(number % 1000))
        FROM numbers(50000000)
        SETTINGS max_insert_threads = 8
    