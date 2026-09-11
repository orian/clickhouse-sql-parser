
        select key1_16, key2_16, sum(value) from test_group_by_strings group by key1_16, key2_16 settings max_threads = 16
    