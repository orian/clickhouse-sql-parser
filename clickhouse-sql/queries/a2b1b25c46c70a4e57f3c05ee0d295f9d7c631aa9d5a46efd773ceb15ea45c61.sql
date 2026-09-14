
        INSERT INTO sort_nearly_sorted
        SELECT number + if(rand() % 20 = 0, rand() % 1000000, 0) AS key
        FROM numbers(10000000)
    