
        select key1_64, key2_64, sum(value) from test_group_by_strings group by key1_64, key2_64 settings max_threads = 16
    