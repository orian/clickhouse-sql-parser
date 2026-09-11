
        INSERT INTO str_sort
        SELECT 
            randomString(4) AS a,
            rand()  AS b, cast(randomString(5) AS FixedString(5)) as c,
            toString(rand() % 100000) as d,
            rand() % 1000 as e,
            rand() % 1000 as f
        FROM numbers(1e7)
    