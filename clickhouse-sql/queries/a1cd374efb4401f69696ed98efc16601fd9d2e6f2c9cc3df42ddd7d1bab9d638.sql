
        select key1_32, key2_32, sum(value) from test_group_by_strings group by key1_32, key2_32 settings max_threads = 32
    