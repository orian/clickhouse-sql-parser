
        select key1_512, key2_512, sum(value) from test_group_by_strings group by key1_512, key2_512 settings max_threads = 16
    