
        INSERT INTO array_join_function_cross
        SELECT arrayMap(x -> x + (number % 7),  range(50)),
               arrayMap(x -> x + (number % 11), range(50)),
               arrayMap(x -> x + (number % 13), range(50))
        FROM numbers(600)
    