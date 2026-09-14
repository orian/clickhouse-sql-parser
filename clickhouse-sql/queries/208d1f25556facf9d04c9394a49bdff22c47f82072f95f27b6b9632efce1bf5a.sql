
        INSERT INTO t_subcolumns SELECT range(number % 20), toString(number), mapFromArrays(range(number % 20), range(number % 20)) FROM numbers_mt(50000000)
    