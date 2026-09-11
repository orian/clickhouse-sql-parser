
        select key1_8, key2_8, sum(value) from test_group_by_strings group by key1_8, key2_8 settings max_threads = 32
    